package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/httpreserve/httpreserve"
	"github.com/httpreserve/wayback"
)

var (
	//version
	vers bool

	//individual links
	link  string
	label string
	save  bool

	//output methods
	boltdb  bool
	jsonout bool
	csvout  bool

	//list processing
	list string
)

func init() {
	// Return version information.
	flag.BoolVar(&vers, "v", false, "return application versions")
	flag.BoolVar(&vers, "version", false, "")

	// Flags to return a single result.
	flag.StringVar(&link, "link", "", "seek the status of a single URL: JSON")
	flag.StringVar(&label, "fname", "", "annotate a response with a filename")
	flag.BoolVar(&save, "save", false, "save the link we're querying")

	// Flags to batch results.
	flag.StringVar(&list, "list", "", "provide a list of URLs to test against in CSV format.")

	// Output method flags.
	flag.BoolVar(&boltdb, "bolt", false, "output to static BoltDB.")
	flag.BoolVar(&jsonout, "json", false, "output to JSON.")
	flag.BoolVar(&csvout, "csv", false, "output to CSV.")
}

func getJSONFromLocal(link string, label string) string {

	if save {
		log.Println("saving url to wayback")
		_, err := wayback.SubmitToInternetArchive(link, httpreserve.VersionText())
		if err != nil {
			if !strings.Contains(fmt.Sprintf("%s", err), wayback.SaveTooMany) {
				log.Printf("error: %s, attempting to return stats if they exist", err)
			}
		}

	}

	ls, err := httpreserve.GenerateLinkStats(link, label, false)
	if err != nil {
		log.Println("Error retrieving linkstat JSON may be incorrect:", err)
	}
	ls.ScreenShot = ""
	js := httpreserve.MakeLinkStatsJSON(ls)

	// throttle requests to the server somehow...
	time.Sleep(500 * time.Millisecond)

	// return json...
	return js
}

func getLocalLink() {
	log.Println("using httpreserve libs to retrieve data")
	js := getJSONFromLocal(link, label)
	fmt.Fprintf(os.Stdout, "%s", js)
}

func linkstat() {
	if link != "" {
		getLocalLink()
		return
	}
	if list == "" {
		log.Println("url list not supplied, other modes require a list of urls")
		os.Exit(1)
	}
	if jsonout {
		fmt.Fprintf(os.Stdout, "%s", outputJSONHeader())
		listHandler(jsonHandler)
		outputjsonpool()
		fmt.Fprintf(os.Stdout, "%s", outputJSONFooter())
		return
	}
	if csvout {
		//output JSON header
		fmt.Fprintf(os.Stdout, "%s", outputCSVHeader())
		listHandler(csvHandler)
		return
	}
	if boltdb {
		openKVALBolt()
		defer closeKVALBolt()
		listHandler(boltdbHandler)
		return
	}
}

func main() {
	setLog("linkstat", true)
	flag.Parse()
	if vers {
		fmt.Fprintf(os.Stderr, "%s\n", getVersion())
		fmt.Fprintf(os.Stderr, "%s\n", httpreserve.VersionText())
		os.Exit(0)
	} else if flag.NFlag() <= 0 {
		flag.Usage()
		os.Exit(0)
	}
	startTime := time.Now()
	linkstat()
	log.Println("data retrieved in:", time.Since(startTime))
}

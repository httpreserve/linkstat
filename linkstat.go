package main

import (
	"flag"
	"fmt"
	"github.com/httpreserve/httpreserve"
	"log"
	"os"
	"time"
)

var (
	//version
	vers bool

	//individual links
	link  string
	label string

	//output methods
	boltdb  bool
	jsonout bool
	csvout  bool

	//list processing
	list string
)

func init() {
	// Return version information.
	flag.BoolVar(&vers, "version", false, "Return httpreserve version.")
	flag.BoolVar(&vers, "v", false, "Return httpreserve version.")

	// Flags to return a single result.
	flag.StringVar(&link, "link", "", "Seek the status of a single URL: JSON")
	flag.StringVar(&label, "label", "", "Annotate single URL check response with label.")

	// Flags to batch results.
	flag.StringVar(&list, "list", "", "Provide a list of URLs to test against in CSV format.")

	// Output method flags.
	flag.BoolVar(&boltdb, "bolt", false, "Output to static BoltDB.")
	flag.BoolVar(&jsonout, "json", false, "Output to JSON.")
	flag.BoolVar(&csvout, "csv", false, "Output to CSV.")
}

func getJSONFromLocal(link string, label string) string {
	ls, err := httpreserve.GenerateLinkStats(link, label, true)
	if err != nil {
		log.Println("Error retrieving linkstat JSON may be incorrect:", err)
	}
	js := httpreserve.MakeLinkStatsJSON(ls)

	// throttle requests to the server somehow...
	time.Sleep(500 * time.Millisecond)

	// return json...
	return js
}

func getLocalLink() {
	js := getJSONFromLocal(link, label)
	fmt.Fprintln(os.Stderr, "Using httpreserve libs to retrieve data.")
	fmt.Fprintf(os.Stdout, "%s", js)
}

var htmcomplete bool
var starttime time.Time
var elapsedtime time.Duration

func programrunner() {
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
	if link != "" {
		getLocalLink()
	}
}

func main() {
	flag.Parse()
	if vers {
		fmt.Fprintf(os.Stderr, "%s\n", getVersion())
		fmt.Fprintf(os.Stderr, "%s\n", httpreserve.VersionText())
		os.Exit(0)
	} else if flag.NFlag() <= 0 {
        fmt.Fprintln(os.Stderr, "Usage:  linkstat [Optional -link] [Optional -label]")
		fmt.Fprintln(os.Stderr, "                 [Optional -list] [Optional -json]")
		fmt.Fprintln(os.Stderr, "                                  [Optional -bolt]")
		fmt.Fprintln(os.Stderr, "                                  [Optional -csv]")
		fmt.Fprintln(os.Stderr, "                 [Optional -version -v]")
		fmt.Fprintln(os.Stderr, "")
		fmt.Fprintln(os.Stderr, "Output: [Json]")
		fmt.Fprintln(os.Stderr, "Output: [CSV]")
		fmt.Fprintln(os.Stderr, "Output: [BoltDB]")
		fmt.Fprintf(os.Stderr, "Output: [Version] '%s ...'\n", httpreserve.VersionText())
		fmt.Fprintln(os.Stderr, "")
		flag.Usage()
		os.Exit(0)
	}
	programrunner()
}

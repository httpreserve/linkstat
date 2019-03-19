package main

import (
	"encoding/json"
	"fmt"
	"github.com/httpreserve/httpreserve"
	kval "github.com/kval-access-language/kval-bbolt"
	"github.com/speps/go-hashids"
	"log"
	"os"
	"time"
)

// values to use to create hashid
var salt = "httpreserve"
var namelen = 8

// bucket constants
const linkIndex = "link index"

//const fnameIndex = "filename index"
const hashIndex = "hash index"

// location for bolt databases. Names are random at present because
// I'm unsure how they're going to be used in future so may write naming
// functions and flags at a later date.
const boltdir = "db/"

// For stdout the name of the database
var boltoutput string

// getNewDBName provides three integers based on the time at
// which we run the code to help us create a hashid name for the db.
func getNewDBName() []int {
	t := time.Now()
	i1 := t.Minute()
	i2 := t.Second()
	i3 := t.Nanosecond()
	return []int{i1, i2, i3}
}

// configureHashID will create a hashids name for our database
func configureHashID() string {

	name := getNewDBName()

	//hashdata
	hd := hashids.NewData()
	hd.Salt = salt
	hd.MinLength = namelen

	//hash
	h, _ := hashids.NewWithData(hd)
	e, _ := h.Encode(name)
	return e
}

// makeIDIndex will write rows to the BoldDB based on an MD5 hash value
// associated with the lmap passed to the function (a deconstructed LinkStats)
func makeIDIndex(kb kval.Kvalboltdb, lmap map[string]interface{}) {
	for k, v := range lmap {
		_, err := kval.Query(kb, "INS "+convertInterface(lmap["response code"])+">>"+convertInterface(lmap["link"])+" >>>> "+k+" :: "+convertInterface(v))
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
	}
}

// makeBoltDir will create a database for all BoldDB files generated
// if the database doesn't already exist.
func makeBoltDir() {
	if _, err := os.Stat(boltdir); os.IsNotExist(err) {
		err := os.Mkdir(boltdir, 0700)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
	}
}

// boltGetResultContainers returns the names of all top level buckets
// n.b. these functions are heavily linked to the database schema and
// could be made more generic with more effort.
func boltGetResultContainers(kb kval.Kvalboltdb) []string {
	var buckets []string
	q := "GET " + hashIndex
	res, _ := kval.Query(kb, q)
	for k := range res.Result {
		buckets = append(buckets, k)
	}
	return buckets
}

// boltGetSingleRecord will return a single record for a given md5 key
// n.b. these functions are heavily linked to the database schema and
// could be made more generic with more effort.
func boltGetSingleRecord(kb kval.Kvalboltdb, md5Key string) map[string]string {
	records := make(map[string]string)
	q := "GET " + hashIndex + " >> " + md5Key
	res, _ := kval.Query(kb, q)
	for k, v := range res.Result {
		records[k] = v
	}
	return records
}

// boltGetAllRecords returns all records in all top level buckets in the
// database.
// n.b. these functions are heavily linked to the database schema and
// could be made more generic with more effort.
func boltGetAllRecords(kb kval.Kvalboltdb) []map[string]string {
	var records []map[string]string
	keys := boltGetResultContainers(kb)
	for _, v := range keys {
		records = append(records, boltGetSingleRecord(kb, v))
	}
	return records
}

var kb kval.Kvalboltdb

func openKVALBolt() {
	var err error
	boltname := configureHashID()
	makeBoltDir()

	boltoutput = boltdir + "HP_" + boltname + ".bolt"
	kb, err = kval.Connect(boltoutput)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error opening bolt database: %+v\n", err)
		os.Exit(1)
	}
}

func closeKVALBolt() {
	kval.Disconnect(kb)
}

var id []string

// boltdbHandler is the primary handler for writing to a BoltDB
// from our httpreserve results rsets.
func boltdbHandler(js string) {

	var ls httpreserve.LinkStats

	err := json.Unmarshal([]byte(js), &ls)
	if err != nil {
		fmt.Fprintln(os.Stderr, "problem unmarshalling data.", err)
	}

	var add = true

	// retrieve a map from the structure and write it out to the
	// bolt db.
	lmap := storeStruct(ls, js)
	if len(lmap) > 0 {
		makeIDIndex(kb, lmap)

		lmapid := convertInterface(lmap["id"])
		for x := range id {
			if lmapid == id[x] {
				add = false
				log.Println("Already seen:", lmap["filename"], lmap["title"])
				break
			}
		}
		if add {
			makeIDIndex(kb, lmap)
			id = append(id, lmapid)
		}
	}
}

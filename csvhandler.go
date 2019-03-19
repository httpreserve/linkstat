package main

import (
	"encoding/json"
	"fmt"
	"github.com/httpreserve/httpreserve"
	"os"
	"strings"
)

var csvHeader = []string{"id", "filename", "link", "response code", "response text", "title",
	"content-type",
	"archived", "internet archive response code", "internet archive response text",
	"wayback earliest date", "internet archive earliest",
	"wayback latest date", "internet archive latest", "internet archive save link",
	"protocol error", "protocol error",
	"analysis version number", "analysis version text", "stats creation time"}

func outputCSVHeader() string {
	var header string
	header = "\"" + strings.Join(csvHeader, "\",\"") + "\"" + "\n"
	return header
}

func outputCSVRow(lmap map[string]interface{}) string {
	var row []string
	for x := range csvHeader {
		if val, ok := lmap[csvHeader[x]]; ok {
			var v string
			switch val.(type) {
			case string:
				v = fmt.Sprintf("%s", val)
				v = strings.Replace(v, "\"", "'", -1)
				v = fmt.Sprintf("\"%s\"", v)
			case int:
				v = fmt.Sprintf("\"%d\"", val)
			case bool:
				v = fmt.Sprintf("\"%t\"", val)
			}
			row = append(row, v)
		} else {
			row = append(row, "\"\"")
		}
	}
	return strings.Join(row, ",")
}

// TODO: consider more idiomatic approaches to achieving what we do here,
// that is, fmt.Println() is not really my approved approach (but it works)
func csvHandler(js string) {
	var ls httpreserve.LinkStats
	err := json.Unmarshal([]byte(js), &ls)
	if err != nil {
		fmt.Fprintln(os.Stderr, "problem unmarshalling data.", err)
	}

	// retrieve a map from the structure and write it out to the CSV
	lmap := storeStruct(ls, js)
	if len(lmap) > 0 {
		fmt.Fprintf(os.Stdout, "%s\n", outputCSVRow(lmap))
	}
}

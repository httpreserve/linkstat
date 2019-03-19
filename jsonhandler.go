package main

import (
	"fmt"
	"os"
	"sync"
)

func outputJSONHeader() string {
	var header string
	header = header + fmt.Sprintf("%s\n", "{")
	header = header + fmt.Sprintf("  \"%s\": \"%s\",\n", "title", "httpreserve-linkstat")
	header = header + fmt.Sprintf("  \"%s\": \"%s\",\n", "description", "httpreserve-linkstat-output")
	header = header + fmt.Sprintf("  \"%s\": %s\n", "data", "[")
	return header
}

func outputJSONFooter() string {
	var footer string
	footer = footer + fmt.Sprintf("%s\n%s\n", "]", "}")
	return footer
}

var jsonCount int

// webappHanlder enables us to establish the web server and create
// the structures we need to present our data to the user...
func jsonHandler(js string) {
	wg := new(sync.WaitGroup)
	wg.Add(1)
	go makejsonpool(js, wg)
	wg.Wait()
	return
}

var jsonpool []string

func makejsonpool(js string, wg *sync.WaitGroup) {
	defer wg.Done()
	jsonpool = append(jsonpool, js)
}

func outputjsonpool() {
	for j := range jsonpool {
		if j+1 < len(jsonpool) {
			fmt.Fprint(os.Stdout, jsonpool[j]+",")
		} else {
			fmt.Fprint(os.Stdout, jsonpool[j])
		}
	}
}

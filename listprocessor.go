package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"sync"
)

var pscomplete = false

// list handler to help us kick off some go channels
// we pass a first class function to help route our output
func listHandler(outputHandler func(js string)) {

	link := make(chan map[string]string)
	results := make(chan string)

	wg := new(sync.WaitGroup)

	// batches of two... helps us to batch out work, e.g. to throttle
	// server requests... two requests per second, IN THEORY!
	for w := 0; w <= 2; w++ {
		wg.Add(1)
		go getJSON(link, results, wg)
	}

	// Create a link map for output to the output handlers
	go func() {
		// Read all the lines in a file...
		file, err := os.Open(list)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error with scanner: %s\n", err.Error())
		}
		defer file.Close()

		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			l := make(map[string]string)
			row := scanner.Text()
			var split []string
			if strings.Contains(row, "\",\"") {
				split = strings.Split(scanner.Text(), "\",\"")
			} else if strings.Contains(row, "\", \"") {
				split = strings.Split(scanner.Text(), "\", \"")
			} else {
				split = strings.Split(scanner.Text(), ",")
			}
			if len(split) != 2 {
				fmt.Fprintf(os.Stderr, "ignoring: issue reading string from file: %s\n", scanner.Text())
			} else {
				l[strings.Trim(split[1], " ")] = strings.Trim(split[0], " ")
				link <- l
			}
		}

		if err := scanner.Err(); err != nil {
			fmt.Fprintf(os.Stderr, "Error with scanner: %s\n", scanner.Text())
		}
		close(link)
	}()

	go func() {
		wg.Wait()
		pscomplete = true
		linklen++
		close(results)
	}()

	for js := range results {
		outputHandler(js)
	}
}

var linklen int

func getJSON(link <-chan map[string]string, results chan<- string, wg *sync.WaitGroup) {
	defer wg.Done()
	for m := range link {
		// k filename, v link...
		for k, v := range m {
			k = strings.Trim(k, "\"")
			v = strings.Trim(v, "\"")
			results <- getJSONFromLocal(k, v)
		}
	}
}

// retrieve a JSON output from HTTPreserve without talking to the server
func httpreserveJSONOutput(link string, filename string) string {
	js := getJSONFromLocal(link, filename)
	return js
}

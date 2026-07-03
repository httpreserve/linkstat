package main

import (
	"fmt"
	"log"
	"os"
	"time"
)

type LogWriter struct {
	Appname string
	UTC     bool
}

const logTimeFormat = "2006-01-02 15:04:05"

func (lw *LogWriter) Write(logString []byte) (int, error) {
	logTime := time.Now().UTC().Format(logTimeFormat)
	if !lw.UTC {
		logTime = time.Now().Format(logTimeFormat)
	}
	return fmt.Fprintf(os.Stderr, "%s :: %s :: %s",
		logTime,
		lw.Appname,
		string(logString),
	)
}

func setLog(app string, utc bool) {
	log.SetFlags(0 | log.Lshortfile | log.LUTC)
	lw := &LogWriter{}
	lw.Appname = app
	lw.UTC = utc
	log.SetOutput(lw)
}

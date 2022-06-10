package utils

import (
	"log"
	"os"
)

var (
	// Warning warning log
	Warning *log.Logger
	// Info information log
	Info *log.Logger
	// Error error log
	Error *log.Logger
)

func init() {
	Warning = log.New(os.Stdout, "[WARN] ", log.Ldate|log.Ltime|log.Lshortfile)
	Info = log.New(os.Stdout, "[INFO] ", log.Ldate|log.Ltime|log.Lshortfile)
	Error = log.New(os.Stderr, "[ERR] ", log.Ldate|log.Ltime|log.Lshortfile)
}

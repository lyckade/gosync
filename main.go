package main

import (
	"flag"
	"fmt"
	"os"
)

// Version defines the version number of the app
const Version = "0.7.0"

var flagVersion = flag.Bool("version", false, "Version of this app.")
var flagVerbose = flag.Bool("v", false, "increase verbosity")
var flagLog = flag.Bool("log", false, "Actions are written to gosync.log")

func init() {
	flag.Parse()
}

func main() {
	if *flagVersion {
		fmt.Printf("Version: %s", Version)
		os.Exit(-1)
	}
}

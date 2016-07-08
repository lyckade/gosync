package main

import (
	"errors"
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/lyckade/gosync/osfsyncer"
	"github.com/lyckade/gosync/sync"
)

// Version defines the version number of the app
const Version = "0.7.1"

var ErrNotEnoughArguments = errors.New("Not enough arguments!")

var flagVersion = flag.Bool("version", false, "Version of this app.")
var flagVerbose = flag.Bool("v", false, "increase verbosity")

//var flagLog = flag.Bool("log", false, "Actions are written to gosync.log")

var logger = log.New(os.Stdout, "", log.LstdFlags)

func init() {
	flag.Parse()
}

func main() {
	if *flagVersion {
		fmt.Printf("Version: %s\n", Version)
		os.Exit(-1)
	}
	if len(os.Args) < 3 {
		fmt.Print(ErrNotEnoughArguments)
	}
	syncFolder, _ := filepath.Abs(os.Args[1])
	distFolder, _ := filepath.Abs(os.Args[2])
	logger.Printf("Sync started for %s\n", syncFolder)
	filepath.Walk(syncFolder, func(fpath string, info os.FileInfo, err error) error {
		if info.IsDir() {
			return err
		}
		/*ignore, err := MatchPath(fpath, properties.Ignore)
		if err != nil || ignore {
			fmt.Println(fpath)
			return err
		}*/
		var syncer osfsyncer.Osfsyncer
		dpath, err := sync.MakeDistPath(fpath, syncFolder, distFolder)
		if err != nil {
			logger.Fatalln(err)
		}
		lg(fpath)
		err = sync.Sync(&syncer, fpath, dpath)
		if err != nil {
			logger.Fatalln(err)
		}
		return err
	})
}

func lg(s string) {
	if *flagVerbose {
		logger.Println(s)
	}
	logger.Println(s)
}

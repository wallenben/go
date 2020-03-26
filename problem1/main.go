package main

import (
	"flag"
	//for scanning the file
	"fmt"
	//formatting package
	"os"
	//os-related package
)

func main() {
	csvFilename := flag.String("csv", "problems.csv", "flags: -csv for csvfile")
	//csvfilename equals the string values called by flag.string
	flag.Parse()

	file, err := os.Open(*csvFilename)
	_ = file
	if err != nil {
		fmt.Printf("failed to open file %s", *csvFilename)
		os.Exit(1)
	}
}

package main

import (
	"flag"
	"fmt"
	"os"
)


func printHelpAndExit() {
	help := `Usage: hex2bin [FILE]
The inversion of bin2hex, converts hex decimal representation to binary.
If no FILE has been given, the tool will read from STDIN.
`
	fmt.Fprintf(os.Stderr, help)
	flag.PrintDefaults()
	os.Exit(1)
}

func main() {
	flag.Usage = printHelpAndExit
	flag.Parse()
	if flag.NArg() > 1 {
		return
	}
	var file *os.File
	if flag.NArg() > 0 {
		var err error
		fileName := flag.Arg(0)
		file, err = os.Open(fileName)
		if err != nil {
			fmt.Fprintf(os.Stderr, "cannot read %s: %v", fileName, err)
			os.Exit(1)
			return
		}
	} else {
		file = os.Stdin
	}
	_ = file
}

package main

import (
	"bufio"
	"encoding/hex"
	"flag"
	"fmt"
	"os"
	"strings"
)


func printHelpAndExit() {
	help := `Usage: hex2bin [FILE]
The inversion of bin2hex, converts hex decimal representation to binary.
If no FILE has been given, the tool well read from STDIN.
`
	fmt.Fprintf(os.Stderr, help)
	flag.PrintDefaults()
	os.Exit(1)
}

func main() {
	flag.Usage = printHelpAndExit
	flag.Parse()
	if flag.NArg() > 1 {
		printHelpAndExit()
		return
	}
	var file *os.File
	if flag.NArg() > 0 {
		var err error
		fileName := flag.Arg(1)
		file, err = os.Open(fileName)
		if err != nil {
			fmt.Fprintf(os.Stderr, "cannot open '%s': %v", fileName, err)
			os.Exit(1)
			return
		}
	} else {
		file = os.Stdin
	}
	eof := false
	reader := bufio.NewReader(file)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()
	for !eof {
		line, err := reader.ReadString('\n')
		if err != nil {
			eof = true
		}
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}
		parts := strings.Split(line, "  ")
		hexString := strings.Join([]string{parts[1], parts[2]}, " ")
		hexString = strings.ReplaceAll(hexString, " ", "")
		data, err := hex.DecodeString(hexString)
		if err != nil {
			fmt.Fprintf(os.Stderr, "failed decode hex string '%s': %v", hexString, err)
			os.Exit(1)
			return
		}
		_, err = writer.Write(data)
		if err != nil {
			fmt.Fprintf(os.Stderr, "failed tor write data: %v", err)
			os.Exit(1)
			return
		}
	}
}

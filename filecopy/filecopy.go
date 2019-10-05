/*
 * HomeWork-6: FileCopy utility like dd
 * Created on 05.10.2019 14:26
 * Copyright (c) 2019 - Eugene Klimov
 */

package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
)

var fromFile, toFile string
var offset, limit int

func init() {

	// set the custom Usage function
	fileName := filepath.Base(os.Args[0])
	flag.Usage = func() {
		fmt.Printf("usage: %s -from <source> -to <destination> [-offset bytes] [-limit bytes]\n", fileName)
		fmt.Printf("example: %s -from /path/to/source -to /path/to/dest -offset 1024 -limit 2048\n", fileName)
		flag.PrintDefaults()
	}

	flag.StringVar(&fromFile, "from", "", "file name to read from")
	flag.StringVar(&toFile, "to", "", "file name to write to")
	flag.IntVar(&offset, "offset", 0, "offset in input file, bytes")
	flag.IntVar(&limit, "limit", 0, "limit, bytes")
}

func main() {
	flag.Parse()

	if fromFile == "" || toFile == "" {
		flag.Usage()
		os.Exit(2)
	}

	fmt.Println(fromFile, toFile, offset, limit)

}

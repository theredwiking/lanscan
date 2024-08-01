package main

import (
	"flag"

	"github.com/theredwiking/lanscan/cmd"
)

func main() {
	var scan = flag.Bool("scan", false, "Starts scan of lan")
	var _ = flag.String("file", "lan.json", "Name of file output")
	var display = flag.Bool("v", false, "Shows output")

	flag.Parse()

	if *scan {
		cmd.Scan(*display)
	}
}

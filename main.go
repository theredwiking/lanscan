package main

import (
	"flag"

	"github.com/theredwiking/lanscan/cmd"
	"github.com/theredwiking/lanscan/models"
)

func main() {
	scanConf := models.Scan{}

	var scan = flag.Bool("scan", false, "Starts scan of lan")
	flag.StringVar(&scanConf.File, "file", "./lan.json", "Name of file output")
	flag.BoolVar(&scanConf.Display, "all", false, "Shows all scanned devices")

	flag.Parse()

	if *scan {
		cmd.Scan(scanConf)
	}
}

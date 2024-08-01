package cmd

import (
	"fmt"
	"log"
	"time"

	"github.com/theredwiking/lanscan/core"
	"github.com/theredwiking/lanscan/models"
)

func Scan(config models.Scan) {
	fmt.Println("Starting scan")

	if err := core.CreateFile(config.File); err != nil {
		log.Println(err)
		return
	}

	ips := make(chan string, 254)
	devices := make(chan models.Device, 254)

	scanResult := []models.Device{}
	
	lan, err := core.LanIP()
	if err != nil {
		log.Println(err)
		return
	}
	
	ipRange, err := core.IpRange(lan)
	if err != nil {
		log.Println(err)
		return
	}
	
	for range 14 {
		go core.Ping(ips, devices)
	}

	if config.Display {
		go core.OutputDisplay(devices)
	}

	ticker := time.NewTicker(500 * time.Millisecond)

	for {
		select {
		case device := <- devices:
			scanResult = append(scanResult, device)
		case _ = <- ticker.C:
			if len(ipRange) != 0 {
				ips <- ipRange[0]
				if len(ipRange) >= 2 {
					ipRange = append(ipRange[:0], ipRange[1:]...)
				} else {
					ipRange = []string{}
				}
			}

			if len(ipRange) == 0 && len(devices) == 0 && len(ips) == 0 {
				fmt.Println("Scan complete")	
				if err = core.WriteFile(config.File, scanResult); err != nil {
					log.Println(err)
					return
				}
				fmt.Printf("Wrote data to file: %s", config.File)
				return
			}
		}
	}


}

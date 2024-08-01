package cmd

import (
	"fmt"
	"log"
	"time"

	"github.com/theredwiking/lanscan/core"
	"github.com/theredwiking/lanscan/models"
)

func Scan(display bool) {
	fmt.Println("Starting scan")

	ips := make(chan string, 254)
	devices := make(chan models.Device, 254)
	
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

	if display {
		go core.OutputDisplay(devices)
	}

	ticker := time.NewTicker(500 * time.Millisecond)

	for {
		select {
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
				return
			}
		}
	}
}

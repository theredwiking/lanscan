package main

import (
	"log"
	"time"

	"github.com/theredwiking/lanscan/core"
	"github.com/theredwiking/lanscan/models"
)

func main() {
	ips := make(chan string, 254)
	alive := make(chan models.Device, 254)
	
	ipRange := []string{"192.168.0.1", "192.168.0.7", "192.168.0.8", "192.168.0.9", "192.168.0.10"}
	
	go core.Ping(ips, alive)
	go core.Ping(ips, alive)
	go core.Ping(ips, alive)

	ips <- ipRange[0]
	ipRange = append(ipRange[:0], ipRange[1:]...)
	ips <- ipRange[0]
	ipRange = append(ipRange[:0], ipRange[1:]...)
	ips <- ipRange[0]
	ipRange = append(ipRange[:0], ipRange[1:]...)

	ticker := time.NewTicker(500 * time.Millisecond)

	for {
		select {
		case msg := <-alive:
			log.Println(msg)
		case _ = <- ticker.C:
			if len(ipRange) != 0 {
				ips <- ipRange[0]
				log.Println(len(ipRange))
				if len(ipRange) >= 2 {
					ipRange = append(ipRange[:0], ipRange[1:]...)
				} else {
					ipRange = []string{}
				}
			}
		}
	}
}

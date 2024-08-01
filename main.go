package main

import (
	"log"

	"github.com/theredwiking/lanscan/core"
)

func main() {
	ips := make(chan string)
	alive := make(chan bool, 254)
	
	ipRange := []string{"192.168.0.1", "192.168.0.7", "192.168.0.8", "192.168.0.9", "192.168.0.10"}
	
	go core.Ping(ips, alive)
	go core.Ping(ips, alive)
	go core.Ping(ips, alive)

	ips <- ipRange[0]
	ipRange = append(ipRange[:0], ipRange[1:]...)

	for {
		select {
		case msg := <-alive:
			log.Println(msg)
			ips <- ipRange[0]
			if len(ipRange) > 1 {
				ipRange = append(ipRange[:0], ipRange[1:]...)
			}
		}
	}
}

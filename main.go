package main

import (
	"log"

	"github.com/theredwiking/lanscan/core"
)

func main() {
	_, err := core.LanIP()
	if err != nil {
		log.Printf("No ip address found: %v", err)
		return
	}
	messages := make(chan core.IP, 3)
	
	pinger := core.NewPinger()
	
	go pinger.Ping(messages)

	pinger.Reciever <- core.IP{Ip: "192.168.0.1", CIDR: 24, Active: false}

	//go core.Ping(messages)

	for {
		select {
		case msg := <-messages:
			log.Println(msg)
			pinger.Reciever <- core.IP{Ip: "192.168.0.7", CIDR: 24, Active: false}
		}
	}
}

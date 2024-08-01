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
	message := make(chan bool)

	go core.Ping("192.168.0.1", message)

	for {
		select {
		case msg := <-message:
			log.Println(msg)
			return
		}
	}
}

package core

import (
	"log"
	"time"

	"github.com/go-ping/ping"
)

func Ping(ip string, message chan bool) {
	pinger, err := ping.NewPinger(ip)
	if err != nil {
		log.Println(err)
		message <- false
		return
	}
	
	pinger.SetPrivileged(true)
	pinger.Count = 2
	pinger.Timeout = 5 * time.Second

	err = pinger.Run()
	if err != nil {
		pinger.Stop()
		log.Println(err)
		message <- false
		return
	}

	stats := pinger.Statistics()

	if stats.PacketLoss == 0 {
		message <- true
	} else {
		message <- false
	}
}

package core

import (
	"log"
	"time"

	"github.com/go-ping/ping"
	"github.com/theredwiking/lanscan/models"
)

// Pings device from string, uses 2 channels
// so as many can be started and used with only 2 channels
func Ping(in chan string, out chan models.Device) {
	for {
		select {
		case ip := <- in:
			device := models.Device{Ip: ip, Active: false, Ports: []int{}}
			pinger, err := ping.NewPinger(ip)
			if err != nil {
				log.Println(err)
				out <- device
				return
			}
			
			pinger.SetPrivileged(true)
			pinger.Count = 2
			pinger.Timeout = 2 * time.Second

			err = pinger.Run()
			if err != nil {
				pinger.Stop()
				log.Println(err)
				out <- device
				return
			}

			stats := pinger.Statistics()

			if stats.PacketLoss == 0 {
				device.Active = true
				out <- device
			} else {
				out <- device
			}
		}
	}
}

package core

import (
	"log"
	"time"

	"github.com/go-ping/ping"
	"github.com/theredwiking/lanscan/models"
)

// Pings device from string, uses 2 channels
// so as many can be started and used with only 2 channels
func Ping(in chan string, out chan bool) {
	for {
		select {
		case ip := <- in:
			device := models.Device{ip, false, []}
			pinger, err := ping.NewPinger(ip)
			if err != nil {
				log.Println(err)
				out <- false
				return
			}
			
			pinger.SetPrivileged(true)
			pinger.Count = 2
			pinger.Timeout = 5 * time.Second

			err = pinger.Run()
			if err != nil {
				pinger.Stop()
				log.Println(err)
				out <- false
				return
			}

			stats := pinger.Statistics()

			if stats.PacketLoss == 0 {
				out <- true
			} else {
				out <- false
			}
		}
	}
}

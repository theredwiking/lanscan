package core

import (
	"log"
	"time"

	"github.com/go-ping/ping"
)

// Ping Struct
type Pinger struct {
	Reciever chan IP
}

// Generates new pinger
func NewPinger() *Pinger {
	return &Pinger{
		Reciever: make(chan IP),
	}
}

// Pings ip provide and returns result through channel
// So it can run as an goroutine
func (p *Pinger) Ping(ips chan IP) {
	for {
		select {
		case ip := <- p.Reciever:
			pinger, err := ping.NewPinger(ip.Ip)
			if err != nil {
				log.Println(err)
				ips <- ip
				return
			}
			
			pinger.SetPrivileged(true)
			pinger.Count = 2
			pinger.Timeout = 5 * time.Second

			err = pinger.Run()
			if err != nil {
				pinger.Stop()
				log.Println(err)
				ips <- ip
				return
			}

			stats := pinger.Statistics()

			if stats.PacketLoss == 0 {
				ip.Active = true
				ips <- ip
			} else {
				ips <- ip
			}
		}
	}
}

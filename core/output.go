package core

import (
	"fmt"

	"github.com/theredwiking/lanscan/models"
)

func OutputDisplay(devices chan models.Device) {
	for {
		select {
		case device := <- devices:
			fmt.Printf("Ip: %s, Active: %t, Ports: %v\n", device.Ip, device.Active, device.Ports)
		}
	}
}

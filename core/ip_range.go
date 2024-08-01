package core

import (
	"errors"
	"fmt"
	"strings"

	"github.com/theredwiking/lanscan/models"
)

// Takes in LAN model and returns possible ips
func IpRange(lan models.LAN) ([]string, error) {
	ipSegments := strings.Split(lan.Ip, ".")
 
	var ips []string

	switch {
	case lan.CIDR == 24:
		for i := range 255 {
			ips = append(ips, fmt.Sprintf("%s.%s.%s.%d", ipSegments[0], ipSegments[1], ipSegments[2], i))
		}
	case lan.CIDR >= 16 :
		for j := range 255 {
			for i := range 255 {
				ips = append(ips, fmt.Sprintf("%s.%s.%d.%d", ipSegments[0], ipSegments[1], j, i))
			}
		}
	case lan.CIDR == 12:
		for k:=16; k < 32; k++ {
			for j := range 255 {
				for i := range 255 {
					ips = append(ips, fmt.Sprintf("%s.%d.%d.%d", ipSegments[0], k, j, i))
				}
			}
		}
	case lan.CIDR >= 8:
		for k := range 255 {
			for j := range 255 {
				for i := range 255 {
					ips = append(ips, fmt.Sprintf("%s.%d.%d.%d", ipSegments[0], k, j, i))
				}
			}
		}
	default:
		return ips, errors.New("unsupported CIDR provide")
	}

	return ips, nil
}

package core

import (
	"errors"
	"net"
	"strconv"
	"strings"
)

type IP struct {
	Ip string `json:"ip"`
	CIDR int `json:"cidr"`
	Active bool `json:"active"`
}

// Finds lan information and returns Lan struct
func LanIP() (IP, error){
	ifaces, err := net.Interfaces()
	if err != nil {
		return IP{"", 0, false}, errors.New("failed to get interfaces")
	}

	for _, iface := range ifaces {
		if iface.Flags&net.FlagUp == 0 {
			continue
		}
		if iface.Flags&net.FlagLoopback != 0 {
			continue
		}

		addrs, err := iface.Addrs()
		if err != nil {
			return IP{"", 0, false}, errors.New("failed to load iface addresses")
		}
		for _, addr := range addrs {
			var ip net.IP
			switch v := addr.(type) {
			case *net.IPNet:
				ip = v.IP
			case *net.IPAddr:
				ip = v.IP
			}
			if ip == nil || ip.IsLoopback() {
				continue
			}
			ip = ip.To4()
			if ip == nil {
				continue
			}
			cidr := getCIDR(addr.String()); if cidr == 0 {
				return IP{"", 0, false}, errors.New("could not get cidr from iface")
			}
			return IP{ip.String(), cidr, true}, nil
		}
	}
	return IP{"", 0, false}, errors.New("no clue");
}

// Takes in addr.String() and return CIDR
func getCIDR(cidr string) int {
	cidrArr := strings.Split(cidr, "/")
	newCidr, err := strconv.Atoi(cidrArr[1])
	if err != nil {
		return 0
	}
	return newCidr
}

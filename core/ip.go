package core

import (
	"errors"
	"net"
	"strconv"
	"strings"
)

type Lan struct {
	Ip string `json:"ip"`
	CIDR int `json:"cidr"`
}

func LanIP() (Lan, error){
	ifaces, err := net.Interfaces()
	if err != nil {
		return Lan{"", 0}, errors.New("failed to get interfaces")
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
			return Lan{"", 0}, errors.New("failed to load iface addresses")
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
				return Lan{"", 0}, errors.New("could not get cidr from iface")
			}
			return Lan{ip.String(), cidr}, nil
		}
	}
	return Lan{"", 0}, errors.New("no clue");
}

func getCIDR(cidr string) int {
	cidrArr := strings.Split(cidr, "/")
	newCidr, err := strconv.Atoi(cidrArr[1])
	if err != nil {
		return 0
	}
	return newCidr
}

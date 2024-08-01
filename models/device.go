package models

type Device struct {
	Ip string `json:"ip"`
	Active bool `json:"active"`
	Ports []int `json:"ports"`
}

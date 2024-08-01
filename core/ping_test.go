package core

import (
	"testing"

	"github.com/attic-labs/testify/assert"
	"github.com/theredwiking/lanscan/models"
)

func TestPingSuccess(t *testing.T) {
	assert := assert.New(t)

	in := make(chan string, 1)
	out := make(chan models.Device, 1)

	in <- "1.1.1.1"

	go Ping(in, out)

	select {
	case msg := <- out:
		assert.Equal(msg, models.Device{Ip: "1.1.1.1", Active: true, Ports: []int{}}, "Checks ip works on 1.1.1.1")
	}
}models.Device{Ip: ip, Active: false, Ports: []int{}}

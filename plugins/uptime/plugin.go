package uptime

// uptime: the simplest useful plugin possible

import (
	"fmt"
	"time"

	"github.com/netflix/hal-9001/hal"
)

var booted time.Time

func init() {
	booted = time.Now()
}

func Register(broker hal.Broker) {
	p := hal.Plugin{
		Name:   "uptime",
		Func:   uptime,
		Regex:  "^!uptime",
		Broker: broker,
	}
	p.Register()
}

func uptime(evt hal.Evt) {
	ut := time.Since(booted)
	evt.Reply(fmt.Sprintf("uptime: %s", ut.String()))
}

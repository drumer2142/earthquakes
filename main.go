package main

import (
	"github.com/drumer2142/earthquakes/alerts"
)

func main() {
	smsSender := alerts.NewSMSSender("6900000000")
	poll := alerts.NewPoller(smsSender)
	poll.Start()
}

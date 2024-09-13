package main

import "github.com/jasonlvhit/gocron"

func BridgeGasTest() {
	s := gocron.NewScheduler()
	_ = s.Every(20).Seconds().From(gocron.NextTick()).Do(CheckBalance)
	<-s.Start()
}

package main

import (
	"github.com/jasonlvhit/gocron"
	"os"
	"strconv"
)

func bridgeGasTest() {
	interval := os.Getenv("Report_Interval")
	num, _ := strconv.ParseInt(interval, 10, 64)
	s := gocron.NewScheduler()
	_ = s.Every(uint64(num)).Seconds().From(gocron.NextTick()).Do(CheckBalance)
	<-s.Start()
}

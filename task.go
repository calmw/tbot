package main

import (
	"github.com/jasonlvhit/gocron"
	"os"
	"strconv"
	"tbot/blockchain"
)

func bridgeGasTest() {
	interval := os.Getenv("REPORT_INTERVAL")
	num, _ := strconv.ParseInt(interval, 10, 64)
	s := gocron.NewScheduler()
	_ = s.Every(uint64(num)).Seconds().From(gocron.NextTick()).Do(blockchain.CheckBalance)
	<-s.Start()
}

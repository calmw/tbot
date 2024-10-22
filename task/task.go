package task

import (
	"github.com/jasonlvhit/gocron"
	"os"
	"strconv"
	"tbot/blockchain"
	"tbot/monitor"
)

func Alarm() {
	interval := os.Getenv("REPORT_INTERVAL")
	num, _ := strconv.ParseInt(interval, 10, 64)
	s := gocron.NewScheduler()
	_ = s.Every(uint64(num)).Seconds().From(gocron.NextTick()).Do(blockchain.CheckBalance)
	_ = s.Every(600).Seconds().From(gocron.NextTick()).Do(monitor.IsNodeMatch714TestnetAlive)
	_ = s.Every(600).Seconds().From(gocron.NextTick()).Do(monitor.IsNodeMatch698TestnetAlive)
	_ = s.Every(600).Seconds().From(gocron.NextTick()).Do(monitor.BridgeFailedOrderCount)
	_ = s.Every(600).Seconds().From(gocron.NextTick()).Do(blockchain.CheckErc20Balance)
	<-s.Start()
}

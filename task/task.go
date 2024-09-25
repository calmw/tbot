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
	_ = s.Every(600).Seconds().From(gocron.NextTick()).Do(monitor.IsNodeBscTestnetAlive)
	_ = s.Every(600).Seconds().From(gocron.NextTick()).Do(monitor.IsNodeMatchTestnetOldAlive)
	_ = s.Every(600).Seconds().From(gocron.NextTick()).Do(monitor.BridgeFailedOrderCount)
	<-s.Start()
}

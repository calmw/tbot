package main

import (
	"tbot/blockchain"
	"tbot/db"
	"tbot/task"
)

func main() {
	db.InitMysql()
	blockchain.NewChains()
	task.Alarm()
}

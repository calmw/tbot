package main

import (
	"tbot/blockchain"
	"tbot/task"
)

func main() {
	blockchain.NewChains()
	task.Alarm()
}

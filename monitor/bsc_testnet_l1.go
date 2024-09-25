package monitor

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/shopspring/decimal"
	"log"
	"math/big"
	"os"
	"tbot/telegram"
	"time"
)

const BSC_TESTNET_L1_RPC = "http://52.195.158.175:8545"

func IsNodeBscTestnetAlive() {
	var err error
	var chainId *big.Int
	var cli *ethclient.Client
	for i := 0; i < 3; i++ {
		cli, err = ethclient.Dial(BSC_TESTNET_L1_RPC)
		if err == nil {
			break
		}
		time.Sleep(time.Second * 2)
	}
	if cli == nil {
		sendAlarm("BSC testnet L1 测试链")
		return
	}

	for i := 0; i < 3; i++ {
		chainId, err = cli.ChainID(context.Background())
		if err == nil {
			break
		}
		time.Sleep(time.Second * 2)
	}
	if chainId == nil {
		sendAlarm("BSC testnet L1 测试链")
		return
	}
	if chainId.Int64() != 714 {
		sendAlarm("BSC testnet L1 测试链")
		return
	}

}

func sendAlarm(chainName string) {
	var bot *telegram.Telegram
	bot, err := telegram.NewTelegram("https://api.telegram.org", os.Getenv("TG_TOKEN"))
	if err != nil {
		log.Fatal(err)
		return
	}
	chatId, err := decimal.NewFromString(os.Getenv("CHAT_ID_DEBUG_GROUP"))
	if err != nil {
		log.Println(err)
		return
	}
	ChatId := 0 - chatId.BigInt().Int64()

	err = bot.SendMessage(fmt.Sprintf("%s异常, 请检查", chainName), ChatId, false)
	if err != nil {
		log.Println(err)
	}
}

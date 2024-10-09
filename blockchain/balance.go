package blockchain

import (
	"context"
	"errors"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/shopspring/decimal"
	"log"
	"os"
	"tbot/telegram"
	"time"
)

var chains = make([]map[string]interface{}, 0)

func NewChains() {
	// BSC 主网
	chains = append(chains, map[string]interface{}{
		"name":      "BSC mainnet",
		"threshold": "100000000000000000", // 0.1
		"rpc":       os.Getenv("RPC_BSC"),
		"accounts":  []string{"0xeDDe26D0638d61daFd5dF7717D10d2646bb46B1A", "0xA064C8397dB01AF331ECE59D5b22C18c5DC50a31"},
	})
	// MATCH 旧主网
	chains = append(chains, map[string]interface{}{
		"name":      "Match mainnet old",
		"threshold": "2000000000000000000", // 2e18
		"rpc":       os.Getenv("RPC_MATCH"),
		"accounts":  []string{"0xeDDe26D0638d61daFd5dF7717D10d2646bb46B1A", "0xA064C8397dB01AF331ECE59D5b22C18c5DC50a31"},
	})
	// MATCH L2 主网
	chains = append(chains, map[string]interface{}{
		"name":      "Match mainnet L2",
		"threshold": "100000000000000000", // 0.1
		"rpc":       os.Getenv("RPC_MATCH_L2"),
		"accounts":  []string{"0xeDDe26D0638d61daFd5dF7717D10d2646bb46B1A", "0xA064C8397dB01AF331ECE59D5b22C18c5DC50a31"},
	})
}

func client(rpc string) (*ethclient.Client, error) {
	var cli *ethclient.Client
	var err error
	for i := 0; i < 20; i++ {
		cli, err = ethclient.Dial(rpc)
		if err == nil {
			return cli, nil
		}
		time.Sleep(time.Second)
	}

	return nil, errors.New("rpc error")
}

func Balance(account, rpc string) (string, error) {
	cli, err := client(rpc)
	if err != nil {
		log.Println(err)
		return "0", err
	}

	for i := 0; i < 20; i++ {
		b, errB := cli.BalanceAt(context.Background(), common.HexToAddress(account), nil)
		if errB == nil {
			return b.String(), nil
		}
		time.Sleep(time.Second)
	}

	return "0", errors.New("get balance error")
}

func CheckBalance() {
	var bot *telegram.Telegram
	var balance1, balance2 string
	var balanceThreshold, b1, b2 decimal.Decimal

	deci18 := decimal.NewFromInt(1e18)
	chatId, err := decimal.NewFromString(os.Getenv("CHAT_ID"))
	if err != nil {
		log.Println(err)
		return
	}
	ChatId := 0 - chatId.BigInt().Int64()

	for _, chain := range chains {
		balanceThreshold, err = decimal.NewFromString(chain["threshold"].(string))
		if err != nil {
			log.Println(err)
			return
		}
		rpc := chain["rpc"].(string)
		chainName := chain["name"].(string)
		accounts := chain["accounts"].([]string)
		balance1, err = Balance(accounts[0], rpc)
		if err != nil {
			log.Println(err)
			continue
		}
		b1, err = decimal.NewFromString(balance1)
		if err != nil {
			log.Println(err)
			return
		}
		balance2, err = Balance(accounts[1], rpc)
		if err != nil {
			log.Println(err)
			continue
		}
		b2, err = decimal.NewFromString(balance2)
		if err != nil {
			log.Println(err)
			return
		}

		bot, err = telegram.NewTelegram("https://api.telegram.org", os.Getenv("TG_TOKEN"))
		if err != nil {
			log.Fatal(err)
			return
		}

		log.Println(fmt.Sprintf("%s, 账户：%s, 当前剩余：%s,阀值:%s", chainName, accounts[0], b1.Div(deci18).String(), balanceThreshold.Div(deci18)))
		log.Println(fmt.Sprintf("%s, 账户：%s, 当前剩余：%s,阀值:%s", chainName, accounts[1], b2.Div(deci18).String(), balanceThreshold.Div(deci18)))

		if b1.Cmp(balanceThreshold) == -1 {
			bal1 := b1.Div(deci18)
			err = bot.SendMessage(fmt.Sprintf("@Abraham_Zero @barlow_node 跨链桥 🛢️gas不足\n\n链 🔗 : %s\n\n账户: %s, 当前剩余: %s, 低于阀值: %s", chainName, accounts[0], bal1.String(), balanceThreshold.Div(deci18)), ChatId, false)
			if err != nil {
				log.Println(err)
			}
		}
		if b2.Cmp(balanceThreshold) == -1 {
			bal2 := b2.Div(deci18)
			err = bot.SendMessage(fmt.Sprintf("@Abraham_Zero @barlow_node 跨链桥 🛢️gas不足\n\n链 🔗 : %s\n\n账户: %s, 当前剩余: %s, 低于阀值: %s", chainName, accounts[1], bal2.String(), balanceThreshold.Div(deci18)), ChatId, false)
			if err != nil {
				log.Println(err)
			}
		}
	}
}

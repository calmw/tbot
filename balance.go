package main

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/shopspring/decimal"
	"log"
	"os"
)

const (
	Account1 = "0xeDDe26D0638d61daFd5dF7717D10d2646bb46B1A"
	Account2 = "0xA064C8397dB01AF331ECE59D5b22C18c5DC50a31"
)

func client() *ethclient.Client {
	cli, err := ethclient.Dial(os.Getenv("RPC"))
	if err != nil {
		panic(err)
	}
	return cli
}

func Balance(account string) string {
	cli := client()
	b, err := cli.BalanceAt(context.Background(), common.HexToAddress(account), nil)
	if err != nil {
		return "0"
	}

	log.Println(fmt.Sprintf("%s balanced at %s", account, b.String()))
	return b.String()
}

func CheckBalance() {
	deci18 := decimal.NewFromInt(1e18)
	balanceThreshold, err := decimal.NewFromString(os.Getenv("BalanceThreshold"))
	if err != nil {
		log.Println(err)
		return
	}
	b1, err := decimal.NewFromString(Balance(Account1))
	if err != nil {
		log.Println(err)
		return
	}
	b2, err := decimal.NewFromString(Balance(Account2))
	if err != nil {
		log.Println(err)
		return
	}

	bot, err := NewTelegram("https://api.telegram.org", os.Getenv("TG_TOKEN"))
	if err != nil {
		log.Fatal(err)
		return
	}

	log.Println(fmt.Sprintf("账户：%s, 当前剩余：%s,阀值:%s", Account1, b1.Div(deci18).String(), balanceThreshold.Div(deci18)))
	log.Println(fmt.Sprintf("账户：%s, 当前剩余：%s,阀值:%s", Account1, b2.Div(deci18).String(), balanceThreshold.Div(deci18)))

	if (b1.Cmp(balanceThreshold) == -1) && (b2.Cmp(balanceThreshold) == -1) {
		bal := b1.Div(deci18)
		err = bot.SendMessage(fmt.Sprintf("@Abraham_Zero 跨链桥 🛢️gas不足\n账户：%s, 当前剩余：%s", Account1, bal.String()), 5222613687, false)
		if err != nil {
			log.Println(err)
		}
		bal = b2.Div(deci18)
		err = bot.SendMessage(fmt.Sprintf("@Abraham_Zero 跨链桥 🛢️gas不足\n账户：%s, 当前剩余：%s", Account2, bal.String()), 5222613687, false)
		if err != nil {
			log.Println(err)
		}
	}
}

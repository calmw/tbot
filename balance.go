package main

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/shopspring/decimal"
	"log"
	"math/big"
	"os"
)

const (
	RPC      = "https://bsc-mainnet.rpcfast.com/?api_key=xbhWBI1Wkguk8SNMu1bvvLurPGLXmgwYeC4S6g2H7WdwFigZSmPWVZRxrskEQwIf"
	Account1 = "0xeDDe26D0638d61daFd5dF7717D10d2646bb46B1A"
	Account2 = "0xA064C8397dB01AF331ECE59D5b22C18c5DC50a31"
)

func client() *ethclient.Client {
	cli, err := ethclient.Dial(RPC)
	if err != nil {
		panic(err)
	}
	return cli
}

func Balance(account string) *big.Int {
	cli := client()
	b, err := cli.BalanceAt(context.Background(), common.HexToAddress(account), nil)
	if err != nil {
		return nil
	}

	log.Println(fmt.Sprintf("%s balanced at %s", account, b.String()))
	return b
}

func CheckBalance() {
	balanceThreshold := big.NewInt(1e16 * 2) //0.02
	//balanceThreshold := big.NewInt(1e18 * 2) //0.02
	bot, err := NewTelegram("https://api.telegram.org", os.Getenv("TG_TOKEN"))
	if err != nil {
		log.Fatal(err)
		return
	}
	b1 := Balance(Account1)
	b2 := Balance(Account2)
	deci18 := decimal.NewFromInt(1e18)
	if b1.Cmp(balanceThreshold) == -1 {
		bal := decimal.NewFromInt(b1.Int64()).Div(deci18)
		err = bot.SendMessage(fmt.Sprintf("@Abraham_Zero 跨链桥 🛢️gas不足\n账户：%s, 当前剩余：%s", Account1, bal.String()), 5222613687, false)
		if err != nil {
			log.Println(err)
		}
	}
	if b2.Cmp(balanceThreshold) == -1 {
		bal := decimal.NewFromInt(b2.Int64()).Div(deci18)
		err = bot.SendMessage(fmt.Sprintf("@Abraham_Zero 跨链桥 🛢️gas不足\n账户：%s, 当前剩余：%s", Account2, bal.String()), 5222613687, false)
		if err != nil {
			log.Println(err)
		}
	}
}

package blockchain

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/shopspring/decimal"
	"log"
	"os"
	"tbot/telegram"
)

const (
	Account1 = "0xeDDe26D0638d61daFd5dF7717D10d2646bb46B1A"
	Account2 = "0xA064C8397dB01AF331ECE59D5b22C18c5DC50a31"
)

var chains = make([]map[string]interface{}, 0)

func NewChains() {
	// BSC 主网
	chains = append(chains, map[string]interface{}{
		"name":      "BSC mainnet",
		"threshold": "20000000000000000", // 2e16
		"rpc":       os.Getenv("RPC_BSC"),
		"accounts":  []string{"0xeDDe26D0638d61daFd5dF7717D10d2646bb46B1A", "0xA064C8397dB01AF331ECE59D5b22C18c5DC50a31"},
	})
	// MATCH 主网
	chains = append(chains, map[string]interface{}{
		"name":      "Match mainnet",
		"threshold": "2000000000000000000", // 2e18
		"rpc":       os.Getenv("RPC_MATCH"),
		"accounts":  []string{"0xeDDe26D0638d61daFd5dF7717D10d2646bb46B1A", "0xA064C8397dB01AF331ECE59D5b22C18c5DC50a31"},
	})
	// MATCH L2 主网
	//chains = append(chains, map[string]interface{}{
	//	"name":      "Match L2 mainnet",
	//	"threshold": "2000000000000000000", // 2e18
	//	"rpc":       os.Getenv("RPC_MATCH_L2"),
	//	"accounts":  []string{"0xeDDe26D0638d61daFd5dF7717D10d2646bb46B1A", "0xA064C8397dB01AF331ECE59D5b22C18c5DC50a31"},
	//})
}

func client(rpc string) *ethclient.Client {
	cli, err := ethclient.Dial(rpc)
	if err != nil {
		panic(err)
	}
	return cli
}

func Balance(account, rpc string) string {
	cli := client(rpc)
	b, err := cli.BalanceAt(context.Background(), common.HexToAddress(account), nil)
	if err != nil {
		return "0"
	}

	return b.String()
}

func CheckBalance() {
	deci18 := decimal.NewFromInt(1e18)
	chatId, err := decimal.NewFromString(os.Getenv("CHAT_ID"))
	fmt.Println(chatId, 8889)
	if err != nil {
		log.Println(err)
		return
	}
	ChatId := 0 - chatId.BigInt().Int64()

	for _, chain := range chains {
		balanceThreshold, err := decimal.NewFromString(chain["threshold"].(string))
		if err != nil {
			log.Println(err)
			return
		}
		chainName := chain["name"].(string)
		accounts := chain["accounts"].([]string)
		b1, err := decimal.NewFromString(Balance(accounts[0], chain["rpc"].(string)))
		if err != nil {
			log.Println(err)
			return
		}
		b2, err := decimal.NewFromString(Balance(accounts[1], chain["rpc"].(string)))
		if err != nil {
			log.Println(err)
			return
		}

		bot, err := telegram.NewTelegram("https://api.telegram.org", os.Getenv("TG_TOKEN"))
		if err != nil {
			log.Fatal(err)
			return
		}

		log.Println(fmt.Sprintf("%s, 账户：%s, 当前剩余：%s,阀值:%s", chainName, accounts[0], b1.Div(deci18).String(), balanceThreshold.Div(deci18)))
		log.Println(fmt.Sprintf("%s, 账户：%s, 当前剩余：%s,阀值:%s", chainName, accounts[1], b2.Div(deci18).String(), balanceThreshold.Div(deci18)))

		if (b1.Cmp(balanceThreshold) == -1) && (b2.Cmp(balanceThreshold) == -1) {
			bal := b1.Div(deci18)
			err = bot.SendMessage(fmt.Sprintf("@Abraham_Zero @barlow_node 跨链桥 🛢️gas不足\n\n链 🔗 : %s\n\n账户: %s, 当前剩余: %s, 低于阀值: %s", chainName, accounts[0], bal.String(), balanceThreshold.Div(deci18)), ChatId, false)
			if err != nil {
				log.Println(err)
			}
			bal = b2.Div(deci18)
			err = bot.SendMessage(fmt.Sprintf("@Abraham_Zero @barlow_node 跨链桥 🛢️gas不足\n\n链 🔗 : %s\n\n账户: %s, 当前剩余: %s, 低于阀值: %s", chainName, accounts[1], bal.String(), balanceThreshold.Div(deci18)), ChatId, false)
			if err != nil {
				log.Println(err)
			}
		}
	}
}

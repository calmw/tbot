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
	erc20 "tbot/blockchain/binding"
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

func Client(rpc string) (*ethclient.Client, error) {
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
	cli, err := Client(rpc)
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

func Erc20Balance(account, token, rpc string) (string, error) {
	cli, err := Client(rpc)
	if err != nil {
		log.Println(err)
		return "0", err
	}
	var tokenObj *erc20.Erc20
	tokenObj, err = erc20.NewErc20(common.HexToAddress(token), cli)
	if err != nil {
		return "", err
	}

	for i := 0; i < 20; i++ {
		b, errB := tokenObj.BalanceOf(nil, common.HexToAddress(account))
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
		// 检查 Match mainnet L2 跨链桥TOX余额
		if chainName == "Match mainnet L2" {
			balanceTox, err := Erc20Balance("0x660406112E2Db5639a2FdD7Fffb21D59aDe69a55", "0x1F294E3B71189dAD7dce235d6FAFBC845C7CD177", rpc)
			if err != nil {
				log.Println(err)
				continue
			}
			bTox, err := decimal.NewFromString(balanceTox)
			if err != nil {
				log.Println(err)
				return
			}
			bToxETH := bTox.Div(deci18)
			eth10000Decimal, err := decimal.NewFromString("10000000000000000000000")
			if err != nil {
				return
			}
			log.Println(fmt.Sprintf("%s, 账户：0x660406112E2Db5639a2FdD7Fffb21D59aDe69a55, 当前剩余：%s TOX,阀值: 10000", chainName, bToxETH))
			if bTox.Cmp(eth10000Decimal) == -1 {
				err = bot.SendMessage(fmt.Sprintf("@Abraham_Zero @barlow_node 跨链桥 TOX不足\n\n链 🔗 : %s\n\n账户: 0x660406112E2Db5639a2FdD7Fffb21D59aDe69a55, 当前剩余: %s, 低于阀值: 10000", chainName, bToxETH.String()), ChatId, false)
				if err != nil {
					log.Println(err)
				}
			}
		} else if chainName == "BSC mainnet" {
			balanceTox, err := Erc20Balance("0x087F0957A8218BA06AdB0465C0aE5B0b57ae0649", "0x837656c3f5858692cCdce13BA66e09d2685073df", rpc)
			if err != nil {
				log.Println(err)
				continue
			}
			bTox, err := decimal.NewFromString(balanceTox)
			if err != nil {
				log.Println(err)
				return
			}
			bToxETH := bTox.Div(deci18)
			eth10000Decimal, err := decimal.NewFromString("100000000000000000000000")
			if err != nil {
				return
			}
			log.Println(fmt.Sprintf("%s, 账户：0x087F0957A8218BA06AdB0465C0aE5B0b57ae0649, 当前剩余：%s TOX,阀值: 100000", chainName, bToxETH))
			if bTox.Cmp(eth10000Decimal) == -1 {
				err = bot.SendMessage(fmt.Sprintf("@Abraham_Zero 跨链桥 TOX不足\n\n链 🔗 : %s\n\n账户: 0x087F0957A8218BA06AdB0465C0aE5B0b57ae0649, 当前剩余: %s, 低于阀值: 100000", chainName, bToxETH.String()), ChatId, false)
				if err != nil {
					log.Println(err)
				}
			}
		}
	}
}

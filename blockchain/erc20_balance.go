package blockchain

import (
	"fmt"
	"github.com/shopspring/decimal"
	"log"
	"os"
	"tbot/db"
	"tbot/model"
	"tbot/telegram"
)

type Bot struct {
	Tg     *telegram.Telegram
	ChatId int64
}

func NewCheckBalanceBot() *Bot {
	chatId, err := decimal.NewFromString(os.Getenv("CHAT_ID_CHECK_BALANCE_GROUP"))
	if err != nil {
		log.Fatal(err)
	}
	cId := 0 - chatId.BigInt().Int64()
	bot, err := telegram.NewTelegram("https://api.telegram.org", os.Getenv("TG_TOKEN"))
	if err != nil {
		log.Fatal(err)
	}

	return &Bot{
		Tg:     bot,
		ChatId: cId,
	}
}

func CheckErc20Balance() {
	var accounts []model.Account
	err := db.Mysql.Model(&model.Account{}).Find(&accounts).Error
	if err != nil {
		log.Println(err)
		return
	}
	log.Println(len(accounts), 1122)
	for _, account := range accounts {
		check(account)
	}
}

func check(account model.Account) {
	var balanceThreshold decimal.Decimal

	bot := NewCheckBalanceBot()

	balance, err := Erc20Balance(account.Account, account.TokenAddress, account.Rpc)
	if err != nil {
		NotifyError(bot, account, err.Error())
		return
	}
	balanceDecimal, err := decimal.NewFromString(balance)
	if err != nil {
		NotifyError(bot, account, err.Error())
		return
	}
	balanceThreshold, err = decimal.NewFromString(account.BalanceThreshold)
	if err != nil {
		log.Println(err)
		return
	}
	deci18 := decimal.NewFromInt(1e18)
	balanceThreshold = balanceThreshold.Div(deci18)
	balanceDecimal = balanceDecimal.Div(deci18)
	log.Println(fmt.Sprintf("账户：%s, 当前剩余：%s,阀值:%s", account.Account, balanceDecimal.String(), balanceThreshold.String()))
	if balanceDecimal.Cmp(balanceThreshold) == -1 {
		NotifyBalance(bot, account, balanceDecimal.String(), balanceThreshold.String())
	}
}

func NotifyError(bot *Bot, account model.Account, errMsg string) {
	err := bot.Tg.SendMessage(fmt.Sprintf("@%s %s %s", account.NotifyWho, account.Title, errMsg), bot.ChatId, false)
	if err != nil {
		log.Println(err)
	}
}

func NotifyBalance(bot *Bot, account model.Account, balance, balanceThreshold string) {
	err := bot.Tg.SendMessage(fmt.Sprintf("@%s %s %s余额不足, %s当前余额: %s, 阀值: %s", account.NotifyWho, account.Title, account.TokenName, account.Account, balance, balanceThreshold), bot.ChatId, false)
	if err != nil {
		log.Println(err)
	}
}

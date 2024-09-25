package monitor

import (
	"encoding/json"
	"fmt"
	"github.com/shopspring/decimal"
	"io"
	"log"
	"net/http"
	"os"
	"tbot/telegram"
)

/// 检控bridge失败订单数量

const MaxFailedOrder = 10

type OrderCount struct {
	Code int `json:"code"`
	Data int `json:"data"`
}

func BridgeFailedOrderCount() {
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

	err, data := GetWithHeader(fmt.Sprintf("http://%s:9002/failed_order_count", os.Getenv("BRIDGE_IP")), nil)
	if err != nil {
		log.Println(err)
		return
	}

	var orderCount OrderCount
	err = json.Unmarshal(data, &orderCount)
	if err != nil || orderCount.Code != 0 {
		log.Println(err)
		return
	}

	log.Println(fmt.Sprintf("跨链桥当前失败订单数量：%d", orderCount.Data))

	if orderCount.Data >= MaxFailedOrder {
		err = bot.SendMessage(fmt.Sprintf("跨链桥当前失败订单数量:%d, 请检查状态", orderCount.Data), ChatId, false)
		if err != nil {
			log.Println(err)
		}
	}

}

func GetWithHeader(url string, headers map[string]string) (error, []byte) {
	client := &http.Client{}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return err, []byte{}
	}
	for key, header := range headers {
		req.Header.Set(key, header)
	}
	resp, err := client.Do(req)
	if err != nil {
		return err, []byte{}
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err, []byte{}
	}
	return nil, body
}

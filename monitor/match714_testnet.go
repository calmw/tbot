package monitor

import (
	"context"
	"crypto/ecdsa"
	"errors"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/shopspring/decimal"
	"log"
	"math/big"
	"os"
	"tbot/telegram"
	"time"
)

const BSC_TESTNET_L1_RPC = "http://52.195.158.175:8545"

func IsNodeMatch714TestnetAlive() {
	var err error
	var chainId *big.Int
	var cli *ethclient.Client
	for i := 0; i < 5; i++ {
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

	for i := 0; i < 5; i++ {
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
	// 转账测试出块节点
	to := common.HexToAddress("0x89876D12A4cB4d19957cEBE3663EA485E05fD3f2")
	err = sendTransactionEvmLtLondon(cli, big.NewInt(714), &to, big.NewInt(1))
	if err != nil {
		sendAlarm("BSC testnet L1 测试链")
		return
	}

}

func sendTransactionEvmLtLondon(cli *ethclient.Client, chainId *big.Int, to *common.Address, value *big.Int) error {
	privateKey, err := crypto.HexToECDSA(os.Getenv("PRIVATE_KEY"))

	if err != nil {
		log.Println(err)
		return err
	}
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		err = errors.New("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
		log.Println(err)
		return err
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	at, err := cli.BalanceAt(context.Background(), fromAddress, nil)
	if err != nil {
		return err
	}
	if at.Int64() == 0 {
		err = errors.New("test account balance is 0")
		log.Println(err)
	}

	nonce, err := cli.NonceAt(context.Background(), fromAddress, nil)
	if err != nil {
		log.Println(err)
		return err
	}

	gasPrice, err := cli.SuggestGasPrice(context.Background())
	if err != nil {
		log.Println(err)
		return err
	}

	auth := types.NewEIP155Signer(chainId)
	tx := types.NewTx(&types.LegacyTx{
		Nonce:    nonce,
		To:       to,
		Value:    value,
		Gas:      210000,
		GasPrice: gasPrice,
		Data:     nil,
	})

	// 签名交易
	signedTx, err := types.SignTx(tx, auth, privateKey)
	if err != nil {
		log.Println(err)
		return err
	}

	// 发送交易
	for i := 0; i < 10; i++ {
		err = cli.SendTransaction(context.Background(), signedTx)
		if err == nil {
			break
		}
		time.Sleep(time.Second * 2)
	}
	if err != nil {
		log.Println(err)
		return err
	}

	var receipt *types.Receipt
	for i := 0; i < 10; i++ {
		receipt, err = cli.TransactionReceipt(context.Background(), signedTx.Hash())
		if err == nil {
			break
		}
		time.Sleep(time.Second * 2)
	}

	if err != nil || receipt == nil {
		log.Println(err, receipt)
		return err
	}

	log.Println(fmt.Sprintf("测试链%d,转账测试成功，Hash:%s,测试账户:%s,余额:%s", chainId.Int64(), receipt.TxHash.String(), fromAddress, at.String()))

	return nil
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

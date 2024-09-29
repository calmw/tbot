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
	"log"
	"math/big"
	"os"
	"time"
)

const MATCH_TESTNET_OLD = "http://35.78.192.225:8545/"

func IsNodeMatchTestnetOldAlive() {
	var err error
	var chainId *big.Int
	var cli *ethclient.Client
	for i := 0; i < 3; i++ {
		cli, err = ethclient.Dial(MATCH_TESTNET_OLD)
		if err == nil {
			break
		}
		time.Sleep(time.Second * 2)
	}
	if cli == nil {
		sendAlarm("Match testnet old 测试链")
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
		sendAlarm("Match testnet old 测试链")
		return
	}
	if chainId.Int64() != 698 {
		sendAlarm("Match testnet old 测试链")
		return
	}

	// 转账测试出块节点
	to := common.HexToAddress("0x89876D12A4cB4d19957cEBE3663EA485E05fD3f2")
	err = sendTransactionEvmEgtLondon(cli, big.NewInt(698), &to, big.NewInt(1))
	if err != nil {
		sendAlarm("Match testnet old 测试链")
		return
	}
}

func sendTransactionEvmEgtLondon(cli *ethclient.Client, chainId *big.Int, to *common.Address, value *big.Int) error {
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

	header, err := cli.HeaderByNumber(context.Background(), nil)
	if err != nil {
		log.Println(err)
		return err
	}

	nonce, err := cli.NonceAt(context.Background(), fromAddress, nil)
	if err != nil {
		log.Println(err)
		return err
	}

	gasTipCap, err := cli.SuggestGasTipCap(context.Background())
	if err != nil {
		log.Println(err)
		return err
	}

	// gasprice = min( MaxPriorityFeePerGas + basefee, MaxFeePerGas )
	txData := &types.DynamicFeeTx{
		ChainID:   chainId,
		Nonce:     nonce,
		To:        to,
		Value:     value,
		Gas:       21000,
		GasFeeCap: header.BaseFee, // 烧掉
		GasTipCap: gasTipCap,      // 给矿工
	}

	tx, err := types.SignNewTx(privateKey, types.LatestSignerForChainID(chainId), txData)
	for i := 0; i < 5; i++ {
		err = cli.SendTransaction(context.Background(), tx)
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
	for i := 0; i < 5; i++ {
		receipt, err = cli.TransactionReceipt(context.Background(), tx.Hash())
		if err == nil {
			break
		}
		time.Sleep(time.Second * 2)
	}

	if err != nil || receipt == nil {
		log.Println(err)
		return err
	}

	log.Println(fmt.Sprintf("测试链%d,转账测试成功，Hash:%s,测试账户:%s,余额:%s", chainId.Int64(), receipt.TxHash.String(), fromAddress, at.String()))

	return nil
}

package main

import (
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/qjpcpu/ethereum/contracts"
	"github.com/qjpcpu/ethereum/key"
	"github.com/urfave/cli"
	"math/big"
	"os"
)

const donate_address = "0xE35f3e2A93322b61e5D8931f806Ff38F4a4F4D88"

func main() {
	app := cli.NewApp()
	app.Name = "canceltx"
	app.Usage = "取消挂起的交易"
	app.Authors = []cli.Author{
		cli.Author{
			Name:  "JasonQu",
			Email: "qjpcpu@gmail.com",
		},
	}
	app.Flags = []cli.Flag{
		cli.Uint64Flag{
			Name:  "nonce",
			Usage: "需要取消的交易nonce",
		},
		cli.StringFlag{
			Name:  "private",
			Usage: "私钥",
		},
		cli.Uint64Flag{
			Name:  "gas",
			Usage: "本次使用多少gas price Gwei(为0则自动计算)",
		},
		cli.Float64Flag{
			Name:  "eth",
			Usage: "捐赠多少ETH",
		},
		cli.Uint64Flag{
			Name:  "finney",
			Usage: "捐赠finney",
		},
		cli.StringFlag{
			Name:  "node",
			Usage: "节点地址",
			Value: "https://api.myetherapi.com/eth",
		},
	}
	app.Action = func(c *cli.Context) error {
		return cancelTx(c)
	}
	if err := app.Run(os.Args); err != nil {
		fmt.Println(err)
	}
}

func cancelTx(c *cli.Context) error {
	node := c.String("node")
	private_str := c.String("private")
	nonce := c.Uint64("nonce")
	gasPrice := c.Uint64("gas")
	if private_str == "" {
		fmt.Println("请指定私钥")
		return nil
	}
	if nonce == 0 {
		fmt.Println("请指定nonce")
		return nil
	}
	pk, err := key.StringToPrivateKey(private_str)
	if err != nil {
		fmt.Println("解析私钥失败", err)
		return err
	}
	keypwd := "123456"
	from, keyjson, err := key.ImportPrivateKey(pk, keypwd, keystore.StandardScryptN, keystore.StandardScryptP)
	if err != nil {
		fmt.Println("导入私钥失败", err)
		return err
	}
	conn, err := ethclient.Dial(node)
	if err != nil {
		fmt.Printf("连接到节点%s失败:%v\n", node, err)
		return err
	}
	fmt.Println("连接到节点:", node)
	var amount *big.Int
	if c.Float64("eth") > 0 {
		amount = new(big.Int)
		new(big.Float).Mul(new(big.Float).SetFloat64(c.Float64("eth")), new(big.Float).SetInt64(1000000000000000000)).Int(amount)
	} else if c.Uint64("finney") > 0 {
		amount = new(big.Int).Mul(new(big.Int).SetUint64(c.Uint64("finney")), big.NewInt(1000000000000000))
	}
	var price *big.Int
	if gasPrice > 0 {
		price = new(big.Int).SetUint64(gasPrice * 1000000000)
	}
	tx, err := contracts.TransferETH(conn, from, common.HexToAddress(donate_address), amount, contracts.SignerFuncOf(string(keyjson), keypwd), nonce, price)
	if err != nil {
		fmt.Printf("取消交易失败:%v\n", err)
		return err
	}
	fmt.Printf(
		`尝试向以太坊提交取消请求:
Hash:     %v
GasPrice: %vGwei
Nonce:    %v
Value:    %v(%.05f eth)`+"\n",
		tx.Hash().Hex(),
		new(big.Int).Div(tx.GasPrice(), big.NewInt(1000000000)),
		tx.Nonce(),
		tx.Value().Uint64(),
		AsEth(tx.Value()),
	)
	return nil
}

func AsEth(num *big.Int) float64 {
	one_eth := big.NewFloat(1000000000000000000)
	f, _ := new(big.Float).Quo(new(big.Float).SetInt(num), one_eth).Float64()
	return f
}

```
NAME:
   canceltx - 取消挂起的交易

USAGE:
   canceltx [global options] command [command options] [arguments...]

VERSION:
   0.0.0

AUTHOR:
   JasonQu <qjpcpu@gmail.com>

COMMANDS:
     help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --nonce value    需要取消的交易nonce (default: 0)
   --private value  私钥
   --gas value      本次出价的GasPrice单位Gwei(为0则自动计算) (default: 0)
   --tx value       交易hash(--nonce和--tx参数任选其一)
   --eth value      捐赠多少ETH(可选) (default: 0)
   --finney value   捐赠finney(可选) (default: 0)
   --node value     节点地址 (default: "https://api.myetherapi.com/eth")
   --help, -h       show help
   --version, -v    print the version
```

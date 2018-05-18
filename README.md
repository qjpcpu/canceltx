# 安装

```
wget https://raw.githubusercontent.com/qjpcpu/canceltx/master/release/canceltx.mac -O canceltx
wget https://raw.githubusercontent.com/qjpcpu/canceltx/master/release/canceltx.linux -O canceltx
```

# 示例

```
canceltx --tx 0x71b6d84be19b1cee9c1a81c48fda97843ace6f47f06c4f1a961288a951f87abf --private xxxxxx  --finney 2
```

# 用法

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

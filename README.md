# 配置环境

## git，cURL，docker,docker-compose

```sh
sudo apt-get install curl
sudo apt-get -y install docker-compose
```

## 安装go

https://go.dev/doc/install

### 注意！！！

要保证在root用户下也有go命令可以执行

方法

```
PATH=$PATH:$<path to go exection>
```



## 安装fabric的镜像 二进制文件

```sh
curl -sSLO https://raw.githubusercontent.com/hyperledger/fabric/main/scripts/install-fabric.sh && chmod +x install-fabric.sh
```

如果命令行错误，可以浏览器打开，下载。

将install-fabric.sh保存到一个新目录下，执行安装程序。

```
sudo ./install-fabric.sh docker samples binary
```

# 启动test-network

```sh
cd fabric-samples/test-network
```

```sh
# 首先关闭一下
sudo ./network.sh down
# up
sudo ./network.sh up
```

检查是否运行起来

```sh
sudo docker ps -a
```

```sh
sudo ./network.sh createChannel
```

创建channel

# 部署chaincode

将我的代码clone下来

https://github.com/yy158775/blockchain-exp

```sh
cd vote-smartcontract
go mod tidy
```

```sh
sudo ./network.sh deployCC -ccn basic -ccp $<dir to vote-smartcontract> -ccl go
```

替换$为vote-smartcontract代码的目录

## 这一步可能会出问题

小心debug



# 运行app

## 关键！！！

代码里面的目录：

```go
const (
   mspID         = "Org1MSP"
   cryptoPath    = "../test-network/organizations/peerOrganizations/org1.example.com"
   certPath      = cryptoPath + "/users/User1@org1.example.com/msp/signcerts/cert.pem"
   keyPath       = cryptoPath + "/users/User1@org1.example.com/msp/keystore/"
   tlsCertPath   = cryptoPath + "/peers/peer0.org1.example.com/tls/ca.crt"
   peerEndpoint  = "localhost:7051"
   gatewayPeer   = "peer0.org1.example.com"
   channelName   = "mychannel"
   chaincodeName = "basic"
)
```

程序里面的目录一定要记得与自己的程序运行目录相匹配。

修改cryptoPath即可

## 运行

```sh
cd vote-app
go mod tidy
```

```sh
go run main.go
```

即可显示输入和输出界面
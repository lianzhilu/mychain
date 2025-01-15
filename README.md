# 使用Go搭建的简易区块链

## 项目启动
在项目根目录下: 
```
make build-my-chain
```
片刻后可在output目录下看到编译成功的可执行文件, 启动项目: 

```
./output/mychain run
```

```
mychain
├── Makefile                         # 项目编译
├── README.md
├── blockchain
│         ├── block.go               # 定义了区块的结构和相关操作   
│         ├── blockchain.go          # 定义了区块链和链上的交易操作
│         └── pow.go                 # 实现pow共识
├── cmd                              # 定义命令行的指令              
│         └── chain.go                            
├── output                           # 存放
│         └── mychain
├── transaction                      # 定义了交易的结构和相关操作
│         ├── io.go
│         └── transaction.go
├── utils                            # 工具包 
│   ├── const.go
│   └── utils.go
├── go.mod
├── go.sum
└── main.go                          # 程序的入口, 通过cobra构建命令行

```
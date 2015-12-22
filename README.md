# SimpleDistributedSys
分布式系统大作业:设计一个分布式处理系统.一个节点(进程)作为Master,负责节点注册,任务分配;其他多个节点
(进程)作为worker,执行客户端发送的请求.系统执行流程为:

1. Master首先启动
2. Worker节点启动后向Master进行"注册";Master维护所有已注册几点的信息(比如IP地址和TCP/UDP端口,服务类型等)
3. 客户端提交任务前首先向Master发送"查询"消息,询问将任务提交给哪个Worker
4. Master根据一定的策略(如轮训)从所有Worker中选择一个worker,比如A,将其信息发送给Client
5. Client向worker A发送任务
6. Worker A完成任务后将结果返回给客户端
7. 计算任务为: 2个数的+, -, *, / . 数据类型支持32位整数\64位整数\浮点数.

## 运行方法

    git clone https://github.com/lawlietxxl/SimpleDistributedSys.git
    export GOPATH=path/to/SimpleDistributedSys
    export GOBIN=path/to/SimpleDistributedSys/bin
    cd path/to/SimpleDistributedSys
    #方法一
    go run src/main/main.go
    #方法二
    go install src/main/main.go
    ./bin/main

## 简单说明
+ mac系统下golang的rpc貌似存在内存泄露问题,在client充分close的情况下,在循环6000+次后就会
有runtime error. ubuntu一切正常.

+ 作业要求2的服务类型没有考虑/作业要求7的类型没有考虑,只是全部用了float64
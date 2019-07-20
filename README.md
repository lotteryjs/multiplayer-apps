# 探索 Golang 高实时 Web 应用 &amp; 轻量级游戏服务端编程

`5G 时代的到临，高实时应用应该会有所爆发……`

### 本笔记📒阅读前置知识

* 感兴趣就好(不懂的知识 `google` 都会给你🆙)
* 基本的服务端知识(最好有 Golang 基础)

本系列教程将以 [Nano](https://github.com/lonng/nano)(Lightweight, facility, high performance golang based game server framework) 项目为探索对象。

* 它是一个轻量级的项目，可以很好的让我们入门 ` Golang 高实时 Web 应用`
* 重量级 MMORPG 框架目前不在此讨论范围内🤣

Server 端源码分析项目：[NanoServer](https://github.com/lonng/nanoserver)

* 大家也可以直接看源码学习，不用管我下面👇的废话🤣（快速应用于自己业务线才是最重要的）

为什么选择 Nano？

* 很多项目很美，但并没有提供一个合适(太大 or 太小)的、完整的、可应用于生产的的业务示例(也许是我没有找到合适的🤦‍♀️)

针对 Apk：[Android逆向破解：Android Killer使用](https://www.jianshu.com/p/61a93a6c0c1b)

由于客户端并没有开源（比较蛋疼🤦‍♀️）

我这边会陆陆续续使用 [Cocos Creator](https://www.cocos.com/creator) 去开发一套完整的客户端并且开源。

### 快速上手 GO & WebSocket

扫盲视频教程：[GO实现千万级WebSocket消息推送服务](https://github.com/owenliang/go-push)

示例源码：[go-push](https://github.com/owenliang/go-push)
* 一些 DevOps 额外话(也许你不需要关注……)
* 关于如何优雅的部署(traffic,k8s,docker,swarm……说实话相关内容还挺多的🤦‍♀️)
* 本地可以采用 [Vagrant + VirtualBox] 或 [Docker Machine + VirtualBox] 快速搭建实验环境,可以参看我的[DevOps](https://github.com/Kirk-Wang/DevOps)仓库

以下文字是针对上述视频的脱敏处理(`感谢小鱼儿大佬提供如此通俗易懂的入门教程`)

#### 什么是推送系统？

![push1](./websocket/images/push1.png)

#### 弹幕系统的技术挑战

技术复杂度：

1 个直播间
* 在线人数：100万
* 发送弹幕：1000条/秒
* 推送频率：100万 * 1000条/秒 = 10亿/秒

N 个直播间
* 推送频率：N * 10亿/秒 (有点可怕了，怕老板钱不够🤣……)



### Nano 基本术语脑图
脑图是根据 [如何构建你的第一个nano应用](https://github.com/lonng/nano/blob/master/docs/get_started_zh_CN.md) 来整理的。

![nano-get-started](image/nano-get-started.png)

### NanoServer 整体架构
`🙅‍♀️🙅‍♂️先不要在此意淫分布式架构，微服务架构，DevOps……☞这里旨在先上手后演变🤣`





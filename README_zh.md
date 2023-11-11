# GoSentinel

[English](README.md) | Chinese

GoSentinel： 一款基于 Go 的综合工具，用于 Linux 进程管理、监控和系统性能分析。它提供集群管理、服务器管理、进程守护、安全审计、用户管理和实时性能分析。


### 优点
- 强大的性能：
    - 基于 Golang 的开发
- 特点
    - 管理集群中的服务器
    - 进程收守护
    - 性能可视化
    - 权限管理
- 丰富的使用场景

## 部署和使用

### 二进制部署
```bash
# 获取 https://github.com/WingYat/GoSentinel/releases/new
$ . /gosentinel --config=./config/config.yaml 
$ nohup . /gosentinel --config=./config/config.yaml > /dev/null 2>&1 &
```

### 用 docker 方式部署
```bash
docker run -d \
  --name gosentinel
  --privileged=true \
  --net=host \
  --p 8080:8080 \
  wingyat/gosentinel:latest
```

### 部署源代码的方式
```bash
$ git clone git@github.com:WingYat/GoSentinel.git
cd GoSentinel
$ make
$ . /gosentinel --config=./config/config.yaml 
$ nohup . /gosentinel --config=./config/config.yaml > /dev/null 2>&1 &
```

## 文档
请访问 [GoSentinel 文档](docs/main.md)。

## 作者。
* [WingYat Cheung](https://github.com/WingYat)

## 许可证
NetShieldControl 采用 [Apache License, Version 2.0](https://www.apache.org/licenses/LICENSE-2.0) 许可。

# GoSentinel
基于golang开发的Linux流程管理工具。

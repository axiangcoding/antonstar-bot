# AntonStar-Bot

![GitHub](https://img.shields.io/github/license/axiangcoding/antonstar-bot) ![GitHub Workflow Status](https://img.shields.io/github/actions/workflow/status/axiangcoding/antonstar-bot/build_docker_image.yml) 

AntonStar-Bot，即安东星机器人（下文简称bot），是一个能在QQ、Kook等社交软件上使用的，具有一些有趣的交互功能的机器人，其设计初衷是为了使用者能便捷地查询游戏《战争雷霆》中的一些数据，包括玩家战绩，游戏数据等。

> “安东星”三个字在中国玩家群体中习惯性代指《战争雷霆》，其本身并不具有贬义

## 使用手册
如果你想要了解这个机器人具备的功能和使用的方式，请看文档：https://www.yuque.com/axiangcoding/anton_star/sfw0d8

## 部署指南

如果你想要在自己的服务器中部署机器人，那么请参考以下的步骤：

- 服务单独部署

  1. 准备环境

     bot需要redis和postgres作为数据存储，cqhttp作为交互客户端，因此请在你的服务器中事先部署好redis，postgres和cqhttp。

     这里就不再累述这些服务的部署方式，如果你的服务器中已有相关的服务，也可直接复用（强烈推荐使用docker方式部署，以便进行维护）

     - Redis：https://redis.io/
     - Postgres：https://www.postgresql.org/
     - Cqhttp：https://docs.go-cqhttp.org/

  2. 部署antonstar-bot

     在完成环境的准备后，你需要部署antonstar-bot服务作为机器人的交互核心。

     - 普通方式：（未提供）

     - Docker方式（推荐）：

       使用docker拉取镜像

       ```
       docker pull axiangcoding/antonstar-bot:latest
       ```

       配置文件的路径为 `/app/config/app.toml`，请根据情况修改配置项

       日志文件的路径为 `/app/logs/`，如果你需要保留日志，请将该路径挂载出来

- docker-compose部署

  1. 下载并安装docker

  2. 下载docker-compose.yml文件

     ```bash
     TODO
     ```

  3. 启动服务

     ```
     TODO
     ```

     



## 开发指南

如果你想要开发机器人，那么请参考以下的内容：

- 环境准备

  bot是使用golang语言进行开发的，因此请准备好golang相关的开发环境，这里不再累述

- 启动必要服务

  要在本地开发bot，请事先启动必要的服务。如果你有docker环境，可以执行

  ```
  docker-compose -f ./services/docker-compose.yml
  ```

  **注意：** 请配置好cqhttp的相关配置项

- 启动bot

  ```bash
  go run ./cmd/app
  ```

## 开源协议

本仓库遵循 [GPL-3.0 license](https://github.com/axiangcoding/antonstar-bot/blob/master/LICENSE) 协议

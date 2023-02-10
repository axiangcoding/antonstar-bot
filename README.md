# AntonStar-Bot 安东星机器人

[![GitHub](https://img.shields.io/github/license/axiangcoding/antonstar-bot)](https://github.com/axiangcoding/antonstar-bot/blob/master/LICENSE) [![Docker Image Version (latest semver)](https://img.shields.io/docker/v/axiangcoding/antonstar-bot?sort=semver)](https://hub.docker.com/r/axiangcoding/antonstar-bot) [![Build docker image](https://github.com/axiangcoding/antonstar-bot/actions/workflows/build_docker_image.yml/badge.svg)](https://github.com/axiangcoding/antonstar-bot/actions/workflows/build_docker_image.yml) [![CodeQL](https://github.com/axiangcoding/antonstar-bot/actions/workflows/codeql.yml/badge.svg)](https://github.com/axiangcoding/antonstar-bot/actions/workflows/codeql.yml) [![Qodana](https://github.com/axiangcoding/antonstar-bot/actions/workflows/qodana.yml/badge.svg)](https://github.com/axiangcoding/antonstar-bot/actions/workflows/qodana.yml) 

AntonStar-Bot，即安东星机器人（下文简称 `asbot`），是一个能在QQ、Kook等社交软件上使用的，具有一些有趣的交互功能的机器人，其设计初衷是为了使用者能便捷地查询游戏《战争雷霆》中的一些数据，包括玩家战绩，游戏数据等

> “安东星”三个字在中国玩家群体中习惯性代指《战争雷霆》，其本身并不具有贬义

## 使用手册
如果你想要了解这个机器人具备的功能和使用的方式，请看文档

https://www.yuque.com/axiangcoding/anton_star/sfw0d8

## 部署指南

如果你想要在自己的服务器中部署机器人，请仔细阅读该栏目下的内容

asbot的运行不仅依赖核心组件，同时也依赖其他独立的服务，依赖如下：

- Redis数据库：https://redis.io/
- Postgres数据库：https://www.postgresql.org/
- Cqhttp客户端：https://docs.go-cqhttp.org/

如果你想要部署的服务器上可以连通上述的依赖服务，或者你更愿意独立部署诸如此类的服务，那么推荐采用“单独部署asbot“，这种方式只会部署asbot的核心组件；

如果你想要一次性部署asbot的全部服务，那么推荐采用”docker-compose完整部署“的方式，这种方式尽可能减少了部署的不确定性；

但是不管以何种方式部署asbot核心，均建议并推荐使用 [Docker](https://www.docker.com/) 以容器化的形式进行部署

### 单独部署asbot

单独部署指的是只部署asbot核心组件，且默认你可以连通依赖服务

1. 安装docker
   在部署前，请根据你的服务器的操作系统，选择合适的docker安装方式。安装文档请看官网，这里不再累述：https://docs.docker.com/desktop/

2. 下载asbot镜像

   下载完docker后，可以使用docker命令行拉取asbot镜像

   ```sh
   docker pull axiangcoding/antonstar-bot:latest
   ```

   镜像推荐采用 `latest` 标签的版本，来保证你的服务永远处在最新的状态。当然，在当前的开发进度下， `latest` 版本的更新无法保证向下兼容，所以你也可以指定诸如 `1.0.0` 之类的标签来使用特定的版本

   镜像端口 `8888` 为服务的端口，可以视情况进行映射

   镜像中有两个路径可以根据需要进行挂载，来持久化数据

   - 配置文件的路径为 `/app/config/app.toml`
   - 日志文件的生成路径为 `/app/logs/`

3. 部署antonstar-bot
   下载完镜像后，便可以部署antonstar-bot了。不过在部署前，请先调整配置文件。你可以点击 [这里](https://github.com/axiangcoding/antonstar-bot/blob/master/api-system/config/app.toml) 查看推荐的生产环境的配置文件，并根据实际情况进行修改。修改完的文件建议以挂载的方式挂载到容器中，这样保证重启容器不会丢失配置
   当然你也可以通过环境变量的形式进行配置项的修改，环境变量以 `AS_XXX` 的形式进行命名，并以特定的规则和配置文件中的配置项一一对应。对应的规则为小写转大写，点号修改为下划线，并添加 ‘AS_’ 前缀。如：app.verison = AS_APP_VERSION
   **注意：环境变量的配置项会覆盖文件配置项，其优先级最高**

   - 修改好配置项后，你可以执行下面的命令进行部署：

     ```sh
     # 将/path/to替换为实际的路径
     docker run -p 8080:8888 -v /path/to/app.toml:/app/config/app.toml -v /path/to/logs/:/app/logs/ --name asbot -d axiangcoding/antonstar-bot:latest
     ```

   - 当然你也可以使用docker compose的形式进行部署：
     docker-compose.yaml

     ```yaml
     version: "3"
     services:
       asbot:
         image: axiangcoding/antonstar-bot:latest
         volumes:
           - "/path/to/app.toml:/app/config/app.toml"
           - "/path/to/logs/:/app/logs/"
         ports:
           - "8080:8888"
     ```

     然后在同级目录下执行下面的命令进行部署

     ```sh
     docker compose up -d
     ```

### docker-compose完整部署

完整部署指的是一次性部署全部所需的服务，在使用 docker compose的前提下，部署的步骤尽可能简化了

1. 创建docker-compose.yml文件
   ```yaml
   version: "3"
   services:
     db:
       image: postgres:15-alpine
       restart: always
       environment:
         # postgres用户名，需要和app.toml中的配置对应
         POSTGRES_USER: user
         # postgres密码，需要和app.toml中的配置对应
         POSTGRES_PASSWORD: password
         # postgres数据库，默认即可
         POSTGRES_DB: anton_star
       # 可视情况是否映射数据库端口到宿主机
       ports:
         - "5432:5432"
       # 挂载数据库文件
       volumes:
         - "/path/to/postgres:/var/lib/postgresql/data"
     redis:
       image: redis:7-alpine
       restart: always
       # 可视情况是否映射端口到宿主机
       ports:
         - "6379:6379"
     cqhttp:
       image: ghcr.io/mrs4s/go-cqhttp:1.0.0-rc4
       # 可视情况是否映射端口到宿主机
       ports:
         - "5700:5700"
       # 挂载数据文件
       volumes:
         - "/path/to/cqhttp:/data"
     asbot:
       image: axiangcoding/antonstar-bot:latest
        # 挂载数据文件
       volumes:
         - "/path/to/app.toml:/app/config/app.toml"
         - "/path/to/logs/:/app/logs/"
       # 可视情况是否映射端口到宿主机
       ports:
         - "8080:8888"
       # 代表该服务需要依赖db和redis
       depends_on:
         - db
         - redis
   ```

2. 修改配置项
   当使用docker compose的方式进行部署时，最关键的是需要配置一些必要的配置项：

   - db（postgres）
     默认情况下建议修改账号和密码即可。

   - redis
     默认情况下无需额外配置。

   - cqhttp
     默认情况下需要调整config.yml到合适的配置。

   - 注意，cqhttp部署时较为繁琐，建议查看官网的相关教程。

     参考文档：https://docs.go-cqhttp.org/guide/docker.html

   - asbot
     注意，app.toml需要在宿主机中事先放好，以免映射错误。docker只会自动创建文件夹类型的宿主机目录

3. 部署完整的服务
   在docker-compose.yaml文件的同级目录下，执行下列的命令

   ```sh
   docker-compose up -d
   ```

### 内存要求

asbot镜像启动后需要约30MB内存，完整服务需要总共约100MB内存（低负载情况下，仅供参考）

## 开发指南

如果你想要开发asbot，那么你需要：

- 具备golang的开发知识
- 遵循golang的开发规范
- 懂得查看和使用各种服务的文档
- 正确向仓库提交PR
- ~~热爱《战争雷霆~~》

## 版本兼容性

通常情况下，形如 v1.0.x 的版本之间是相互兼容的，并确保各项功能无较大出入，并且代码呈稳定结构，形如 v1.x.x 的版本之间可能不向下兼容，功能上可能存在较大差异，并且代码结构可能会发生调整；大版本间相互不兼容

但由于目前正在快速开发中，无法保证每次更新的向下兼容性，请根据实际情况谨慎升级，如果遇到问题欢迎提出ISSUE

如果你提交的PR和当前master分支存在较大的冲突，请拉取最新的master代码用以解决冲突

## 开源协议

本仓库遵循 [GPL-3.0 license](https://github.com/axiangcoding/antonstar-bot/blob/master/LICENSE) 协议

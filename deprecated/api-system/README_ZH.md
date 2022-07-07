# Go Gin Template 

[![Go Report Card](https://goreportcard.com/badge/github.com/axiangcoding/go-gin-template)](https://goreportcard.com/report/github.com/axiangcoding/go-gin-template) [![MIT](https://img.shields.io/badge/license-MIT-green.svg)](./LICENSE) [![codebeat badge](https://codebeat.co/badges/25562f5b-a6ee-4ec8-a93d-97b55cec9a22)](https://codebeat.co/projects/github-com-axiangcoding-go-gin-template-master)

切换语言: 中文|[English](./README.md)

本项目的母语是中文，如果其他语言版本的翻译有不到之处，请提Issue修改，不胜感激

## 介绍

本项目是使用Go语言开发的一个Web应用脚手架，集成了Go语言生态中最常用的一些框架，提供较为轻便的开发体验，特别适合小型项目作为初始模板使用。

本项目的Web框架是 [Gin](https://github.com/gin-gonic/gin)，它是目前最流行也是最快的Web框架。同时也集成了相当多但是又十分常用的特性，具体的列表可以看[特性](./README_ZH.md#特性)

当然，Go是一个开放的语言，Gin也提供了十分简洁的API以供使用，因此本项目结构在内的任何东西都可以根据实际情况进行调整，规范化的结构只是为了团队能够按照统一的风格编码，提供最佳的编码体验

## 如何运行

### 环境

通常情况下，Windows的环境配置会略微复杂，但只要遵循官网的安装指引，这些都不会是问题。接下来，请按照下面的步骤安装必要的环境和工具

1. 请确保你的开发机器上已经安装了Go。安装的方式可以参考 https://golang.org/dl/
2. 本项目采用Make来简化构建，请确保安装了Make。Windows下的安装方式可以参考 http://gnuwin32.sourceforge.net/packages/make.htm
3. 本项目采用了Docker来建立开发环境并构建项目的镜像，请确保安装了Docker。安装的方式可以参考官网 https://www.docker.com/get-started
4. 本项目需要Redis和Mysql作为数据源。我们**强烈建议**您使用项目提供的docker-compose搭建。当然，您也可以自己搭建Redis和Mysql服务

### 配置

需要运行本项目，首先需要了解本项目的配置生效规则。遵循以下规则：

- 如果`config/app.toml`配置文件存在，则读取`config/app.toml`作为配置文件
- 如果`config/app.toml`不存在，则使用默认配置文件`config/default/app.toml`

因此，建议您将`config/default/app.toml`复制到`config/app.toml`，并修改`config/app.toml`，这样可以比对两者的区别。配置文件内容大致如下：

```
[app]
version = "1.0.0"
name = "axiangcoding/go-gin-template"

[app.log]
level = "INFO"
file.enable = false
file.path = "./logs/"

···
```

### 运行

本项目的运行非常简单，请按照以下流程进行操作。注意，标注`（可选）`的内容代表这一种方式是可选项。

1. 如果您还未将本项目的代码拉取到本地，请执行下面的命令：

   ```bash
   $ git clone https://github.com/axiangcoding/go-gin-template.git
   ```

2. （可选）本项目在Github中设置为了`Public Template`，这意味着你可以很方便地复制它，成为你自己项目的初始模板。请在Github项目的首页界面，点击 `Use this template`，并按提示完成自己项目的建立。然后，再执行`git clone`命令，将你自己的项目clone到本地，类似：

   ```bash
   $ git clone https://github.com/your-github-namespace/your-project-name.git
   ```

3. 启动Mysql和Redis服务。如果你希望采用预置的docker-compose快速启动，请确保你的docker正在运行，然后执行下面的命令：

   ```bash
   $ make run-dev-env
   ```

   上述命令是使用make的封装，其等价于下面的命令：

   ```bash
   $ docker-compose -f docker/docker-compose.yaml up -d
   ```


4. （可选）你也可以自己运行一个本机可访问的Mysql和Redis，并将配置文件中的对应配置项修改为正确的值

5. 当一切准备就绪后，请执行下面的命令启动本项目：

   ```bash
   $ make run
   ```

   上述命令同样等价于：

   ```bash
   $ go run ./main.go
   ```

6. 如果一切正常，您将会看到下列的运行信息出现在你的控制台中：

   ```
   [GIN-debug] [WARNING] Running in "debug" mode. Switch to "release" mode in produ
   ction.
    - using env:   export GIN_MODE=release
    - using code:  gin.SetMode(gin.ReleaseMode)
   
   [GIN-debug] GET    /swagger/*any             --> github.com/swaggo/gin-swagger.C
   ustomWrapHandler.func1 (3 handlers)
   [GIN-debug] GET    /api/v1/demo/get          --> github.com/axiangcoding/go-gin-
   template/api/v1.DemoGet (4 handlers)
   [GIN-debug] POST   /api/v1/demo/post         --> github.com/axiangcoding/go-gin-
   template/api/v1.DemoPost (4 handlers)
   [GIN-debug] GET    /api/v1/test/test-log     --> github.com/axiangcoding/go-gin-
   template/api/v1.TestLog (3 handlers)
   [GIN-debug] POST   /api/v1/user/login        --> github.com/axiangcoding/go-gin-
   template/api/v1.UserLogin (3 handlers)
   [GIN-debug] POST   /api/v1/user/register     --> github.com/axiangcoding/go-gin-
   template/api/v1.UserRegister (3 handlers)
   [GIN-debug] POST   /api/v1/user/logout       --> github.com/axiangcoding/go-gin-
   template/api/v1.UserLogout (4 handlers)
   [GIN-debug] GET    /api/v1/system/info       --> github.com/axiangcoding/go-gin-
   template/api/v1.SystemInfo (4 handlers)
   2021-12-20T18:02:53.073+0800    INFO    logging/logging.go:72   database mysql c
   onnected success
   ```

   接着，访问地址：[http://localhost:8888/swagger/index.html](http://localhost:8888/swagger/index.html)，如果看到swagger界面，那代表一切应该正常，可以开始您的新体验啦

当然，并非每次开始都能一帆风顺，如果你遇到了任何问题，并在自己多次尝试解决无果时，欢迎提出Issue寻求帮助

## 特性

- Casbin（权限和授权管理）
- Gin（Web框架）
- Redis（Token管理、缓存）
- JWT（Token鉴权）
- Viper（配置项管理）
- Swagger（开放式API接口文档工具）
- Zap （日志框架）
- Mysql（主要数据库）
- Gorm（ORM库）
- Graceful stop （Web服务优雅暂停）

## 常问的问题

1. Go语言社区中不乏Web框架的脚手架，为什么费尽心思又做了一个呢，这是不是重复造轮子？

  答：相比于常见的Go语言web框架的脚手架，本项目的体量较小，更适合小项目的使用。同时也尽可能做到了不在代码中使用复杂特性，更适合初学者学习。当然，本项目也吸取了优秀开源项目的优点，绝非闭门造车，很多看起来不合理的地方，有可能是为了某个目标的妥协。当然。如果您有更好的方式，我们十分期待您提出Issue，创建Merge Request为本项目添砖加瓦

2. 本项目可以在生产环境中使用吗？

  答：本项目完全可以在生产中使用。但是如果是大型的公司项目，那么我们还是建议您采用更高社区活跃度的脚手架项目作为项目的启动模板。社区活跃度高的项目意味着发现的BUG可以得到及时的帮助，同时团队间对代码结构和规范的接受程度也越高

## 联系我

邮件发送至 <axiangcoding@gmail.com>


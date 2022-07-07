# github.com/axiangcoding/ax-web

[![Go Report Card](https://goreportcard.com/badge/github.com/axiangcoding/go-gin-template)](https://goreportcard.com/report/github.com/axiangcoding/go-gin-template) [![MIT](https://img.shields.io/badge/license-MIT-green.svg)](./LICENSE) [![codebeat badge](https://codebeat.co/badges/25562f5b-a6ee-4ec8-a93d-97b55cec9a22)](https://codebeat.co/projects/github-com-axiangcoding-go-gin-template-master)

切换语言: 中文|[English](./README.md)

***开发者的母语是中文，英文的水平较差，因此英文版本使用了Google翻译。如果英文版本的内容有不到之处，请提Issue修改，不胜感激***

## 介绍

`ax-web`是Go语言开发的一个轻量级Web应用模板，是对一些框架的二次集成，特别适合小型项目开发使用。

`ax-web`使用了 [Gin](https://github.com/gin-gonic/gin) 这个web框架作为基本结构，同时也集成了一些常用的特性，具体的列表可以看[特性](./README_ZH.md#特性)。同时，也通过一些实际的例子规范化了代码的编写，希望能够提供一种统一的编码风格，提供最佳的编码体验

***注意：v1.x版本和v2.x版本互不兼容，请不要再使用v1.x版本，而使用v2.x版本，v2.x版本中代码结构做了更改，更加符合新的go开发规范，同时强化了一些特性，修复了大量的BUG***

## 如何运行

### 环境

`ax-web`的开发环境配置非常简单，只需要几个必要的工具安装妥当即可。接下来，请按照下面的步骤安装必要的环境和工具。（Windows的环境配置会略微复杂，但只要遵循安装指引，这些都不会是问题）

1. 项目采用Go语言作为开发语言，请确保你的开发机器上已经正确安装了Go。安装的方式可以参考 https://golang.org/dl/
2. 项目采用Make来简化构建，请确保正确安装了Make。Windows下的安装方式可以参考 http://gnuwin32.sourceforge.net/packages/make.htm
3. 项目采用了Docker来建立开发环境并构建项目的镜像，请确保正确安装了Docker。安装的方式可以参考官网 https://www.docker.com/get-started
4. 作为开发环境，项目需要Redis和Mysql作为数据源。我们**建议**您使用项目提供的`docker-compose`搭建。当然，您也可以自己搭建Redis和Mysql服务或使用现成的
5. 请确保正确安装了IDE。例如装了必要插件的Vscode或者是Jetbrains Goland

### 配置

需要运行项目，首先需要了解项目的配置项优先级设定。项目的配置优先级从最高到最低如下：

1. 环境变量。例如，`app.version` 对应的是 `AX_APP_VERSION`
2. 配置文件 `/config/app.toml`
3. 配置文件 `/config/default/app.toml`

如果你在本地开发，建议您将`/config/default/app.toml`复制到`/config/app.toml`，并修改`config/app.toml`，这样可以比对两者的区别。配置文件内容大致如下：

```
[app]
version = "1.0.0"
name = "github.com/axiangcoding/ax-web"

[app.log]
level = "INFO"
file.enable = false
file.encoder = "console"
file.path = "./logs/"

···
```

### 运行

需要在开发环境运行项目，请按照以下流程进行操作。注意，标注`（可选）`的内容代表这一种方式是可选项。

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
   GOROOT=D:\Program Files\Go #gosetup
   GOPATH=C:\Users\Administrator\go #gosetup
   "D:\Program Files\Go\bin\go.exe" build -o C:\Users\Administrator\AppData\Local\Temp\GoLand\___go_build_github_com_axiangcoding_go_gin_template.exe github.com/axiangcoding/go-gin-template #gosetup
   C:\Users\Administrator\AppData\Local\Temp\GoLand\___go_build_github_com_axiangcoding_go_gin_template.exe
   2022-05-24T23:58:47.455+0800    INFO    data/data.go:36 Database mysql connected success   
   2022-05-24T23:58:47.464+0800    INFO    data/data.go:54 Auto migrate database table success
   [GIN-debug] [WARNING] Running in "debug" mode. Switch to "release" mode in production.
    - using env:   export GIN_MODE=release
    - using code:  gin.SetMode(gin.ReleaseMode)
   
   [GIN-debug] GET    /api/swagger/*any         --> github.com/swaggo/gin-swagger.CustomWrapHandler.func1 (4 handlers)
   [GIN-debug] GET    /api/v1/demo/get          --> github.com/axiangcoding/go-gin-template/controller/v1.DemoGet (4 handlers) 
   [GIN-debug] POST   /api/v1/demo/post         --> github.com/axiangcoding/go-gin-template/controller/v1.DemoPost (4 handlers)
   [GIN-debug] POST   /api/v1/user/login        --> github.com/axiangcoding/go-gin-template/controller/v1.UserLogin (4 handlers)
   [GIN-debug] POST   /api/v1/user/register     --> github.com/axiangcoding/go-gin-template/controller/v1.UserRegister (4 handlers)
   [GIN-debug] POST   /api/v1/user/me           --> github.com/axiangcoding/go-gin-template/controller/v1.UserMe (4 handlers)
   [GIN-debug] GET    /api/v1/system/info       --> github.com/axiangcoding/go-gin-template/controller/v1.SystemInfo (4 handlers)
   2022-05-24T23:58:48.768+0800    INFO    middleware/logger.go:24 GET /api/swagger/index.html  --> status=200, latency_time=322µs, ip=::1
   2022-05-24T23:58:48.790+0800    INFO    middleware/logger.go:24 GET /api/swagger/swagger-ui.css  --> status=200, latency_time=539.5µs, ip=::1
   2022-05-24T23:58:48.815+0800    INFO    middleware/logger.go:24 GET /api/swagger/swagger-ui-bundle.js  --> status=200, latency_time=1.587ms, ip=::1
   2022-05-24T23:58:48.843+0800    INFO    middleware/logger.go:24 GET /api/swagger/swagger-ui-standalone-preset.js  --> status=200, latency_time=949.8µs, ip=::1
   2022-05-24T23:58:49.140+0800    INFO    middleware/logger.go:24 GET /api/swagger/swagger-ui.css.map  --> status=200, latency_time=1.1284ms, ip=::1
   2022-05-24T23:58:49.143+0800    INFO    middleware/logger.go:24 GET /api/swagger/swagger-ui-standalone-preset.js.map  --> status=200, latency_time=3.3462ms, ip=::1
   2022-05-24T23:58:49.144+0800    INFO    middleware/logger.go:24 GET /api/swagger/swagger-ui-bundle.js.map  --> status=200, latency_time=4.621ms, ip=::1
   2022-05-24T23:58:49.227+0800    INFO    middleware/logger.go:24 GET /api/swagger/doc.json  --> status=200, latency_time=307.5µs, ip=::1
   2022-05-24T23:58:49.235+0800    INFO    middleware/logger.go:24 GET /api/swagger/favicon-32x32.png  --> status=200, latency_time=596.4µs, ip=::1
   
   ```
   
   接着，访问地址：[http://localhost:8888/api/swagger/index.html](http://localhost:8888/api/swagger/index.html)，如果看到swagger界面，那代表一切正常，可以开始您的开发之旅了。

当然，并非每次开始都能一帆风顺，如果你遇到了任何问题，并在自己多次尝试解决无果时，欢迎提出Issue寻求帮助

## 特性

- Gin（Web框架）
- Casbin（权限和授权管理）
- Axth （开发者的另外一个开源项目，简易的登录注册框架）
- Redis（Token管理、缓存）
- Cookie-Session （状态管理）
- Viper（配置项管理）
- Swagger（开放式API接口文档工具）
- Zap （日志框架）
- Mysql（主要数据库）
- Gorm（ORM库）
- Graceful stop （Web服务优雅暂停）

## 常问的问题

1. Go语言社区中不乏Web框架的脚手架，为什么费尽心思又做了一个呢，这是不是重复造轮子？

  答：相比于常见的Go语言web框架的脚手架，本项目的体量较小，更适合小项目的使用。同时也尽可能做到了不在代码中使用复杂特性，更适合初学者学习。当然，本项目也吸取了优秀开源项目的优点，绝非闭门造车。当然。如果您有更好的实现方式，我们十分期待您提出Issue，创建Pull Request为本项目添砖加瓦

2. 本项目可以在生产环境中使用吗？

  答：本项目完全可以在生产中使用。但是如果是大型的公司项目，那么我们还是建议您采用更高社区活跃度的脚手架项目作为项目的启动模板。社区活跃度高的项目意味着发现的BUG可以得到及时的帮助，同时团队间对代码结构和规范的接受程度也越高

## 联系我

邮件发送至 <axiangcoding@gmail.com>


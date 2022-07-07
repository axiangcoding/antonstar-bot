# github.com/axiangcoding/ax-web

[![Go Report Card](https://goreportcard.com/badge/github.com/axiangcoding/go-gin-template)](https://goreportcard.com/report/github.com/axiangcoding/go-gin-template) [![MIT](https://img.shields.io/badge/license-MIT-green.svg)](./LICENSE) [![codebeat badge](https://codebeat.co/badges/25562f5b-a6ee-4ec8-a93d-97b55cec9a22)](https://codebeat.co/projects/github-com-axiangcoding-go-gin-template-master)

Switch language: [中文](./README_ZH.md)|English

***The native language of the developer is Chinese, and the level of English is poor, so the English version uses Google Translate. If the content of the English version is not correct, please submit an Issue to modify it, it would be greatly appreciated***

## Introduce

`ax-web` is a lightweight web application template developed in Go language. It is a secondary integration of some frameworks, especially suitable for small project development.

`ax-web` uses the [Gin](https://github.com/gin-gonic/gin) web framework as the basic structure, and also integrates some commonly used features. For a specific list, see [features]( ./README.md#features). At the same time, it also standardized the writing of the code through some practical examples, hoping to provide a unified coding style and provide the best coding experience

***Note: v1.x version and v2.x version are incompatible with each other, please do not use v1.x version, but use v2.x version, the code structure in v2.x version has been changed, more in line with the new go Developed specifications, enhanced some features, and fixed a lot of bugs***

## How to run

### Environment

The development environment configuration of `ax-web` is very simple, only a few necessary tools need to be installed properly. Next, follow the steps below to install the necessary environment and tools. (The Windows environment configuration will be slightly more complicated, but as long as the installation instructions are followed, these will not be a problem)

1. The project uses Go as the development language, please make sure that Go has been installed correctly on your development machine. For the installation method, please refer to https://golang.org/dl/
2. The project uses Make to simplify the build, please make sure that Make is installed correctly. For the installation method under Windows, please refer to http://gnuwin32.sourceforge.net/packages/make.htm
3. The project uses Docker to build the development environment and build the image of the project. Please make sure that Docker is installed correctly. For the installation method, please refer to the official website https://www.docker.com/get-started
4. As a development environment, the project requires Redis and Mysql as data sources. We **recommend** that you use the `docker-compose` build provided by the project. Of course, you can also build Redis and Mysql services yourself or use ready-made
5. Make sure the IDE is installed correctly. For example Vscode or Jetbrains Goland with the necessary plugins installed

### Configuration

To run a project, you first need to understand the project's configuration item priority settings. The configuration priorities of the projects are as follows, from highest to lowest:

1. Environment variables. For example, `app.version` corresponds to `AX_APP_VERSION`
2. Configuration file `/config/app.toml`
3. Configuration file `/config/default/app.toml`

If you develop locally, it is recommended that you copy `/config/default/app.toml` to `/config/app.toml`, and modify `config/app.toml`, so that you can compare the difference between the two. The content of the configuration file is roughly as follows:

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

### Run

To run the project in the development environment, please follow the steps below. Note that the content marked `(optional)` indicates that this method is optional.

1. If you have not pulled the code of this project locally, please execute the following command:

   ```bash
   $ git clone https://github.com/axiangcoding/go-gin-template.git
   ```

2. (Optional) This project is set to `Public Template` in Github, which means you can easily copy it and become the initial template of your own project. Please click `Use this template` on the home page of the Github project, and follow the prompts to complete the creation of your own project. Then, execute the `git clone` command to clone your own project locally, similar to:

   ```bash
   $ git clone https://github.com/your-github-namespace/your-project-name.git
   ```

3. Start Mysql and Redis services. If you want to use the preset docker-compose to start quickly, please make sure your docker is running, and then execute the following command:

   ```bash
   $ make run-dev-env
   ```

   The above command is a package using make, which is equivalent to the following command:

   ```bash
   $ docker-compose -f docker/docker-compose.yaml up -d
   ```


4. (Optional) You can also run a locally accessible Mysql and Redis by yourself, and modify the corresponding configuration items in the configuration file to the correct value

5. When everything is ready, execute the following command to start the project:

   ```bash
   $ make run
   ```

   The above command is also equivalent to:

   ```bash
   $ go run ./main.go
   ```

6. If everything is normal, you will see the following operating information appear in your console:

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

   Next, visit the address: [http://localhost:8888/api/swagger/index.html](http://localhost:8888/api/swagger/index.html), if you see the swagger interface, it means everything is normal , you can start your development journey.

Of course, not everything will be smooth sailing at the beginning. If you encounter any problems and try to solve them many times to no avail, you are welcome to raise an Issue for help.

## Features

- Gin (web framework)
- Casbin (permission and authorization management)
- Axth (another open source project of the developer, a simple login and registration framework)
- Redis (Token management, cache)
- Cookie-Session (state management)
- Viper (configuration item management)
- Swagger (Open API Interface Documentation Tool)
- Zap (logging framework)
- Mysql (main database)
- Gorm (ORM library)
- Graceful stop (graceful stop of the web service)

## Frequently Asked Questions

1. There is no shortage of scaffolding for web frameworks in the Go language community. Why did you work hard to make another one? Is this a duplication of wheels?

  A: Compared with the scaffolding of the common Go language web framework, this project is smaller in size and is more suitable for small projects. At the same time, it does not use complex features in the code as much as possible, which is more suitable for beginners to learn. Of course, this project also draws on the advantages of excellent open source projects, and is by no means behind closed doors. certainly. If you have a better way to implement it, we are looking forward to your raising an Issue and creating a Pull Request to contribute to this project

2. Can this project be used in a production environment?

  A: This project can be used in production. However, if it is a large-scale company project, then we still recommend that you use a scaffolding project with higher community activity as the start-up template of the project. Projects with high community activity mean that BUGs found can be helped in time, and the team’s acceptance of code structure and specifications is also higher

## Contact me

Email to <axiangcoding@gmail.com>
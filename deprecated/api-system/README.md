# Go Gin Template

[![Go Report Card](https://goreportcard.com/badge/github.com/axiangcoding/go-gin-template)](https://goreportcard.com/report/github.com/axiangcoding/go-gin-template) [![MIT](https://img.shields.io/badge/license-MIT-green.svg)](./LICENSE) [![codebeat badge](https://codebeat.co/badges/25562f5b-a6ee-4ec8-a93d-97b55cec9a22)](https://codebeat.co/projects/github-com-axiangcoding-go-gin-template-master)

Switch language: [中文](./README_ZH.md)|English

The native language of this project is Chinese. If there is something wrong with the translation of other language, please open an Issue for modification. Thank you very much.

## Introduce

This project is a web application scaffold developed using the Go language. It integrates some of the most commonly used frameworks in the Go language ecology and provides a relatively light development experience, which is especially suitable for small projects as an initial template.

The web framework of this project is [Gin](https://github.com/gin-gonic/gin), which is currently the most popular and fastest web framework. At the same time, it also integrates quite a lot of but very commonly used features. For a detailed list, please see [Features](./README_ZH.md#Features)

Of course, Go is an open language, Gin also provides a very concise API for use, so everything including the structure of this project can be adjusted according to the actual situation, the standardized structure is only for the team to be able to code in a unified style To provide the best coding experience

## How to run

### Environment

Under normal circumstances, the environment configuration of Windows will be slightly more complicated, but as long as you follow the installation instructions on the official website, these will not be a problem. Next, please follow the steps below to install the necessary environment and tools

1. Make sure that Go has been installed on your development machine. For the installation method, please refer to https://golang.org/dl/
2. This project uses Make to simplify the build. Please make sure that Make is installed. For the installation method under Windows, please refer to http://gnuwin32.sourceforge.net/packages/make.htm
3. This project uses Docker to establish the development environment and build the project's image. Please make sure that Docker is installed. For the installation method, please refer to the official website https://www.docker.com/get-started
4. This project requires Redis and Mysql as data sources. We **strongly recommend** that you use the docker-compose provided by the project to build. Of course, you can also build Redis and Mysql services by yourself

### Configuration

To run this project, you first need to understand the effective rules of the configuration of this project. Follow the following rules:

-If the `config/app.toml` configuration file exists, read `config/app.toml` as the configuration file
-If `config/app.toml` does not exist, use the default configuration file `config/default/app.toml`

Therefore, it is recommended that you copy `config/default/app.toml` to `config/app.toml` and modify `config/app.toml` so that the difference between the two can be compared. The content of the configuration file is roughly as follows:

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

### Run

The operation of this project is very simple, please follow the procedure below. Note that the content marked with `(optional)` means that this method is optional.

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
   [GIN-debug] [WARNING] Running in "debug" mode. Switch to "release" mode in produ
   ction.
    -using env: export GIN_MODE=release
    -using code: gin.SetMode(gin.ReleaseMode)
   
   [GIN-debug] GET /swagger/*any --> github.com/swaggo/gin-swagger.C
   ustomWrapHandler.func1 (3 handlers)
   [GIN-debug] GET /api/v1/demo/get --> github.com/axiangcoding/go-gin-
   template/api/v1.DemoGet (4 handlers)
   [GIN-debug] POST /api/v1/demo/post --> github.com/axiangcoding/go-gin-
   template/api/v1.DemoPost (4 handlers)
   [GIN-debug] GET /api/v1/test/test-log --> github.com/axiangcoding/go-gin-
   template/api/v1.TestLog (3 handlers)
   [GIN-debug] POST /api/v1/user/login --> github.com/axiangcoding/go-gin-
   template/api/v1.UserLogin (3 handlers)
   [GIN-debug] POST /api/v1/user/register --> github.com/axiangcoding/go-gin-
   template/api/v1.UserRegister (3 handlers)
   [GIN-debug] POST /api/v1/user/logout --> github.com/axiangcoding/go-gin-
   template/api/v1.UserLogout (4 handlers)
   [GIN-debug] GET /api/v1/system/info --> github.com/axiangcoding/go-gin-
   template/api/v1.SystemInfo (4 handlers)
   2021-12-20T18:02:53.073+0800 INFO logging/logging.go:72 database mysql c
   onnected success
   ```

   Next, visit the address: [http://localhost:8888/swagger/index.html](http://localhost:8888/swagger/index.html), if you see the swagger interface, it means everything should be normal and you can start Your new experience

Of course, not everything will be smooth sailing at the beginning. If you encounter any problems and try to solve them many times to no avail, you are welcome to raise an Issue for help.

## Features

- Casbin (permission and authorization management)
- Gin (Web Framework)
- Redis (Token management, cache)
- JWT (Token authentication)
- Viper (Configuration Item Management)
- Swagger (Open API interface documentation tool)
- Zap (Log Framework)
- Mysql (main database)
- Gorm (ORM library)
- Graceful stop (Web service gracefully suspended)

## Frequently Asked Questions

1. There is no shortage of scaffolding for web frameworks in the Go language community. Why did you work hard to make another one? Is this a duplication of wheels?

Answer: Compared with the scaffolding of the common Go language web framework, this project has a smaller size and is more suitable for small projects. At the same time, it does not use complex features in the code as much as possible, which is more suitable for beginners to learn. Of course, this project has also absorbed the advantages of excellent open source projects. It is by no means behind closed doors. Many things that seem unreasonable may be a compromise for a certain goal. Of course. If you have a better way, we look forward to your issue and create a Merge Request to contribute to this project

2. Can this project be used in a production environment?

Answer: This project can be used in production. However, if it is a large-scale company project, then we still recommend that you use a scaffolding project with higher community activity as the start-up template of the project. Projects with high community activity mean that BUGs found can be helped in time, and the team’s acceptance of code structure and specifications is also higher

## Contact me

Email to <axiangcoding@gmail.com>
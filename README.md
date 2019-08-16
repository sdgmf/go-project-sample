
## Prerequisite
1. golang
2. docker、docker-compose
3. [jsonnet](https://jsonnet.org/)
4. make

## Quick Start
```bash
    make docker-compose
```

* **访问接口**： http://localhost:8080/product/1
* **consul**: http://localhost:8500/
* **grafana**: http://localhost:3000/ 
* **jaeger**: http://localhost:16686/search
* **Prometheus**: http://localhost:9090/graph 



## 目的

1. 提供一个完整的go语言项目编程示例
2. 通过示例项目介绍go语言中的编程思想
3. 介绍项目实例遵守的规范

## 示例项目

Github源码[go-project-sample](http://github.com/sdgmf/go-project-sample)

## 包结构

关于golang项目的包结构，Dave Chaney博客《[Five suggestions for setting up a Go project](https://dave.cheney.net/2014/12/01/five-suggestions-for-setting-up-a-go-project)》里讨论了package和command的包设计建议，还有一个社区普遍认可的包结构规范[project-layout](https://github.com/golang-standards/project-layout)。在这两个两篇文章的知道下，结合常见的互联网微服务项目，我又细化了如下的项目结构。

```bash
.
├── api
│   └── proto
├── build
│   ├── details
│   ├── products
│   ├── ratings
│   └── reviews
├── cmd
│   ├── details
│   ├── products
│   ├── ratings
│   └── reviews
├── configs
├── deployments
├── dist
├── internal
│   ├── app
│   │   ├── details
│   │   │   ├── controllers
│   │   │   ├── grpcservers
│   │   │   ├── repositorys
│   │   │   └── services
│   │   ├── products
│   │   │   ├── controllers
│   │   │   ├── grpcclients
│   │   │   └── services
│   │   ├── ratings
│   │   │   ├── controllers
│   │   │   ├── grpcservers
│   │   │   ├── repositorys
│   │   │   └── services
│   │   └── reviews
│   │       ├── controllers
│   │       ├── grpcservers
│   │       ├── repositorys
│   │       └── services
│   └── pkg
│       ├── app
│       ├── config
│       ├── consul
│       ├── database
│       ├── jaeger
│       ├── log
│       ├── models
│       ├── transports
│       │   ├── grpc
│       │   └── http
│       │       └── middlewares
│       │           └── ginprom
│       └── utils
│           └── netutil
├── mocks
└── scripts

53 directories
```

### \cmd

同[project-layout](https://github.com/golang-standards/project-layout)

> "该项目的main方法。 每个应用程序的目录名称应与您要拥有的可执行文件的名称相匹配（例如，/cmd/myapp）。 不要在应用程序目录中放入大量代码。 如果您认为代码可以导入并在其他项目中使用，那么它应该存在于/ pkg目录中。 如果代码不可重用或者您不希望其他人重用它，请将该代码放在/ internal目录中。 你会惊讶于别人会做什么，所以要明确你的意图！ 通常有一个小的main函数可以从/ internal和/ pkg目录中导入和调用代码，而不是其他任何东西。"

### \internal\pkg

同[project-layout](https://github.com/golang-standards/project-layout)

> "私有应用程序和库代码。 这是您不希望其他人在其应用程序或库中导入的代码。 将您的实际应用程序代码放在/internal/app目录（例如/internal/app/myapp）和/internal/ pkg目录中这些应用程序共享的代码（例如/internal/pkg/myprivlib）。"

内部的包采用平铺的方式。

### /internal/pkg/config

加载配置文件,或者从配置中心获取配置和监听配置变动。

### /internal/pkg/database

数据库连接初始化和ORM框架初始化配置。

### /internal/pkg/models

结构体定义。

#### /internal/pkg/repositrys

存储层逻辑代码。

### /internal/pkg/services

领域逻辑层代码。

### /internal/pkg/transport

传输层/控制层逻辑代码

### /internal/app

应用内部代码

### /internal/app/products/controllers

控制层

### /internal/app/products/services

逻辑层

### /internal/app/products/repositrys

存储层

### /internal/app/products/grpcserver

grpc server 实现

### /api

同[project-layout](https://github.com/golang-standards/project-layout)

> OpenAPI/Swagger规范，JSON模式文件，协议定义文件等。

### /scripts

同[project-layout](https://github.com/golang-standards/project-layout)

> sql、部署等脚本

### /build

> Dockerfile

### 

## 分层

MVC、领域模型、ORM 这些都是通过把特定职责的代码拆分到不同的层次对象里，在Java里这些分层概念在各种框架里都有体现(如SSH,SSM等常用框架组合)，并且早已形成了默认的规约，是否还适用go语言吗？答案是肯定的。Martin Fowler在《[企业应用架构模式](https://book.douban.com/subject/1230559/)》就阐述过分层带来的各种好处。

1. 便代码复用,提高代码可维护性.如service的代码可被http协议和grpc协议复用，如果增加thrift协议的接口也很方便。
2. 层次清晰，代码可读性更高。
3. 方便单元测试,单元测试往往因为依赖持久的存储而无法进行,如果持久化代码抽取到单独的对象里，这就变的很简单了.

## 依赖注入

Java 程序员都很熟悉依赖注入和控制翻转这种思想，Spring正式基于依赖注入的思想开发。依赖注入的好处是解耦，对象的组装交给容器来控制(选择需要的实现类、是否单例和初始化).基于依赖注入可以很方便的实现单元测试和提高代码可维护性。

关于Golang依赖注入的讨论《[Dependency Injection in Go](https://blog.drewolson.org/dependency-injection-in-go)》,Golang依赖注入的Package有 Uber的[dig](https://github.com/uber-go/dig),[fx](https://github.com/uber-go/fx),facebook 的 [inject](https://github.com/facebookarchive/inject),google的[wire](https://github.com/google/wire)。dig、fx和inject都是基于反射实现，wire是通过代码生成实现，代码生成的方式是显式的。

本示例通过wire来完成依赖注入. 编写wire.go,wire会根据wire.go生成代码。

``` go
// +build wireinject

package main

import (
    "github.com/google/wire"
    "github.com/zlgwzy/go-project-sample/cmd/app"
    "github.com/zlgwzy/go-project-sample/internal/pkg/config"
    "github.com/zlgwzy/go-project-sample/internal/pkg/database"
    "github.com/zlgwzy/go-project-sample/internal/pkg/log"
    "github.com/zlgwzy/go-project-sample/internal/pkg/services"
    "github.com/zlgwzy/go-project-sample/internal/pkg/repositorys"
    "github.com/zlgwzy/go-project-sample/internal/pkg/transport/http"
    "github.com/zlgwzy/go-project-sample/internal/pkg/transport/grpc"
)

var providerSet = wire.NewSet(
    log.ProviderSet,
    config.ProviderSet,
    database.ProviderSet,
    services.ProviderSet,
    repositorys.ProviderSet,
    http.ProviderSet,
    grpc.ProviderSet,
    app.ProviderSet,
)

func CreateApp(cf string) (*app.App, error) {
    panic(wire.Build(providerSet))
}

```

生成代码

```bash
go get github.com/google/wire/cmd/wire

wire ./...
```

生成后的代码在wire_gen.go

``` go
// Code generated by Wire. DO NOT EDIT.

//go:generate wire
//+build !wireinject

package main

import (
    "github.com/google/wire"
    "github.com/zlgwzy/go-project-sample/cmd/app"
    "github.com/zlgwzy/go-project-sample/internal/pkg/config"
    "github.com/zlgwzy/go-project-sample/internal/pkg/database"
    "github.com/zlgwzy/go-project-sample/internal/pkg/log"
    "github.com/zlgwzy/go-project-sample/internal/pkg/services"
    "github.com/zlgwzy/go-project-sample/internal/pkg/repositorys"
    "github.com/zlgwzy/go-project-sample/internal/pkg/transport/grpc"
    "github.com/zlgwzy/go-project-sample/internal/app/proxy/grpcservers"
    "github.com/zlgwzy/go-project-sample/internal/pkg/transport/http"
    "github.com/zlgwzy/go-project-sample/internal/app/proxy/controllers"
)

// Injectors from wire.go:

func CreateApp(cf string) (*app.App, error) {
    viper, err := config.New(cf)
    if err != nil {
        return nil, err
    }
    options, err := log.NewOptions(viper)
    if err != nil {
        return nil, err
    }
    logger, err := log.New(options)
    if err != nil {
        return nil, err
    }
    httpOptions, err := http.NewOptions(viper)
    if err != nil {
        return nil, err
    }
    databaseOptions, err := database.NewOptions(viper, logger)
    if err != nil {
        return nil, err
    }
    db, err := database.New(databaseOptions)
    if err != nil {
        return nil, err
    }
    productsRepository := repositorys.NewMysqlProductsRepository(logger, db)
    productsService := services.NewProductService(logger, productsRepository)
    productsController := controllers.NewProductsController(logger, productsService)
    initControllers := controllers.CreateInitControllersFn(productsController)
    engine := http.NewRouter(httpOptions, initControllers)
    server, err := http.New(httpOptions, logger, engine)
    if err != nil {
        return nil, err
    }
    grpcOptions, err := grpc.NewOptions(viper)
    if err != nil {
        return nil, err
    }
    productsServer, err := grpcservers.NewProductsServer(logger, productsService)
    if err != nil {
        return nil, err
    }
    initServers := grpcservers.CreateInitServersFn(productsServer)
    grpcServer, err := grpc.New(grpcOptions, logger, initServers)
    if err != nil {
        return nil, err
    }
    appApp, err := app.New(logger, server, grpcServer)
    if err != nil {
        return nil, err
    }
    return appApp, nil
}

// wire.go:

var providerSet = wire.NewSet(log.ProviderSet, config.ProviderSet, database.ProviderSet, services.ProviderSet, repositorys.ProviderSet, http.ProviderSet, grpc.ProviderSet, app.ProviderSet)

```

## 面向接口编程

多态和单元测试必须,比较好理解不再解释。

## 显式编程

Golang的开发推崇这样一种显式编程的思想，显式的初始化、方法调用和错误处理.

1. 尽可能不要使用包级别的全局变量.
2. 尽量不要使用init函数，初始化操作可以在main函数中调用，这样方便阅读代码和控制初始化顺序。
3. 函数都要返回错误，用if err != nil 显式的处理错误.
4. 依赖的参数让调用者去控制（控制翻转的思想),可以看下节依赖注入。

几个大佬都讨论过这个问题，博士Peter的《[A theory of modern Go](http://peter.bourgon.org/blog/2017/06/09/theory-of-modern-go.html)》认为魔法代码的核心是”no package level vars; no func init“.单这也不是绝对。
Dave Cheny在《[go-without-package-scoped-variables](https://dave.cheney.net/2017/06/11/go-without-package-scoped-variables)》做了更详细的说明.

## 打印日志

使用比较多的两个日志库，[logrush](https://github.com/sirupsen/logrus)和[zap](https://github.com/uber-go/zap),个人更喜欢zap。

初始化logger,通过viper加载日志相关配置,lumberjack负责日志切割。

``` go

// Options is log configration struct
type Options struct {
     Filename   string
     MaxSize    int
     MaxBackups int
     MaxAge     int
     Level      string
     Stdout     bool
}

func NewOptions(v *viper.Viper) (*Options, error) {
     var (
          err error
          o   = new(Options)
     )
     if err = v.UnmarshalKey("log", o); err != nil {
          return nil, err
     }

     return o, err
}

// New for init zap log library
func New(o *Options) (*zap.Logger, error) {
     var (
          err    error
          level  = zap.NewAtomicLevel()
          logger *zap.Logger
     )

     err = level.UnmarshalText([]byte(o.Level))
     if err != nil {
          return nil, err
     }

     fw := zapcore.AddSync(&lumberjack.Logger{
          Filename:   o.Filename,
          MaxSize:    o.MaxSize, // megabytes
          MaxBackups: o.MaxBackups,
          MaxAge:     o.MaxAge, // days
     })

     cw := zapcore.Lock(os.Stdout)

     // file core 采用jsonEncoder
     cores := make([]zapcore.Core, 0, 2)
     je := zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig())
     cores = append(cores, zapcore.NewCore(je, fw, level))

     // stdout core 采用 ConsoleEncoder
     if o.Stdout {
          ce := zapcore.NewConsoleEncoder(zap.NewDevelopmentEncoderConfig())
          cores = append(cores, zapcore.NewCore(ce, cw, level))
     }

     core := zapcore.NewTee(cores...)
     logger = zap.New(core)

     zap.ReplaceGlobals(logger)

     return logger, err
}


```

logger应该作为私有变量，这样可以统一添加对象的标示。

``` go

type Object struct {
    logger *zap.Logger
}

// 统一添加标示
func NewObject(logger *zap.Logger){
    return &Object{
        logger:  logger.With(zap.String("type","Object"))
    }

}
```

## 错误处理

错误处理还是看Dave Cheny的博客《[Stack traces and the errors package](https://dave.cheney.net/2016/06/12/stack-traces-and-the-errors-package)》,《[Don’t just check errors, handle them gracefully](https://dave.cheney.net/2016/04/27/dont-just-check-errors-handle-them-gracefully)》。

1. 使用类型判断错误。
2. 包装错误，记录错误的上下文。
3. 使用 pakcage [errors](github.com/pkg/errors)
4. 只处理一次错误，处理错误意味着检查错误值并做出决定。

错误输出示例

``` json
// pc.logger.Error("get product by id error", zap.Error(err))

{
    "level":"error",
    "ts":1564056905.4602501,
    "msg":"get product by id error",
    "error":"product service get product error: get product error[id=2]: record not found",
    "errorVerbose":"record not found get product error[id=2]
github.com/zlgwzy/go-project-sample/internal/pkg/repositorys.(*MysqlProductsRepository).Get
/Users/xxx/code/go/go-project-sample/internal/pkg/repositorys/products.go:29
github.com/zlgwzy/go-project-sample/internal/pkg/services.(*DefaultProductsService).Get
/Users/xxx/code/go/go-project-sample/internal/pkg/services/products.go:27
github.com/zlgwzy/go-project-sample/internal/app/proxy/controllers.(*ProductsController).Get
/Users/xxx/code/go/go-project-sample/internal/app/proxy/controllers/products.go:30
github.com/gin-gonic/gin.(*Context).Next
/Users/xxx/go/pkg/mod/github.com/gin-gonic/gin@v1.4.0/context.go:124
github.com/gin-gonic/gin.RecoveryWithWriter.func1
/Users/xxx/go/pkg/mod/github.com/gin-gonic/gin@v1.4.0/recovery.go:83
github.com/gin-gonic/gin.(*Context).Next
/Users/xxx/go/pkg/mod/github.com/gin-gonic/gin@v1.4.0/context.go:124
github.com/gin-gonic/gin.(*Engine).handleHTTPRequest
/Users/xxx/go/pkg/mod/github.com/gin-gonic/gin@v1.4.0/gin.go:389
github.com/gin-gonic/gin.(*Engine).ServeHTTP
/Users/xxx/go/pkg/mod/github.com/gin-gonic/gin@v1.4.0/gin.go:351
net/http.serverHandler.ServeHTTP
/usr/local/Cellar/go/1.12.6/libexec/src/net/http/server.go:2774
net/http.(*conn).serve
/usr/local/Cellar/go/1.12.6/libexec/src/net/http/server.go:1878
runtime.goexit
/usr/local/Cellar/go/1.12.6/libexec/src/runtime/asm_amd64.s:1337
product service get product error
github.com/zlgwzy/go-project-sample/internal/pkg/services.(*DefaultProductsService).Get
/Users/xxx/code/go/go-project-sample/internal/pkg/services/products.go:28
github.com/zlgwzy/go-project-sample/internal/app/proxy/controllers.(*ProductsController).Get
/Users/xxx/code/go/go-project-sample/internal/app/proxy/controllers/products.go:30
github.com/gin-gonic/gin.(*Context).Next
/Users/xxx/go/pkg/mod/github.com/gin-gonic/gin@v1.4.0/context.go:124
github.com/gin-gonic/gin.RecoveryWithWriter.func1
/Users/xxx/go/pkg/mod/github.com/gin-gonic/gin@v1.4.0/recovery.go:83
github.com/gin-gonic/gin.(*Context).Next
/Users/xxx/go/pkg/mod/github.com/gin-gonic/gin@v1.4.0/context.go:124
github.com/gin-gonic/gin.(*Engine).handleHTTPRequest
/Users/xxx/go/pkg/mod/github.com/gin-gonic/gin@v1.4.0/gin.go:389
github.com/gin-gonic/gin.(*Engine).ServeHTTP
/Users/xxx/go/pkg/mod/github.com/gin-gonic/gin@v1.4.0/gin.go:351
net/http.serverHandler.ServeHTTP
/usr/local/Cellar/go/1.12.6/libexec/src/net/http/server.go:2774
net/http.(*conn).serve
/usr/local/Cellar/go/1.12.6/libexec/src/net/http/server.go:1878
runtime.goexit
/usr/local/Cellar/go/1.12.6/libexec/src/runtime/asm_amd64.s:1337"
}

```

## 单元测试

### 存储层测试

添加repositorys/wire.go 创建要测试的对象，会根据ProviderSet注入合适的依赖。

```go
// +build wireinject

package repositorys

import (
    "github.com/google/wire"
    "github.com/zlgwzy/go-project-sample/internal/pkg/config"
    "github.com/zlgwzy/go-project-sample/internal/pkg/database"
    "github.com/zlgwzy/go-project-sample/internal/pkg/log"
)



var testProviderSet = wire.NewSet(
    log.ProviderSet,
    config.ProviderSet,
    database.ProviderSet,
    ProviderSet,
)

func CreateProductRepository(f string) (ProductsRepository, error) {
    panic(wire.Build(testProviderSet))
}

```

添加repositorys/products_test.go,这里采用表格驱动的方法进行测试,存储层测试会依赖数据库。

``` go
package repositorys

import (
    "flag"
    "github.com/stretchr/testify/assert"
    "testing"
)

var configFile = flag.String("f", "app.yml", "set config file which viper will loading.")

func TestProductsRepository_Get(t *testing.T) {
    flag.Parse()

    sto, err := CreateProductRepository(*configFile)
    if err != nil {
        t.Fatalf("create product Repository error,%+v", err)
    }

    tests := []struct {
        name     string
        id       uint64
        expected bool
    }{
        {"1+1", 1, true},
        {"2+3", 2, false},
        {"4+5", 3, false},
    }

    for _, test := range tests {
        t.Run(test.name, func(t *testing.T) {
            _, err := sto.Get(test.id)

            if test.expected {
                assert.NoError(t, err )
            }else {
                assert.Error(t, err)
            }
        })
    }
}

```

运行测试

``` bash
cd internal/pkg/repositorys
 go test . -v -f ../../../configs/proxy.yml 

=== RUN   TestProductsRepository_Get
2019/07/26 14:29:28 use config file:../../../configs/proxy.yml
2019-07-26T14:29:28.301+0800    INFO    load log options success        {"url": "root:xxx@tcp(127.0.0.1:3306)/shop?charset=utf8&parseTime=True&loc=Local"}
=== RUN   TestProductsRepository_Get/1+1
=== RUN   TestProductsRepository_Get/2+3
=== RUN   TestProductsRepository_Get/4+5
--- PASS: TestProductsRepository_Get (0.04s)
    --- PASS: TestProductsRepository_Get/1+1 (0.00s)
    --- PASS: TestProductsRepository_Get/2+3 (0.00s)
    --- PASS: TestProductsRepository_Get/4+5 (0.00s)
PASS
ok      github.com/zlgwzy/go-project-sample/internal/pkg/repositorys    0.049s
```

###  逻辑层测试

通过mockery自动生成mock对象.

```go
    mockery --all --inpkg
```

添加services/wire.go

```go
// +build wireinject

package services

import (
    "github.com/google/wire"
    "github.com/zlgwzy/go-project-sample/internal/pkg/config"
    "github.com/zlgwzy/go-project-sample/internal/pkg/database"
    "github.com/zlgwzy/go-project-sample/internal/pkg/log"
    "github.com/zlgwzy/go-project-sample/internal/pkg/repositorys"
)

var testProviderSet = wire.NewSet(
    log.ProviderSet,
    config.ProviderSet,
    database.ProviderSet,
    ProviderSet,
)

func CreateProductsService(cf string, sto repositorys.ProductsRepository) (ProductsService, error) {
    panic(wire.Build(testProviderSet))
}

```

编写单元测试services/products_test.go

```go
package services

import (
    "flag"
    "github.com/stretchr/testify/assert"
    "github.com/stretchr/testify/mock"
    "github.com/zlgwzy/go-project-sample/internal/pkg/models"
    "github.com/zlgwzy/go-project-sample/internal/pkg/repositorys"
    "testing"
)

var configFile = flag.String("f", "proxy.yml", "set config file which viper will loading.")

func TestProductsRepository_Get(t *testing.T) {
    flag.Parse()

    sto := new(repositorys.MockProductsRepository)

    sto.On("Get", mock.AnythingOfType("uint64")).Return(func(ID uint64) (p *models.Product) {
        return &models.Product{
            ID: ID,
        }
    }, func(ID uint64) error {
        return nil
    })

    svc, err := CreateProductsService(*configFile, sto)
    if err != nil {
        t.Fatalf("create product serviceerror,%+v", err)
    }

    // 表格驱动测试
    tests := []struct {
        name     string
        id       uint64
        expected uint64
    }{
        {"1+1", 1, 1},
        {"2+3", 2, 2},
        {"4+5", 3, 3},
    }

    for _, test := range tests {
        t.Run(test.name, func(t *testing.T) {
            p, err := svc.Get(test.id)
            if err != nil {
                t.Fatalf("product service get proudct error,%+v", err)
            }

            assert.Equal(t, test.expected, p.ID)
        })
    }
}


```

存储层使用生成的MockProductsRepository，可以直接在用例中定义Mock方法的返回值。

### 控制层测试

添加controllers/products_test.go,利用httptest进行测试

```go

package controllers

import (
    "encoding/json"
    "flag"
    "fmt"
    "github.com/gin-gonic/gin"
    "github.com/stretchr/testify/assert"
    "github.com/stretchr/testify/mock"
    "io/ioutil"
    "net/http/httptest"
    "github.com/zlgwzy/go-project-sample/internal/pkg/models"
    "github.com/zlgwzy/go-project-sample/internal/pkg/repositorys"
    "testing"
)

var r *gin.Engine
var configFile = flag.String("f", "proxy.yml", "set config file which viper will loading.")

func setup() {
    r = gin.New()
}

func TestProductsController_Get(t *testing.T) {
    flag.Parse()
    setup()

    sto := new(repositorys.MockProductsRepository)

    sto.On("Get", mock.AnythingOfType("uint64")).Return(func(ID uint64) (p *models.Product) {
        return &models.Product{
            ID: ID,
        }
    }, func(ID uint64) error {
        return nil
    })

    c, err := CreateProductsController(*configFile, sto)
    if err != nil {
        t.Fatalf("create product serviceerror,%+v", err)
    }

    r.GET("/products/:id", c.Get)

    tests := []struct {
        name     string
        id       uint64
        expected uint64
    }{
        {"1", 1, 1},
        {"2", 2, 2},
        {"3", 3, 3},
    }

    for _, test := range tests {
        t.Run(test.name, func(t *testing.T) {
            uri := fmt.Sprintf("/products/%d", test.id)
            // 构造get请求
            req := httptest.NewRequest("GET", uri, nil)
            // 初始化响应
            w := httptest.NewRecorder()

            // 调用相应的controller接口
            r.ServeHTTP(w, req)

            // 提取响应
            rs := w.Result()
            defer func() {
                _ = rs.Body.Close()
            }()

            // 读取响应body
            body, _ := ioutil.ReadAll(rs.Body)
            p := new(models.Product)
            err := json.Unmarshal(body, p)
            if err != nil {
                t.Errorf("unmarshal response body error:%v", err)
            }

            assert.Equal(t, test.expected, p.ID)
        })
    }

}

```

### grpc测试

添加grpcservers/products_test.go

```go
package grpcservers

import (
    "context"
    "flag"
    "github.com/stretchr/testify/assert"
    "github.com/stretchr/testify/mock"
    "github.com/zlgwzy/go-project-sample/internal/pkg/models"
    "github.com/zlgwzy/go-project-sample/internal/pkg/services"
    "github.com/zlgwzy/go-project-sample/api/proto"
    "testing"
)

var configFile = flag.String("f", "proxy.yml", "set config file which viper will loading.")

func TestProductsService_Get(t *testing.T) {
    flag.Parse()

    service := new(services.MockProductsService)

    service.On("Get", mock.AnythingOfType("uint64")).Return(func(ID uint64) (p *models.Product) {
        return &models.Product{
            ID: ID,
        }
    }, func(ID uint64) error {
        return nil
    })

    server, err := CreateProductsServer(*configFile, service)
    if err != nil {
        t.Fatalf("create product server error,%+v", err)
    }

    // 表格驱动测试
    tests := []struct {
        name     string
        id       uint64
        expected uint64
    }{
        {"1+1", 1, 1},
        {"2+3", 2, 2},
        {"4+5", 3, 3},
    }

    for _, test := range tests {
        t.Run(test.name, func(t *testing.T) {
            req := &proto.ProductGetRequest{
                ID: test.id,
            }
            p, err := server.Get(context.Background(), req)
            if err != nil {
                t.Fatalf("product service get proudct error,%+v", err)
            }

            assert.Equal(t, test.expected, p.ID)
        })
    }

}

```

## Makefile

编写Makefile

``` makefile
.PHONY: run
run:
     go run ./cmd -f cmd/app.yml
.PHONY: wire
wire:
    wire ./...
.PHONY: test
test:
    go test -v ./... -f `pwd`/cmd/app.yml -covermode=count -coverprofile=dist/test/cover.out

.PHONY: build
build:
    GOOS=linux GOARCH="amd64" go build ./cmd -o dist/sample5-linux-amd64
    GOOS=darwin GOARCH="amd64" go build ./cmd -o dist/sample5-darwin-amd64
.PHONY: cover
cover:
    go tool cover -html=dist/test/cover.out
.PHONY: mock
mock:
    mockery --all --inpkg
.PHONY: lint
lint:
    golint ./...
.PHONY: proto
proto:
    protoc -I api/proto ./api/proto/products.proto --go_out=plugins=grpc:api/proto
docker: build
    docker-compose -f docker/sample/docker-compose.yml up
```

1. make run 运行项目
2. make wire 生成依赖注入的代码
3. make mock 生成mock对象
4. make test 运行单元测试
5. cover 查看测试用例覆盖度
6. make build 编译代码
7. make lint 静态代码检查
8. make proto 生成grpc代码
9. make docker 通过docker启动项目,包括依赖的数据

## 框架或库

比较喜欢和常用的几个框架或库

1. [Gin](https://github.com/gin-gonic/gin) MVC库
2. [gorm](https://github.com/jinzhu/gorm) ORM库
3. [viper](https://github.com/spf13/viper) 配置管理库
4. [zap](https://github.com/uber-go/zap) 日志库
5. [grpc](https://github.com/grpc/grpc) RPC库
6. [Cobar](https://github.com/spf13/cobra) Command开发库
7. [Opentracing](https://opentracing.io/) 调用链跟踪
8. [go-prometheus](https://github.com/prometheus/client_golang) 服务监控
9. [wire](https://github.com/google/wire) 依赖注入

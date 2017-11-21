# Go语言web简单配置
---

本次使用的web开发框架是Martini，Martini 是一个非常新的 Go 语言的 Web 框架，使用 Go 的 net/http 接口开发，类似 Sinatra 或者 Flask 之类的框架，也可使用自己的 DB 层、会话管理和模板。这个框架在GitHub上都有中文的解释以及用法，比较容易上手。

其特性如下：

 - 使用非常简单
 - 无侵入设计
 - 可与其他 Go 的包配合工作
 - 超棒的路径匹配和路由
 - 模块化设计，可轻松添加工具
 - 大量很好的处理器和中间件
 - 很棒的开箱即用特性
 - 完全兼容 http.HandlerFunc 接口

首先使用命令`go get`从Github上下载或更新Martini的代码包及其依赖包，其下载的文件会放在`GOPATH/src/github.com`文件夹里面
> go get github.com/go-martini/martini

main.go用的是老师给出的代码：
```
package main

import (
    "os"
    "web/service"
    flag "github.com/spf13/pflag"
)

const (
    //默认8080端口
    PORT string = "8080" 
)

func main() {
    //默认8080端口
    port := os.Getenv("PORT") 
    if len(port) == 0 {
        port = PORT
    }
    //端口号的解析
    pPort := flag.StringP("port", "p", PORT, "PORT for httpd listening")
    flag.Parse()
    if len(*pPort) != 0 {
        port = *pPort
    }
    //启动server
    service.NewServer(port)
}
```

service.go是一个简单的hello world
```
package service

import (
    "github.com/go-martini/martini" 
)

func NewServer(port string) {   
    m := martini.Classic()
    
    //请求处理器
    m.Get("/", func(params martini.Params) string {
        return "hello world"
    })

    m.RunOnAddr(":"+port)   
}
```

编译并运行，运行之后即开始监听8080端口
![1](https://raw.githubusercontent.com/LeungChiHo/CloudGo/master/screenshot/QQ%E6%88%AA%E5%9B%BE20171111132022.png)

接着打开http://localhost:8080，可以看到成功输出hello world
![2](https://raw.githubusercontent.com/LeungChiHo/CloudGo/master/screenshot/QQ%E6%88%AA%E5%9B%BE20171111132056.png)

下面使用 curl 测试
![3](https://raw.githubusercontent.com/LeungChiHo/CloudGo/master/screenshot/QQ%E6%88%AA%E5%9B%BE20171111132144.png)

最后是对 web 执行压力测试，一共请求1000次，50%的请求在30ms内完成，所有请求均在87ms内完成。
![4](https://raw.githubusercontent.com/LeungChiHo/CloudGo/master/screenshot/QQ%E6%88%AA%E5%9B%BE20171111132543.png)





 




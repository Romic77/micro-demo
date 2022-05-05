package main

import (
	"github.com/asim/go-micro/v3/logger"
	"github.com/asim/go-micro/v3/web"
	"net/http"
)

func main() {
	// 创建一个新的服务，服务名为 go.micro.srv.hello
	service := web.NewService(
		// 服务名：go.micro.api.user
		web.Name("go.micro.api.user"),
		// 元信息：使用 http 协议
		web.Metadata(map[string]string{"protocol": "http"}),
		// 服务地址：端口号为 8081
		web.Address(":8081"),
	)

	// 通过命令行参数或者插件初始化服务
	service.Init()

	// /user 路由处理函数
	service.HandleFunc("/user", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(`{"msg":"user api"}`))
	})

	// 运行服务
	if err := service.Run(); err != nil {
		logger.Fatal(err)
	}
}

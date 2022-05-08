package main

import (
	"github.com/asim/go-micro/v3/logger"
	"github.com/asim/go-micro/v3/web"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	// 创建 gin 路由
	r := gin.Default()
	r.GET("/user", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"msg": "user api"})
	})

	// 创建一个新的服务，服务名为 go.micro.srv.hello
	service := web.NewService(
		// 服务名：go.micro.api.user
		web.Name("go.micro.api.user"),
		// 元信息：http 协议
		web.Metadata(map[string]string{"protocol": "http"}),
		// 服务地址
		web.Address(":8081"),
		// 使用gin路由
		web.Handler(r),
	)

	// 通过命令行参数或者插件初始化服务
	service.Init()

	// 运行服务
	if err := service.Run(); err != nil {
		logger.Fatal(err)
	}
}

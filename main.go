// @title Go Music API
// @version 1.0
// @description 这是一个基于底层库构建的跨平台音乐搜索与解析统一 API 服务。
// @host localhost:8080
// @BasePath /
package main

import (
	"fmt"
	"os"

	"github.com/guohuiyuan/go-music-api/router"
	"github.com/guohuiyuan/go-music-api/service"
)

func main() {
	service.CM.Load()
	fmt.Println("Cookies 已加载")

	r := router.SetupRouter()

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	fmt.Printf("Music API Server is running on http://localhost:%s\n", port)
	fmt.Printf("Swagger API 接口文档请访问: http://localhost:%s/swagger/index.html\n", port)
	if err := r.Run(":" + port); err != nil {
		panic("Failed to start server: " + err.Error())
	}
}

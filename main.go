package main

import (
	"github.com/gin-gonic/gin"
	"wang/controllers"
)

func main() {
	r := gin.Default()
	// 定义一个路径为 /ping 的 GET 格式路由，并返回 JSON 数据
	r.GET("/ping", controllers.GetUser)
	
	// 路由定义post请求, url路径为：/user/login, 绑定doLogin控制器函数
	r.POST("/user/login", controllers.DoLogin)

	//定义put请求
	r.PUT("/users/:id", controllers.UpdateUser)

	//定义delete请求
	r.DELETE("/users/:id", controllers.DeleteUser)

	r.Run(":8080") // 启动服务，并监听 8080 端口
}
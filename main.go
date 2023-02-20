package main

import (
	"github.com/gin-gonic/gin"
	"wang/controllers"
)

func main() {
	r := gin.Default()

	// 创建v1组
	v1 := r.Group("/v1")
	{
        // 在v1这个分组下，注册路由
		// 定义一个路径为 /ping 的 GET 格式路由，并返回 JSON 数据
		v1.GET("/ping", controllers.GetUser)

		// 路由定义post请求, url路径为：/user/login, 绑定doLogin控制器函数
		v1.POST("/user/login", controllers.DoLogin)

		//定义put请求
		v1.PUT("/users/:id", controllers.UpdateUser)

		//定义delete请求
		v1.DELETE("/users/:id", controllers.DeleteUser)
	}
	r.Run(":8080") // 启动服务，并监听 8080 端口
}
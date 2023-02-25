package main

import (
	"github.com/gin-gonic/gin"
	"wang/controllers"
)

func main() {
	// 学习框架链接: https://www.tizi365.com/archives/244.html
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

	// 创建v2组
	v2 := r.Group("/v2")
	{
		//获取name参数, 通过Query获取的参数值是String类型。
		v2.GET("/users/query", controllers.Query)
		v2.GET("/users/defaultQuery", controllers.DefaultQuery)
		v2.GET("/users/getQuery", controllers.GetQuery)
		v2.GET("/users/handlejson", controllers.HandlerJson)
		v2.GET("/users/handlexml", controllers.HandlelXml)
		v2.GET("/users/handlefile", controllers.HandlelFile)
		v2.GET("/users/handlefileattachment", controllers.HandlelFileAttachment)
		v2.GET("/users/handlehttp", controllers.HandlelHttp)
	}
	r.Run(":8080") // 启动服务，并监听 8080 端口
}
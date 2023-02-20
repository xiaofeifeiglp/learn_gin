package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	// 定义一个路径为 /ping 的 GET 格式路由，并返回 JSON 数据
	r.GET("/ping", getUser)
	
	// 路由定义post请求, url路径为：/user/login, 绑定doLogin控制器函数
	r.POST("/user/login", doLogin)

	//定义put请求
	r.PUT("/users/:id", updateUser)

	//定义delete请求
	r.DELETE("/users/:id", deleteUser)

	r.Run(":8080") // 启动服务，并监听 8080 端口
}

// 控制器函数
func getUser(c *gin.Context) {
	// 获取get请求参数
	c.JSON(200, gin.H{
		"message": "Hello, Gin by quanxiaoha.com !",
	})
}

// 控制器函数
func doLogin(c *gin.Context) {
	// 获取post请求参数
username := c.PostForm("username")
password := c.PostForm("password")
// 通过请求上下文对象Context, 直接往客户端返回一个字符串
c.String(200, "username=%s,password=%s", username,password)
}

func updateUser(c *gin.Context) {
    
}

func deleteUser(c *gin.Context) {
    
}
package controllers

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

// User 结构体定义
type User struct {
	Name  string `json:"name" form:"name"`
	Email string `json:"email" form:"email"`
  }

// 控制器函数
func GetUser(c *gin.Context) {
	// 获取get请求参数
	ip := c.ClientIP()
	fmt.Printf("ip: %v\n", ip)
	c.JSON(200, gin.H{
		"message": "Hello, Gin by quanxiaoha.com !",
	})
}

// 控制器函数
func DoLogin(c *gin.Context) {
	// 获取post请求参数
	username := c.PostForm("username")
	password := c.PostForm("password")
	
	// 初始化user struct
	u := User{}
	// 通过ShouldBind函数，将请求参数绑定到struct对象， 处理json请求代码是一样的。
	// 如果是post请求则根据Content-Type判断，接收的是json数据，还是普通的http请求参数
	fmt.Printf("u: %v\n", u)
	fmt.Printf("c.ShouldBind(&u): %v\n", c.ShouldBind(&u))
	if c.ShouldBind(&u) == nil {
		// 绑定成功， 打印请求参数
		fmt.Println(u.Name)
		fmt.Println(u.Email)
    }

	// 通过请求上下文对象Context, 直接往客户端返回一个字符串
	c.String(200, "username=%s,password=%s", username,password)
}

func UpdateUser(c *gin.Context) {
    id := c.Param("id")
	fmt.Printf(id)
}

func DeleteUser(c *gin.Context) {
    
}


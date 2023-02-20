package controllers

import (
	"github.com/gin-gonic/gin"
)

// 控制器函数
func GetUser(c *gin.Context) {
	// 获取get请求参数
	c.JSON(200, gin.H{
		"message": "Hello, Gin by quanxiaoha.com !",
	})
}

// 控制器函数
func DoLogin(c *gin.Context) {
	// 获取post请求参数
username := c.PostForm("username")
password := c.PostForm("password")
// 通过请求上下文对象Context, 直接往客户端返回一个字符串
c.String(200, "username=%s,password=%s", username,password)
}

func UpdateUser(c *gin.Context) {
    
}

func DeleteUser(c *gin.Context) {
    
}
package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

//获取name参数, 通过Query获取的参数值是String类型
func Query(c *gin.Context) {
	fmt.Printf("c.Query(\"name\"): %v\n", c.Query("name"))
}

//获取name参数, 跟Query函数的区别是，可以通过第二个参数设置默认值
func DefaultQuery(c *gin.Context) {
	fmt.Printf("c.DefaultQuery(\"name\", \"tizi365\"): %v\n", c.DefaultQuery("name", "tizi365"))
}

//获取id参数, 通过GetQuery获取的参数值也是String类型
func GetQuery(c *gin.Context) {
	dd, ok := c.GetQuery("id")
	fmt.Print(dd, ok)
}
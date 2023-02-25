package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

// UserJson 定义
type UserJson struct {
	Name  string `json:"name"` // 通过json标签定义struct字段转换成json字段的名字。
	Email string `json:"email"`
  }

// Handler_json 控制器
func HandlerJson(c *gin.Context) {
	//初始化user对象
	u := &User{
	  Name:  "tizi365",
	  Email: "tizi@tizi365.com",
	}
	//返回json数据
	//返回结果：{"name":"tizi365", "email":"tizi@tizi365.com"}
	c.JSON(200, u)
}

// UserXml 定义, 默认struct的名字就是xml的根节点名字，这里转换成xml后根节点的名字为User.
type UserXml struct {
	Name  string `xml:"name"` // 通过xml标签定义struct字段转换成xml字段的名字。
	Email string `xml:"email"`
}

// Handler_xml 控制器
func HandlelXml(c *gin.Context){
	//初始化user对象
	u := &User{
	  Name:  "tizi365",
	  Email: "tizi@tizi365.com",
	}
	//返回xml数据
	//返回结果：
	//  <?xml version="1.0" encoding="UTF-8"?>
	//  <User><name>tizi365</name><email>tizi@tizi365.com</email></User>
	c.XML(200, u)
}


func HandlelFile(c *gin.Context) {
	//通过File函数，直接返回本地文件，参数为本地文件地址。
	//函数说明：c.File("文件路径")
	c.File("controllers/wangfei.png")
}

func HandlelFileAttachment(c *gin.Context) {
	//通过FileAttachment函数，返回本地文件，类似File函数，区别是可以指定下载的文件名。
	//函数说明: c.FileAttachment("文件路径", "下载的文件名")
	c.FileAttachment("controllers/wangfei.png", "wangfei.png")
}


func HandlelHttp(c *gin.Context) {
	//设置http响应 header, key/value方式，支持设置多个header
	c.Header("site","tizi365")
	fmt.Println(11111, c.Request.Header)
	fmt.Println(11111, c.Request)
}


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
	// 获取url参数id
	id := c.Param("id")
	fmt.Print(id)
	dd, ok := c.GetQuery("id")
	fmt.Print(dd, ok)
}
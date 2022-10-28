package main

import (
	"collect_manage/Helpers"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"github.com/syyongx/php2go"
)

var SYS_TOOLS = Helpers.NewSystemTools()

func init() {

}

func main() {
	SYS_TOOLS.Logger.WithFields(log.Fields{
		"animal": "walrus",
	}).Info("System init success")

	// 典型读取操作，默认分区可以使用空字符串表示
	fmt.Println("App Mode:", SYS_TOOLS.Conf.Section("").Key("app_mode").String())
	fmt.Println("Data Path:", SYS_TOOLS.Conf.Section("paths").Key("data").String())

	SYS_TOOLS.Funcs.Hello()

	SYS_TOOLS.Logger.Info(SYS_TOOLS.Funcs.Md5("str"))

	SYS_TOOLS.Logger.Info(php2go.Abs(-1))

	//1.创建路由
	r := gin.Default()
	//2.绑定路由规则，执行的函数
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status":  "geted",
			"message": "Hello World!",
		})
	})

	r.POST("/form_post", func(c *gin.Context) {
		message := c.PostForm("message")
		nick := c.DefaultPostForm("nick", "anonymous")

		c.JSON(http.StatusOK, gin.H{
			"status":  "posted",
			"message": message,
			"nick":    nick,
		})
	})

	//3.监听端口，默认8080
	r.Run(":8080")
}

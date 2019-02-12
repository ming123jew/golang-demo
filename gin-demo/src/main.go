package main

import (
	"flag"
	"fmt"
	log "github.com/Sirupsen/logrus"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
	"os"
	"rpc-comment/config"
	"rpc-comment/controller"
	"strconv"
)

type conf struct {
	Host string
	Port int
	Pid  string
}

var DB = make(map[string]string)

/**
 *运行配置
 * 命令行执行 go run main.go -host 127.0.0.1 -port 8802 -pid ./run.pid
 * 参数可省略，省略则使用默认参数
 */
func (envConf *conf) getConf() *conf {

	pid := os.Getpid()
	//log.Info("pid: ",pid)

	//服务IP，接受传参，如 “cmd -host 121.1.41.201”，不传就用默认IP
	host := flag.String("host", config.Host, "HOST")
	//服务端口，接受传参，如 “cmd -port 8001”，不传就用默认端口
	port := flag.Int("port", int(config.Port), "PORT")
	//pid文件，接受传参，如 “cmd -pid /tmp/ggo.pid”，不传就用默认文件
	pidPath := flag.String("pid", config.Pid, "PID")
	flag.Parse()

	//log.Info("pidPath: ", *pidPath)

	err := ioutil.WriteFile(*pidPath, []byte(strconv.Itoa(pid)), 0644)
	if err != nil {
		log.Fatal("error write pid to config file")
	}

	envConf.Port = *port
	envConf.Host = *host
	envConf.Pid = *pidPath

	return envConf
}

func setupRouter() *gin.Engine {
	// Disable Console Color
	// gin.DisableConsoleColor()
	r := gin.Default()

	// Ping test
	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	// Get user value
	r.GET("/user/:name", func(c *gin.Context) {
		user := c.Params.ByName("name")
		value, ok := DB[user]
		if ok {
			c.JSON(http.StatusOK, gin.H{"user": user, "value": value})
		} else {
			c.JSON(http.StatusOK, gin.H{"user": user, "status": "no value"})
		}
	})

	// Authorized group (uses gin.BasicAuth() middleware)
	// Same than:
	// authorized := r.Group("/")
	// authorized.Use(gin.BasicAuth(gin.Credentials{
	//	  "foo":  "bar",
	//	  "manu": "123",
	//}))
	authorized := r.Group("/", gin.BasicAuth(gin.Accounts{
		"foo":  "bar", // user:foo password:bar
		"manu": "123", // user:manu password:123
	}))

	authorized.POST("admin", func(c *gin.Context) {
		user := c.MustGet(gin.AuthUserKey).(string)

		// Parse JSON
		var json struct {
			Value string `json:"value" binding:"required"`
		}

		if c.Bind(&json) == nil {
			DB[user] = json.Value
			c.JSON(http.StatusOK, gin.H{"status": "ok"})
		}
	})

	example := controller.Example{}

	//测试Mysql
	r.GET("/example/mysql", example.HandleMysql)

	//测试Redis
	r.GET("/example/redis", example.HandleRedis)

	//测试Mongo
	r.GET("/example/mongo", example.HandleMongo)

	return r
}

func main() {
	//设置路由
	r := setupRouter()

	//加载global配置文件里的配置
	var envConf conf
	envConf.getConf()

	// Listen and Server in 0.0.0.0:8080
	r.Run(fmt.Sprintf("%s:%d", envConf.Host, envConf.Port))
}

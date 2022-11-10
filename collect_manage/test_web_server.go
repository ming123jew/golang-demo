package main

import (
	"collect_manage/Helpers"
	"encoding/json"
	"fmt"
	"net/http"
	"sync"
	"sync/atomic"
	"time"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"github.com/syyongx/php2go"
)

var SYS_TOOLS = Helpers.NewSystemTools()

func init() {

}

func mutexAdd() {
	var a int32 = 0
	var wg sync.WaitGroup
	var mu sync.Mutex
	start := time.Now()
	for i := 0; i < 100000000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			mu.Lock()
			a += 1
			mu.Unlock()
		}()
	}
	wg.Wait()
	timeSpends := time.Now().Sub(start).Nanoseconds()
	fmt.Printf("use mutex a is %d, spend time: %v\n", a, timeSpends)
}

func AtomicAdd() {
	var a int32 = 0
	var wg sync.WaitGroup
	start := time.Now()
	for i := 0; i < 100000000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			atomic.AddInt32(&a, 1)
		}()
	}
	wg.Wait()
	timeSpends := time.Now().Sub(start).Nanoseconds()
	fmt.Printf("use atomic a is %d, spend time: %v\n", atomic.LoadInt32(&a), timeSpends)
}

type DataItem struct {
	ArticleId          string `json:"articleId"`
	ArticleName        string `json:"articleName"`
	ArticleReportName  string `json:"articleReportName"`
	ArticleUrl         string `json:"articleUrl"`
	ArticleIsTop       string `json:"articleIsTop"`
	ArticleIsImport    string `json:"articleIsImport"`
	ArticleFiledName   string `json:"articleFiledName"`
	ArticleProductName string `json:"articleProductName"`
	ArticleIsLock      string `json:"articleIsLock"`
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

	// mutexAdd()

	//1.创建路由

	r := gin.Default()

	var data1 DataItem
	data1.ArticleId = "1003"
	data1.ArticleName = "以中国新发展为世界提供新机遇"
	data1.ArticleAbstract = "以中国新发展为世界提供新机遇"
	data1.ArticleReportName = "一点资讯"
	data1.ArticleUrl = "www.yidianzixun.com/article/0WBvNwoq"
	data1.ArticleIsTop = "1"
	data1.ArticleIsImport = "1"
	data1.ArticleFiledName = ""
	data1.ArticleProductName = "文字"
	data1.ArticleIsLock = "0"

	data1_str, _ := json.Marshal(data1)

	var data2 DataItem
	data2.ArticleId = "1004"
	data2.ArticleName = "第五届进博会布展完成"
	data2.ArticleAbstract = "第五届进博会布展完成"
	data2.ArticleReportName = "一点资讯"
	data2.ArticleUrl = "www.yidianzixun.com/article/0WBvNwoq2"
	data2.ArticleIsTop = "1"
	data2.ArticleIsImport = "1"
	data2.ArticleFiledName = ""
	data2.ArticleProductName = "文字"
	data2.ArticleIsLock = "0"
	data2_str, _ := json.Marshal(data2)
	fmt.Println(data1_str, data2_str)
	datas := [...]DataItem{
		data1,
		data2,
	}

	//2.绑定路由规则，执行的函数
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status":  "1000",
			"message": "ok",
			"data":    datas,
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

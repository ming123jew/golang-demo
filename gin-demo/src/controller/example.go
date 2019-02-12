package controller

import (
	log "github.com/Sirupsen/logrus"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"rpc-comment/service"

	"rpc-comment/entity"
)

type Example struct{}

func (e *Example) HandleMysql(c *gin.Context) {

	var suid string
	var uid int64
	var err error

	suid = c.Query("uid")
	uid,_ = strconv.ParseInt(suid, 10 ,64)

	userService := new(service.UserService)

	var user entity.Users

	user,err = userService.GetUser(uid)
	if err != nil {
		log.Fatal("error write pid to config file")
	}

	c.JSON(http.StatusOK, gin.H{"status": "ok","suid":suid,"uid":uid,"data":user})
}

func (e *Example) HandleRedis(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "ok","msg":"redis"})
}

func (e *Example) HandleMongo(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "ok","msg":"mongo"})
}

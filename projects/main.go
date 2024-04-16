package main

import (
	"my-project/Controllers"

	"io"
	"os"
	"path"
	"time"

	"github.com/gin-gonic/gin"

	log "github.com/sirupsen/logrus"

	_ "github.com/go-sql-driver/mysql"
)

func init() {

	now := time.Now() //or time.Now().UTC()
	logFileName := now.Format("2006-01-02") + ".log"
	file, err := os.OpenFile(path.Join("./storage/logs", logFileName), os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0644)
	if err != nil {
		log.Error(err)
	}
	log.SetOutput(file)
	log.SetFormatter(&log.JSONFormatter{
		DisableHTMLEscape: true,
		PrettyPrint:       true,
		TimestampFormat:   "2007-02-03 13:05:05",
	})
	gin.DefaultWriter = io.MultiWriter(file)
	log.SetReportCaller(true)
	log.SetLevel(log.DebugLevel)
	log.Debug("logFileName: ", logFileName)
}
func main() {

	//  router settup
	router := gin.Default()
	router.POST("/Req/RespDevice", Controllers.ReqCreateDeviceApi)
	router.POST("/Req/RespVerifyToken", Controllers.RespVerifyTokenApi)
	router.POST("/ListKeys", Controllers.ListKeysApi)
	router.POST("/Req/RespListKeys", Controllers.RespListKeysTknApi)
	router.POST("/Req/RespListAccPvd", Controllers.RespListAccPvdApi)
	router.POST("/Req/RespListAccount", Controllers.RespListAccFetchApi)
	router.POST("/Req/RespListAccount/create", Controllers.ReqListAccountCreateApi)
	router.POST("/Req/RespListKeys-GetWallet", Controllers.RespListKeysGetWalletAPI)
	router.POST("/ReqRegMob", Controllers.ReqRegMobAPI)
	router.POST("/ReqSetCre", Controllers.ReqSetCreAPI)
	router.POST("/ReqUserRegAPI", Controllers.ReqUserRegAPI)
	router.Run(":8080")

}

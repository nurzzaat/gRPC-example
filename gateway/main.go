package main

import (
	"fmt"
	"net/http"
	"os"

	// "net"

	// logrustash "github.com/bshuster-repo/logrus-logstash-hook"
	log "github.com/sirupsen/logrus"
	"go.elastic.co/ecslogrus"

	"github.com/gin-gonic/gin"
)

const (
	appPort = "1232"
)

//	@title			Swagger Example API
//	@version		1.0
//	@description	This is a sample server celler server.

//	@host		localhost:1232
//	@BasePath	/

// @securityDefinitions.apikey	ApiKeyAuth
// @in							header
// @name						Authorization
func main() {
	log.SetFormatter(&ecslogrus.Formatter{})

	// conn, err := net.Dial("tcp", "localhost:5044")
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// hook := logrustash.New(conn, logrustash.DefaultFormatter(log.Fields{
	// 	"type": "ladyTaxi",
	// }))
	// log.AddHook(hook)

	path := "./logs/out.log"
	file, err := os.OpenFile(path, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal(err)
	}
	log.SetOutput(file)

	log.Info("application is running")

	ginRouter := gin.Default()
	ginRouter.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, PATCH, DELETE, OPTIONS ,HEAD")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Access-Control-*, Cross-Origin-Resource-Policy , Origin, X-Requested-With, Content-Type, Accept")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "*")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusOK)
			return
		}
		c.Next()
	})
	Setup(ginRouter)

	ginRouter.Run(fmt.Sprintf(":%s", appPort))
}

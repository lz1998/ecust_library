package main

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/lz1998/ecust_library/config"
	"github.com/lz1998/ecust_library/handler"
	log "github.com/sirupsen/logrus"
)

func init() {
	customFormatter := new(log.TextFormatter)
	customFormatter.TimestampFormat = "2006-01-02 15:04:05"
	customFormatter.FullTimestamp = true
	log.SetFormatter(customFormatter)
}

func main() {
	jwtSecRet := os.Getenv("JWT_SECRET")
	if jwtSecRet != "" {
		config.JwtSecret = []byte(jwtSecRet)
	}

	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()
	{
		group := router.Group("/ecust")

		group.POST("/admin/login", handler.Login)

		group.POST("/admin/list", handler.ListAdmin)
		group.POST("/admin/create", handler.CreateAdmin)
		group.POST("/admin/update", handler.UpdateAdmin)

		group.POST("/book/list", handler.ListBook)
		group.POST("/book/create", handler.CreateBook)
		group.POST("/book/update", handler.UpdateBook)
	}

	log.Infof("success")

	if err := router.Run(":28000"); err != nil {
		panic(err)
	}
}

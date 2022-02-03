package main

import (
	"github.com/Ensena/core/external"
	"github.com/Ensena/core/router"
	"github.com/gin-gonic/gin"
	"go.elastic.co/apm/module/apmgin"
)

func main() {

	ws := gin.Default()
	ws.Use(apmgin.Middleware(ws))
	ws.GET("/me", router.GetInfo)
	ws.GET("/news", router.GetNews)

	/*External APPS  */
	ws.GET("/external", external.ExternalApp)

	ws.Run(":8000")
}

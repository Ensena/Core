package router

import (
	"github.com/Ensena/core/news"
	"github.com/gin-gonic/gin"
)

func GetNews(ctx *gin.Context) {
	username := ctx.Request.URL.Query().Get("username")
	tweets := news.Search(username)
	ctx.JSON(200, tweets)

}

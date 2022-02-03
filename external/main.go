package external

import (
	"net/url"

	"github.com/Ensena/core/env-global"
	"github.com/Ensena/oauth2-server"
	"github.com/elmalba/oauth2-server/jwt"
	"github.com/gin-gonic/gin"
)

var key string

func init() {
	key = env.Check("secretKey", "Missing Params secretKey")

}

type URL struct {
	Url string
}

func ExternalApp(ctx *gin.Context) {
	user := ctx.Request.Header.Get("UserID")
	email := ctx.Request.Header.Get("Email")

	client, err := oauth2.GetApp(ctx, "5sigFSyaY2")
	if !err {
		ctx.AbortWithStatus(401)
		return
	}

	token := jwt.GetToken(user, email, key)

	params, _ := url.ParseQuery("")
	params.Set("code", token)
	uri := client.CallBackURL + `?` + params.Encode()

	ctx.JSON(200, URL{uri})

}

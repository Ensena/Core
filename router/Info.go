package router

import (
	"fmt"
	"strconv"

	"github.com/Ensena/core"
	"github.com/gin-gonic/gin"
	"go.elastic.co/apm"
	"go.elastic.co/apm/module/apmhttp"
)

const TraceID = "traceid"

func GetInfo(ctx *gin.Context) {

	user := ctx.Request.Header.Get("UserID")
	id, err := strconv.Atoi(user)

	var traceID apm.TraceID
	if values := ctx.Request.Header[apmhttp.W3CTraceparentHeader]; len(values) == 1 && values[0] != "" {
		if c, err := apmhttp.ParseTraceparentHeader(values[0]); err == nil {
			traceID = c.Trace
		}
	}

	fmt.Println("trace", traceID.String())
	ctx.Writer.Header().Add(TraceID, traceID.String())
	if err != nil {

		ctx.AbortWithStatus(400)
		return
	}
	ctx.Writer.Write(core.GetInfo(ctx, id))
}

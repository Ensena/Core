package core

import (
	"github.com/gin-gonic/gin"
	"go.elastic.co/apm"
)

func GetInfo(ctx *gin.Context, userID int) []byte {

	span, _ := apm.StartSpan(ctx.Request.Context(), "Get Data From Ensena", "GetData")
	response, user := GetEnsenaData(ctx, userID)
	span.End()

	if user.MoodleUDP {
		span, _ := apm.StartSpan(ctx.Request.Context(), "Get Data From Moodle", "GetData")
		moodleData := GetMoodle(ctx, user.Email)
		span.End()

		response = Mixer(&response, &moodleData, "moodle")
	} else {
		empty := ([]byte("[]"))
		response = Mixer(&response, &empty, "moodle")
	}
	return response
}

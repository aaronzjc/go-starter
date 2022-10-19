package web

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Resp(ctx *gin.Context, errno int, errmsg string, data interface{}) {
	if data == nil {
		data = make(map[string]struct{})
	}
	ctx.JSON(http.StatusOK, map[string]interface{}{
		"errno":  errno,
		"errmsg": errmsg,
		"data":   data,
	})
}

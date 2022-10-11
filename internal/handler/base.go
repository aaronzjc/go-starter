package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func resp(ctx *gin.Context, errno int, errorMsg string, data interface{}) {
	ctx.JSON(http.StatusOK, map[string]interface{}{
		"errno":  errno,
		"errmsg": errorMsg,
		"data":   data,
	})
}

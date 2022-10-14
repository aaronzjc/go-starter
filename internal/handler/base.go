package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type RespData struct {
	Errno  int         `json:"errno"`
	ErrMsg string      `json:"errmsg"`
	Data   interface{} `json:"data"`
}

func resp(ctx *gin.Context, data *RespData) {
	ctx.JSON(http.StatusOK, map[string]interface{}{
		"errno":  data.Errno,
		"errmsg": data.ErrMsg,
		"data":   data.Data,
	})
}

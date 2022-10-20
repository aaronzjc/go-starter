package web

import (
	"go-starter/internal/handler/web/res"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Resp(ctx *gin.Context, errno int, errmsg string, data interface{}) {
	if data == nil {
		data = make(map[string]struct{})
	}
	ctx.JSON(http.StatusOK, &res.RespSt{
		Errno:  errno,
		ErrMsg: errmsg,
		Data:   data,
	})
}

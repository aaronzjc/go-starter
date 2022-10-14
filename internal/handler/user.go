package handler

import (
	"go-starter/internal/constant"
	"go-starter/internal/service"

	"github.com/gin-gonic/gin"
)

type User struct{}

func (h *User) List(ctx *gin.Context) {
	userList, err := service.GetUserList()

	if err != nil {
		resp(ctx, &RespData{Errno: constant.ERR_FAILED, ErrMsg: constant.ERR_MSG_USERLIST})
		return
	}

	resp(ctx, &RespData{
		Errno: constant.ERR_OK,
		Data:  map[string]interface{}{"list": userList},
	})
}

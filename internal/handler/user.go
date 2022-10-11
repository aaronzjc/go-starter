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
		resp(ctx, constant.ERR_FAILED, constant.ERR_MSG_USERLIST, nil)
		return
	}

	resp(ctx, constant.ERR_OK, "", map[string]interface{}{
		"list": userList,
	})
}

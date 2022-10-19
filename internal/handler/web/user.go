package web

import (
	"go-starter/internal/constant"
	"go-starter/internal/handler/web/res"
	"go-starter/internal/service/logic"
	"go-starter/internal/service/store"

	"github.com/gin-gonic/gin"
)

type User struct{}

func (h *User) List(ctx *gin.Context) {
	userRepo, err := store.NewUserRepoImpl()
	if err != nil {
		Resp(ctx, constant.ERR_FAILED, constant.ERR_MSG_USERLIST, nil)
		return
	}
	userLogic := logic.NewUserLogic(userRepo)
	userList, err := userLogic.GetUserList(ctx)
	if err != nil {
		Resp(ctx, constant.ERR_FAILED, constant.ERR_MSG_USERLIST, nil)
		return
	}

	Resp(ctx, constant.ERR_OK, "", res.UserList{List: userList})
}

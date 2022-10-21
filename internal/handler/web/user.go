package web

import (
	"go-starter/internal/application/service"
	"go-starter/internal/application/store"
	"go-starter/internal/constant"
	"go-starter/internal/handler/web/res"

	"github.com/gin-gonic/gin"
)

type User struct {
	svc service.UserService
}

func NewUser() *User {
	repo, _ := store.NewUserRepoImpl()
	svc := service.NewUserService(repo)
	return &User{
		svc: svc,
	}
}

func (h *User) List(ctx *gin.Context) {
	userList, err := h.svc.GetUserList(ctx)
	if err != nil {
		Resp(ctx, constant.ERR_FAILED, constant.ERR_MSG_USERLIST, nil)
		return
	}

	Resp(ctx, constant.ERR_OK, "", res.UserList{List: userList})
}

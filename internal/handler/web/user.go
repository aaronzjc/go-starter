package web

import (
	"go-starter/internal/constant"
	"go-starter/internal/handler/web/res"
	"go-starter/internal/service/logic"
	"go-starter/internal/service/store"

	"github.com/gin-gonic/gin"
)

type User struct {
	l logic.UserLogic
}

func NewUser() *User {
	repo, _ := store.NewUserRepoImpl()
	l := logic.NewUserLogic(repo)
	return &User{
		l: l,
	}
}

func (h *User) List(ctx *gin.Context) {
	userList, err := h.l.GetUserList(ctx)
	if err != nil {
		Resp(ctx, constant.ERR_FAILED, constant.ERR_MSG_USERLIST, nil)
		return
	}

	Resp(ctx, constant.ERR_OK, "", res.UserList{List: userList})
}

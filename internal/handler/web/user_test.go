package web

import (
	"go-starter/internal/constant"
	"go-starter/internal/mocks"
	"go-starter/internal/service/dto"
	"go-starter/test"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func mockUser(t *testing.T) *User {
	l := mocks.NewUserLogic(t)
	l.EXPECT().GetUserList(mock.Anything).Return([]dto.User{{ID: 1, Username: "aaron"}}, nil)
	return &User{l: l}
}

func TestGetUserList(t *testing.T) {
	assert := assert.New(t)
	user := mockUser(t)

	resp := test.NewRequest(t).Handler(user.List).Get("/user/list").Exec()
	assert.Equal(200, resp.Code())
	errno, _, _, err := resp.TryDecode()
	assert.Equal(errno, constant.ERR_OK)
	assert.Nil(err)
}

package web

import (
	"go-starter/internal/application/dto"
	"go-starter/internal/constant"
	"go-starter/internal/mocks"
	"go-starter/test"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func mockUser(t *testing.T) *User {
	svc := mocks.NewUserService(t)
	svc.EXPECT().GetUserList(mock.Anything).Return([]dto.User{{ID: 1, Username: "aaron"}}, nil)
	return &User{svc: svc}
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

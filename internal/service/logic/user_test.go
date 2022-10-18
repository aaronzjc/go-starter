package logic

import (
	"go-starter/internal/domain/model"
	"go-starter/internal/mock"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUserGetAll(t *testing.T) {
	assert := assert.New(t)

	userRepo := mock.NewUserRepo(t)
	userRepo.EXPECT().GetAll().Return([]model.User{{BaseModel: model.BaseModel{ID: 1}, Username: "aaron"}}, nil)

	users, _ := GetUserList(userRepo)
	assert.NotEmpty(users)
	userRepo.AssertExpectations(t)
}

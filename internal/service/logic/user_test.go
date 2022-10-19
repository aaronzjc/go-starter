package logic

import (
	"context"
	"go-starter/internal/domain/model"
	"go-starter/internal/mocks"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestUserGetAll(t *testing.T) {
	assert := assert.New(t)

	userRepo := mocks.NewUserRepo(t)
	userRepo.EXPECT().GetAll(mock.Anything).Return([]model.User{{BaseModel: model.BaseModel{ID: 1}, Username: "aaron"}}, nil)

	userLogic := NewUserLogic(userRepo)
	users, _ := userLogic.GetUserList(context.Background())
	assert.NotEmpty(users)
	userRepo.AssertExpectations(t)
}

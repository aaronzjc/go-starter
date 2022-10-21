package service

import (
	"context"
	"go-starter/internal/application/dto"
	"go-starter/internal/domain/repo"
	"go-starter/pkg/helper"
)

type UserService interface {
	GetUserList(context.Context) ([]dto.User, error)
}

type UserServiceImpl struct {
	repo repo.UserRepo
}

func (l *UserServiceImpl) GetUserList(ctx context.Context) ([]dto.User, error) {
	var users []dto.User
	userModels, err := l.repo.GetAll(ctx)
	if err != nil {
		return nil, err
	}
	for _, v := range userModels {
		users = append(users, dto.User{
			ID:       v.ID,
			Username: v.Username,
			CreateAt: helper.TimeToLocalStr(v.CreatedAt),
		})
	}
	return users, nil
}

func NewUserService(repo repo.UserRepo) *UserServiceImpl {
	return &UserServiceImpl{repo: repo}
}

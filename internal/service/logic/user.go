package logic

import (
	"context"
	"go-starter/internal/domain/repo"
	"go-starter/internal/service/dto"
	"go-starter/pkg/helper"
)

type UserLogic struct {
	repo repo.UserRepo
}

func (l *UserLogic) GetUserList(ctx context.Context) ([]dto.User, error) {
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

func NewUserLogic(repo repo.UserRepo) *UserLogic {
	return &UserLogic{repo: repo}
}

package store

import (
	"context"
	"go-starter/internal/domain/model"
	"go-starter/internal/domain/repo"
)

type UserRepoImpl struct {
}

var _ repo.UserRepo = &UserRepoImpl{}

func (r *UserRepoImpl) GetAll(ctx context.Context) ([]model.User, error) {
	return []model.User{
		{
			Username: "aaronzjc",
		},
	}, nil
}

func NewUserRepoImpl() (*UserRepoImpl, error) {
	return &UserRepoImpl{}, nil
}

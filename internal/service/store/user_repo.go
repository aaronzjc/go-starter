package store

import (
	"context"
	"errors"
	"go-starter/internal/db"
	"go-starter/internal/domain/model"
	"go-starter/internal/domain/repo"

	"gorm.io/gorm"
)

type UserRepoImpl struct {
	db *gorm.DB
}

var _ repo.UserRepo = &UserRepoImpl{}

func (r *UserRepoImpl) GetAll(ctx context.Context) ([]model.User, error) {
	users := []model.User{}
	st := r.db.Find(&users)
	if st.Error != nil {
		return nil, errors.New("get users err")
	}
	return users, nil
}

func NewUserRepoImpl() (*UserRepoImpl, error) {
	demo, ok := db.Get(model.DB_DEMO)
	if !ok {
		return nil, errors.New("db not connected")
	}
	return &UserRepoImpl{
		db: demo,
	}, nil
}

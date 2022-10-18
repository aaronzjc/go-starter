package repo

import (
	"errors"
	"go-starter/internal/db"
	"go-starter/internal/domain/model"
	iRepo "go-starter/internal/domain/repo"

	"gorm.io/gorm"
)

type UserRepoImpl struct {
	db *gorm.DB
}

var _ iRepo.UserRepo = &UserRepoImpl{}

func (r *UserRepoImpl) GetAll() ([]model.User, error) {
	users := []model.User{}
	st := r.db.Find(&users)
	if st.Error != nil {
		return nil, errors.New("get users err")
	}
	return users, nil
}

func NewUserRepoImpl() (*UserRepoImpl, error) {
	demo, err := db.Get(model.DB_DEMO)
	if err != nil {
		return nil, err
	}
	return &UserRepoImpl{
		db: demo,
	}, nil
}

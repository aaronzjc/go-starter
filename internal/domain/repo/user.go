package repo

import "go-starter/internal/domain/model"

// User 用户相关行为
type UserRepo interface {
	GetAll() ([]model.User, error)
}

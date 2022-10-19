package repo

import (
	"context"
	"go-starter/internal/domain/model"
)

// User 用户相关行为
type UserRepo interface {
	GetAll(context.Context) ([]model.User, error)
}

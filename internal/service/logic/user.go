package logic

import (
	"go-starter/internal/domain/repo"
	"go-starter/internal/service/dto"
	"go-starter/pkg/helper"
)

func GetUserList(userRepo repo.UserRepo) ([]dto.User, error) {
	var users []dto.User
	userModels, err := userRepo.GetAll()
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

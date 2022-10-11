package service

import (
	"errors"
	"go-starter/internal/model"
	"go-starter/pkg/helper"
)

type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	RegistAt string `json:"regist_at"`
}

func GetUserList() ([]User, error) {
	userModel := model.User{}
	rows, err := userModel.FetchAll()
	if err != nil {
		return nil, errors.New("get userlist error")
	}
	var users []User
	for _, v := range rows {
		users = append(users, User{
			ID:       v.ID,
			Username: v.Username,
			RegistAt: helper.TimeToLocalStr(v.CreatedAt),
		})
	}
	return users, nil
}

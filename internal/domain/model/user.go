package model

import (
	"go-starter/internal/db"
)

type User struct {
	Username string `gorm:"column:username" json:"username"`

	BaseModel
}

func (User) TableName() string {
	return "user"
}

func (User) FetchAll() ([]User, error) {
	demo, err := db.Get(DB_DEMO)
	if err != nil {
		return nil, err
	}
	var users []User
	demo.Find(&users)
	return users, nil
}

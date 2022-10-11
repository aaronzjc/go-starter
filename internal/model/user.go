package model

import (
	"go-starter/pkg/db"
	"time"
)

type User struct {
	ID        int       `gorm:"column:id" json:"id"`
	Username  string    `gorm:"column:username" json:"username"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
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

package model

type User struct {
	Username string `gorm:"column:username" json:"username"`

	BaseModel
}

func (User) TableName() string {
	return "user"
}

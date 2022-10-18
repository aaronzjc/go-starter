package model

import "time"

const DB_DEMO = "demo"

type BaseModel struct {
	ID        uint      `gorm:"column:id" json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

package model

import "time"

type User struct {
	Id        uint   `gorm:"primaryKey"`
	Email     string `gorm:"unique"`
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
}

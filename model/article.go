package model

import (
	"time"
)

type Article struct {
	Id        int `gorm:"primaryKey"`
	Title     string
	Content   string
	IsAtHome  bool
	CreatedAt time.Time
	UpdatedAt time.Time
}

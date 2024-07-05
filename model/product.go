package model

import (
	"time"
)

type Product struct {
	Id        int `gorm:"primaryKey"`
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
	ProductID int
}

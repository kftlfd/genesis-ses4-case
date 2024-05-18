package models

import (
	"time"
)

type Subscriber struct {
	ID        uint      `gorm:"primaryKey;autoIncrement"`
	Email     string    `gorm:"not null;unique;index"`
	CreatedAt time.Time `gorm:"not null;autoCreateTime"`
}

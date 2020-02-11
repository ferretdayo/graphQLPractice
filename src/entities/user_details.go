package entities

import "time"

type UserDetail struct {
	UserID    uint      `gorm:"primary_key"`
	Nickname  string    `gorm:"not null"`
	Birthday  time.Time `gorm:"not null"`
	HobbyID   uint      `gorm:"not null"`
	Hobby     Hobby
	CreatedAt time.Time
	UpdatedAt time.Time
}

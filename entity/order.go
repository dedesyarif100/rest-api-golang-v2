package entity

import "time"

type Order struct {
	// gorm.Model
	ID			int
	UserId		int
	BookId		int
	Amount		int
	State		string
	CreatedAt	time.Time
	UpdatedAt	time.Time
	DeletedAt   time.Time
}
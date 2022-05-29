package book

import (
	"time"
	// "github.com/jinzhu/gorm"
)

type User struct {
	ID 			int
	Name		string
	Age			int
	CreatedAt	time.Time
	UpdatedAt	time.Time
}

type Book struct {
	ID			int
	Title		string
	Age			int
	Description string
	Price		int
	Rating		int
	Discount	int
	CreatedAt	time.Time
	UpdatedAt	time.Time
}

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

type CreditCard struct {
	ID			int
	BookId		int
	Name		string
	Number		int
	CreatedAt	time.Time
	UpdatedAt	time.Time
}

type Result struct {
	Date  time.Time
	Total int64
}
package book

import "time"

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
	ID			int
	Amount		int
	State		string
	CreatedAt	time.Time
	UpdatedAt	time.Time
	DeletedAt   time.Time
}
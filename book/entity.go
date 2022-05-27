package book

import "time"

type Book struct {
	ID			int
	Title		string
	Description string
	Price		int
	Rating		int
	Discount	int
	CreatedAt	time.Time
	UpdatedAt	time.Time
}

type Assets struct {
	ID			int
	Name		string
	Code		string
	CreatedAt	time.Time
	UpdatedAt	time.Time
}
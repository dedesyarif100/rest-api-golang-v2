package entity

import "time"

type CreditCard struct {
	ID			int
	BookId		int
	Name		string
	Number		int
	CreatedAt	time.Time
	UpdatedAt	time.Time
}
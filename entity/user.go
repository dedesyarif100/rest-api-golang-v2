package entity

import "time"

type User struct {
	ID 			int
	CompanyId	int
	Name		string
	Age			int
	Gender 		string
	Role 		string
	CreatedAt	time.Time
	UpdatedAt	time.Time
}
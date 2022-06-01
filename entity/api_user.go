package entity

import "time"

type ApiUser struct {
	ID  		int
	Name 		string
	CreatedAt	time.Time
	UpdatedAt	time.Time
}
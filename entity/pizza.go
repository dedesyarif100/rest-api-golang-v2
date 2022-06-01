package entity

import "time"

type Pizza struct {
	ID 			int
	Pizza		string
	Size		string
	CreatedAt	time.Time
	UpdatedAt	time.Time
}
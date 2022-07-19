package Contents

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID 			int
	CompanyID 	int
	Name 		string
	Age 		int
	Gender 		string
	Role 		string
	Company		[]Company
}

type Company struct {
	gorm.Model
	ID 			int
	Name 		string
}

func Preload(db *gorm.DB, err any) {
	db.Debug().Preload()
}

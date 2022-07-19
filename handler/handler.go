package handler

import (
	// "rest-api-golang-v2/handler/AllHandler/CrudInterface"
	"rest-api-golang-v2/handler/AllHandler/Associations"
	"gorm.io/gorm"
)

func Handler(db *gorm.DB, err any) {
	// CrudInterface.CrudInterface(db, err)

	Associations.Associations(db, err)
}

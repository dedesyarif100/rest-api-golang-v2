package Constraints

import (
	"gorm.io/gorm"
	"rest-api-golang-v2/handler/AllHandler/AdvancedTopics/Constraints/Contents"
)

func Constraints(db *gorm.DB, err any) {
	Contents.CheckConstraint(db, err)
	Contents.ForeignKeyConstraint(db, err)
	Contents.IndexConstraint(db, err)
}
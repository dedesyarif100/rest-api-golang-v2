package Indexes

import (
	"gorm.io/gorm"
	"rest-api-golang-v2/handler/AllHandler/AdvancedTopics/Indexes/Contents"
)

func Indexes(db *gorm.DB, err any) {
	Contents.CompositeIndexes(db, err)
	Contents.IndexTag(db, err)
	Contents.MultipleIndexes(db, err)
}
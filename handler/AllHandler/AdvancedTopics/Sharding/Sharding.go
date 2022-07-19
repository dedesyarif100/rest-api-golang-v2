package Sharding

import (
	"gorm.io/gorm"
	"rest-api-golang-v2/handler/AllHandler/AdvancedTopics/Sharding/Contents"
)

func Sharding(db *gorm.DB, err any) {
	Contents.Usage(db, err)
}
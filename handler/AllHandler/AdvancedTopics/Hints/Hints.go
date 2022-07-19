package Hints

import (
	"gorm.io/gorm"
	"rest-api-golang-v2/handler/AllHandler/AdvancedTopics/Hints/Contents"
)

func Hints(db *gorm.DB, err any) {
	Contents.CommentHints(db, err)
	Contents.IndexHints(db, err)
	Contents.OptimizerHints(db, err)
}
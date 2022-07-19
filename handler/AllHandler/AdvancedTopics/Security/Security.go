package Security

import (
	"gorm.io/gorm"
	"rest-api-golang-v2/handler/AllHandler/AdvancedTopics/Security/Contents"
)

func Security(db *gorm.DB, err any) {
	Contents.InlineCondition(db, err)
	Contents.QueryCondition(db, err)
	Contents.SQLinjectionMethods(db, err)
}
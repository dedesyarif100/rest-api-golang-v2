package WritePlugins

import (
	"gorm.io/gorm"
	"rest-api-golang-v2/handler/AllHandler/AdvancedTopics/WritePlugins/Contents"
)

func WritePlugins(db *gorm.DB, err any) {
	Contents.Callbacks(db, err)
	Contents.Plugin(db, err)
}
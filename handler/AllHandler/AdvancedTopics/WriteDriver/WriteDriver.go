package WriteDriver

import (
	"gorm.io/gorm"
	"rest-api-golang-v2/handler/AllHandler/AdvancedTopics/WriteDriver/Contents"
)

func WriteDriver(db *gorm.DB, err any) {
	Contents.WriteNewDriver(db, err)
}
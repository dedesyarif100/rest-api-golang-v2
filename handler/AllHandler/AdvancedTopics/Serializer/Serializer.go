package Serializer

import (
	"gorm.io/gorm"
	"rest-api-golang-v2/handler/AllHandler/AdvancedTopics/Serializer/Contents"
)

func Serializer(db *gorm.DB, err any) {
	Contents.CustomizedSerializerType(db, err)
	Contents.RegisterSerializer(db, err)
}
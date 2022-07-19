package CustomizeDataTypes

import (
	"gorm.io/gorm"
	"rest-api-golang-v2/handler/AllHandler/Tutorials/CustomizeDataTypes/Contents"
)

func Conventions(db *gorm.DB, err any) {
	Contents.CustomizedDataTypesCollections(db, err)
	Contents.ImplementsCustomizedDataType(db, err)
}
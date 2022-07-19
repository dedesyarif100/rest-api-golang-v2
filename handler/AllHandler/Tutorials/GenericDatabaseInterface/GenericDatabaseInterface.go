package GenericDatabaseInterface

import (
	"gorm.io/gorm"
	"rest-api-golang-v2/handler/AllHandler/Tutorials/GenericDatabaseInterface/Contents"
)

func GenericDatabaseInterface(db *gorm.DB, err any) {
	Contents.ConnectionPool(db, err)
}
package Logger

import (
	"gorm.io/gorm"
	"rest-api-golang-v2/handler/AllHandler/Tutorials/Logger/Contents"
)

func Logger(db *gorm.DB, err any) {
	Contents.CustomizeLogger(db, err)
	Contents.Logger(db, err)
}
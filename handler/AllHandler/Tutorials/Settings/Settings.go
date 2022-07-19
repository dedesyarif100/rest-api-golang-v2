package Settings

import (
	"gorm.io/gorm"
	"rest-api-golang-v2/handler/AllHandler/Tutorials/Settings/Contents"
)

func Performance(db *gorm.DB, err any) {
	Contents.InstanceSetInstanceGet(db, err)
	Contents.SetGet(db, err)
}
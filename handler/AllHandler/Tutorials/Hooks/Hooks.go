package Hooks

import (
	"gorm.io/gorm"
	"rest-api-golang-v2/handler/AllHandler/Tutorials/Hooks/Contents"
)

func Hooks(db *gorm.DB, err any) {
	Contents.Hooks(db, err)
	Contents.ModifyCurrentOperation(db, err)
	Contents.ObjectLifeCycle(db, err)
}
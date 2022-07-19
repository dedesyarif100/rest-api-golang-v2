package MethodChaining

import (
	"gorm.io/gorm"
	"rest-api-golang-v2/handler/AllHandler/Tutorials/MethodChaining/Contents"
)

func MethodChaining(db *gorm.DB, err any) {
	Contents.ChainMethod(db, err)
	Contents.FinisherMethod(db, err)
	Contents.NewSessionMethod(db, err)
}
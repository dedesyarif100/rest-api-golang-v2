package ErrorHandling

import (
	"gorm.io/gorm"
	"rest-api-golang-v2/handler/AllHandler/Tutorials/ErrorHandling/Contents"
)

func ErrorHandling(db *gorm.DB, err any) {
	Contents.ErrorHandling(db, err)
	Contents.Errors(db, err)
	Contents.ErrRecordNotFound(db, err)
}
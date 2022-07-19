package Conventions

import (
	"gorm.io/gorm"
	"rest-api-golang-v2/handler/AllHandler/Tutorials/Conventions/Contents"
)

func Conventions(db *gorm.DB, err any) {
	Contents.ColumnName(db, err)
	Contents.IdAsPrimaryKey(db, err)
	Contents.PluralizedTableName(db, err)
	Contents.TimestampTracking(db, err)
}
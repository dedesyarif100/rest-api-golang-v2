package Migration

import (
	"gorm.io/gorm"
	"rest-api-golang-v2/handler/AllHandler/Tutorials/Migration/Contents"
)

func MethodChaining(db *gorm.DB, err any) {
	Contents.AutoMigration(db, err)
	Contents.Constraints(db, err)
	Contents.MigratorInterface(db, err)
	Contents.OtherMigrationTools(db, err)
}
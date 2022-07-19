package GORMConfig

import (
	"gorm.io/gorm"
	"rest-api-golang-v2/handler/AllHandler/AdvancedTopics/GORMConfig/Contents"
)

func GORMConfig(db *gorm.DB, err any) {
	Contents.AllowGlobalUpdate(db, err)
	Contents.DisableAutomaticPing(db, err)
	Contents.DisableForeignKeyConstraintWhenMigrating(db, err)
	Contents.DisableNestedTransaction(db, err)
	Contents.DryRun(db, err)
	Contents.Logger(db, err)
	Contents.NamingStrategy(db, err)
	Contents.NowFunc(db, err)
	Contents.PrepareStmt(db, err)
	Contents.SkipDefaultTransaction(db, err)
}
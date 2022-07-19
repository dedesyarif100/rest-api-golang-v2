package Session

import (
	"gorm.io/gorm"
	"rest-api-golang-v2/handler/AllHandler/Tutorials/Session/Contents"
)

func Performance(db *gorm.DB, err any) {
	Contents.AllowGlobalUpdate(db, err)
	Contents.Context(db, err)
	Contents.CreateBatchSize(db, err)
	Contents.Debug(db, err)
	Contents.DisableNestedTransaction(db, err)
	Contents.DryRun(db, err)
	Contents.FullSaveAssociations(db, err)
	Contents.Initialized(db, err)
	Contents.Logger(db, err)
	Contents.NewDB(db, err)
	Contents.NowFunc(db, err)
	Contents.PrepareStmt(db, err)
	Contents.QueryFields(db, err)
	Contents.SkipHooks(db, err)
}
package Performance

import (
	"gorm.io/gorm"
	"rest-api-golang-v2/handler/AllHandler/Tutorials/Performance/Contents"
)

func Performance(db *gorm.DB, err any) {
	Contents.CachesPreparedStatement(db, err)
	Contents.DisableDefaultTransaction(db, err)
	Contents.IndexHints(db, err)
	Contents.IterationFindInBatches(db, err)
	Contents.SelectFields(db, err)
}
package ErrorHandling

import (
	"gorm.io/gorm"
	"rest-api-golang-v2/handler/AllHandler/Tutorials/Context/Contents"
)

func Context(db *gorm.DB, err any) {
	Contents.ChiMiddlewareExample(db, err)
	Contents.ContextInHooksCallbacks(db, err)
	Contents.ContextTimeout(db, err)
	Contents.ContinuousSessionMode(db, err)
	Contents.Logger(db, err)
	Contents.SingleSessionMode(db, err)
}
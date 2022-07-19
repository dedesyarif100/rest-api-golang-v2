package DatabaseResolver

import (
	"gorm.io/gorm"
	"rest-api-golang-v2/handler/AllHandler/AdvancedTopics/DatabaseResolver/Contents"
)

func DatabaseResolver(db *gorm.DB, err any) {
	Contents.AutomaticConnectionSwitching(db, err)
	Contents.ConnectionPool(db, err)
	Contents.LoadBalancing(db, err)
	Contents.ManualConnectionSwitching(db, err)
	Contents.ReadWriteSplitting(db, err)
	Contents.Transaction(db, err)
	Contents.Usage(db, err)
}
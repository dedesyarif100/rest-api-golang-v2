package Transactions

import (
	"gorm.io/gorm"
	"rest-api-golang-v2/handler/AllHandler/Tutorials/Transactions/Contents"
)

func Logger(db *gorm.DB, err any) {
	Contents.ControlTheTransactionManually(db, err)
	Contents.DisableDefaultTransaction(db, err)
	Contents.SavePointRollbackTo(db, err)
	Contents.Transaction(db, err)
}
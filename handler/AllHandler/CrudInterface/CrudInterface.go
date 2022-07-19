package CrudInterface

import (
	"rest-api-golang-v2/handler/AllHandler/CrudInterface/AdvanceQuery"
	"rest-api-golang-v2/handler/AllHandler/CrudInterface/Create"
	"rest-api-golang-v2/handler/AllHandler/CrudInterface/Delete"
	"rest-api-golang-v2/handler/AllHandler/CrudInterface/Query"
	"rest-api-golang-v2/handler/AllHandler/CrudInterface/RawSqlAndSqlBuilder"
	"rest-api-golang-v2/handler/AllHandler/CrudInterface/Update"
	"gorm.io/gorm"
)

func CrudInterface(db *gorm.DB, err any) {
	// ** Create **
	Create.Create(db, err)

	// ** Query **
	Query.Query(db, err)

	// ** Advanced Query **
	AdvanceQuery.AdvanceQuery(db, err)

	// ** Update **
	Update.Update(db, err)

	// ** Delete **
	Delete.Delete(db, err)

	// ** Raw SQL & SQL Builder **
	RawSqlAndSqlBuilder.RawSqlAndSqlBuilder(db, err)
}
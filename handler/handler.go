package handler

import (
	"gorm.io/gorm"
	"rest-api-golang/handler/AllHandler/CrudInterface/Query"
	"rest-api-golang/handler/AllHandler/CrudInterface/Create"
	"rest-api-golang/handler/AllHandler/CrudInterface/Delete"
	"rest-api-golang/handler/AllHandler/CrudInterface/Update"
	"rest-api-golang/handler/AllHandler/CrudInterface/AdvanceQuery"
	"rest-api-golang/handler/AllHandler/CrudInterface/RawSqlAndSqlBuilder"
)

func Handler(db *gorm.DB, err any) {
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
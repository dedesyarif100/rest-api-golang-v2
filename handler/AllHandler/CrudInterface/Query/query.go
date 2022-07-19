package Query

import (
	"rest-api-golang-v2/handler/AllHandler/CrudInterface/Query/Contents/Conditions"
	// "rest-api-golang-v2/handler/AllHandler/CrudInterface/Query/Contents"
	"gorm.io/gorm"
	"fmt"
)

func Query(db *gorm.DB, err any) {
	var amount int
	db.Debug().Raw("SELECT SUM(amount) FROM orders WHERE state = ?", "UNKNOW").Scan(&amount)
	fmt.Println(amount)
	// QUERY
	// Retrieving a single object
		// Contents.RetrievingSingleObject(db, err)

	// Retrieving all objects
		// Contents.RetrievingAllObjects(db, err)

	// CONDITIONS
		Conditions.Conditions(db, err)

	// Order
		// Contents.Order(db, err)

	// Limit & Offset
		// Contents.LimitAndOffset(db, err)

	// Group By & Having
		// Contents.GroupByAndHaving(db, err)

	// Joins
		// Contents.Joins(db, err)

	// Scan
		// Contents.Scan(db, err)

	// Select
		// Contents.Select(db, err)
}

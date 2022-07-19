package Contents

import (
	"fmt"
	"log"
	"rest-api-golang-v2/entity"

	"gorm.io/gorm"
)

func GroupByAndHaving(db *gorm.DB, err any) {
	var results []entity.Result
	err = db.Debug().Table("orders").Select("date(created_at) as date, sum(amount) as total").Group("date(created_at)").Having("sum(amount) >= ?", 20000).Scan(&results)
	if err != nil {
		log.Fatal("============= QUERY ERROR =============")
	}
	println()
	fmt.Println("DATA RESULTS :", results)
	println()
}

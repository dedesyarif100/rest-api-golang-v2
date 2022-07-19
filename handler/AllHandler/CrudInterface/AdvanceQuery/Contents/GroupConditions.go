package Contents

import (
	"fmt"
	// "log"
	"rest-api-golang-v2/entity"

	"gorm.io/gorm"
)

func GroupConditions(db *gorm.DB, err any) {
	var pizzas []entity.Pizza
	db.Debug().Where(
		db.Where("pizza = ?", "pepperoni").Where(db.Where("size = ?", "small").Or("size = ?", "medium")),
	).Or(
		db.Where("pizza = ?", "hawaiian").Where("size = ?", "xlarge"),
	).Find(&pizzas)
	println()
	fmt.Println("DATA : >>>>>>>>>>>>>>>>>>>>")
	for _, data := range pizzas {
		fmt.Println("---------------------------")
		fmt.Println("ID 		:", data.ID)
		fmt.Println("PIZZA 		:", data.Pizza)
		fmt.Println("SIZE 		:", data.Size)
		fmt.Println("---------------------------")
	}
	println()
}

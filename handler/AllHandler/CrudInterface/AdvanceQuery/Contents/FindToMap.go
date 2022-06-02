package Contents

import (
	"fmt"
	// "log"
	"rest-api-golang/entity"
	"gorm.io/gorm"
)

func FindToMap(db *gorm.DB, err any) {
	var users []entity.User
	result := map[string]interface{}{}
	db.Debug().Model(&users).First(&result, "id = ?", 1)
	// fmt.Println(result)
	fmt.Println("----------------------------")
	for key, data := range result {
		fmt.Println(key,"  :",data)
	}
	fmt.Println("----------------------------")

	var results []map[string]interface{}
	db.Debug().Table("users").Find(&results)
	println()
	for _, result := range results {
		fmt.Println("-----------------------------------------------------")
		for key, data := range result {
			fmt.Println(key,"  :",data)
		}
		fmt.Println("-----------------------------------------------------")
	}
	println()
}
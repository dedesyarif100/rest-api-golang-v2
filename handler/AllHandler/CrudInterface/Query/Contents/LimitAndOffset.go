package Contents

import (
	"fmt"
	"log"
	"rest-api-golang-v2/entity"

	"gorm.io/gorm"
)

func LimitAndOffset(db *gorm.DB, err any) {
	// var books []entity.Book
	// err = db.Debug().Limit(2).Find(&books).Error
	// if err != nil {
	// 	log.Fatal("============= QUERY ERROR =============")
	// }
	// println()
	// fmt.Println("DATA BOOK :",books)
	// println()

	var books []entity.Book
	err = db.Debug().Limit(2).Offset(3).Find(&books).Error
	if err != nil {
		log.Fatal("============= QUERY ERROR =============")
	}
	println()
	fmt.Println("DATA BOOK :", books)
	println()
}

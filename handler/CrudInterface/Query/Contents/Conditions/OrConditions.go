package Conditions

import (
	"fmt"
	"log"
	"rest-api-golang/entity"

	"gorm.io/gorm"
)

func OrConditions(db *gorm.DB, err any) {
	// var books []entity.Book
	// err = db.Debug().Where("title = ?", "HENDRS").Or("title = ?", "RIAND").Find(&books).Error
	// if err != nil {
	// 	log.Fatal("============= QUERY ERROR =============")
	// }
	// println()
	// fmt.Println("DATA BOOK :",books)
	// println()


	// var books []entity.Book
	// err = db.Debug().Where("description LIKE ?", "%php%").Or("description LIKE ?", "%golang%").Find(&books).Error
	// if err != nil {
	// 	log.Fatal("============= QUERY ERROR =============")
	// }
	// println()
	// fmt.Println("DATA BOOK :",books)
	// println()


	// var books []entity.Book
	// err = db.Debug().Where("title = 'RIAND'").Or(entity.Book{Title: "HENDRS"}).Find(&books).Error // STRUCT
	// if err != nil {
	// 	log.Fatal("============= QUERY ERROR =============")
	// }
	// println()
	// fmt.Println("DATA BOOK :",books)
	// println()


	var books []entity.Book
	err = db.Debug().Where("title = 'RIAND'").Or(map[string]interface{}{"title" : "HENDRS"}).Find(&books).Error // MAP
	if err != nil {
		log.Fatal("============= QUERY ERROR =============")
	}
	println()
	fmt.Println("DATA BOOK :",books)
	println()
}
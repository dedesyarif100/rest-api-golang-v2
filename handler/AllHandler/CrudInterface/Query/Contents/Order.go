package Contents

import (
	"fmt"
	"log"
	"rest-api-golang/entity"
	"gorm.io/gorm"
)

func Order(db *gorm.DB, err any) {
	// var books []entity.Book
	// err = db.Debug().Order("age desc").Find(&books).Error
	// if err != nil {
	// 	log.Fatal("============= QUERY ERROR =============")
	// }
	// println()
	// fmt.Println("DATA BOOK :",books)
	// println()


	// var books []entity.Book
	// err = db.Debug().Order("age desc").Order("title").Find(&books).Error
	// if err != nil {
	// 	log.Fatal("============= QUERY ERROR =============")
	// }
	// println()
	// fmt.Println("DATA BOOK :",books)
	// println()


	var books []entity.Book
	err = db.Debug().Order("age").Find(&books).Error
	if err != nil {
		log.Fatal("============= QUERY ERROR =============")
	}
	println()
	fmt.Println("DATA BOOK :",books)
	println()
}
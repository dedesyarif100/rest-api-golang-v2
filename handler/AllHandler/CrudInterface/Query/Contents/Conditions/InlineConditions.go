package Conditions

import (
	"fmt"
	"log"
	"rest-api-golang-v2/entity"

	"gorm.io/gorm"
)

func InlineConditions(db *gorm.DB, err any) {
	// var books []entity.Book
	// err = db.Debug().First(&books).Error
	// if err != nil {
	// 	log.Fatal("============= QUERY ERROR =============")
	// }
	// println()
	// fmt.Println("DATA BOOK :",books)
	// println()

	// Plain SQL
	// var books []entity.Book
	// err = db.Debug().Find(&books, "title = ?", "TRISULI").Error
	// if err != nil {
	// 	log.Fatal("============= QUERY ERROR =============")
	// }
	// println()
	// fmt.Println("DATA BOOK :",books)
	// println()

	// var books []entity.Book
	// err = db.Debug().Find(&books, "title <> ? AND rating >= ?", "RIAND", 4).Error
	// if err != nil {
	// 	log.Fatal("============= QUERY ERROR =============")
	// }
	// println()
	// fmt.Println("DATA BOOK :",books)
	// println()

	// struct
	// var books []entity.Book
	// err = db.Debug().Find(&books, entity.Book{Rating: 6}).Error
	// if err != nil {
	// 	log.Fatal("============= QUERY ERROR =============")
	// }
	// println()
	// fmt.Println("DATA BOOK :",books)
	// println()

	// Map
	var books []entity.Book
	err = db.Debug().Find(&books, map[string]any{"rating": 6}).Error
	if err != nil {
		log.Fatal("============= QUERY ERROR =============")
	}
	println()
	fmt.Println("DATA BOOK :", books)
	println()
}

package Conditions

import (
	"fmt"
	"log"
	"rest-api-golang/entity"
	"gorm.io/gorm"
)

func StructAndMapConditions(db *gorm.DB, err any) {
	// var books []entity.Book
	// err = db.Debug().Where(&entity.Book{Title: "DEDE", Age: 22}).Find(&books).Error
	// if err != nil {
	// 	log.Fatal("============= QUERY ERROR =============")
	// }
	// println()
	// fmt.Println("DATA BOOK :",books)
	// println()


	// var books []entity.Book
	// err = db.Debug().Where(map[string]interface{}{"title": "TRISULI", "price": "115000"}).Find(&books).Error // Map		
	// if err != nil {
	// 	log.Fatal("============= QUERY ERROR =============")
	// }
	// println()
	// fmt.Println("DATA BOOK :",books)
	// println()


	var books []entity.Book
	err = db.Debug().Where([]int64{2, 4}).Find(&books).Error; // Slice of primary keys		
	if err != nil {
		log.Fatal("============= QUERY ERROR =============")
	}
	println()
	fmt.Println("DATA BOOK :",books)
	println()
}
package Contents

import (
	"fmt"
	"log"
	"rest-api-golang-v2/entity"

	"gorm.io/gorm"
)

func Scan(db *gorm.DB, err any) {
	// var books []entity.Book
	// err = db.Debug().Table("books").Select("title").Where("title = ?", "dede").Scan(&books).Error
	// if err != nil {
	// 	log.Fatal("============= QUERY ERROR =============")
	// }
	// println()
	// fmt.Println("DATA BOOK :", books)
	// println()

	var books []entity.Book
	err = db.Debug().Raw("SELECT title, age FROM books WHERE title = ?", "RIAND").Scan(&books).Error
	if err != nil {
		log.Fatal("============= QUERY ERROR =============")
	}
	println()
	fmt.Println("DATA BOOK :", books)
	println()
}

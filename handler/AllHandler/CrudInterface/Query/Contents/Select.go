package Contents

import (
	"fmt"
	"log"
	"rest-api-golang-v2/entity"
	"gorm.io/gorm"
)

func Select(db *gorm.DB, err any) {
	var books []entity.Book
	// err = db.Debug().Select("id, title").Find(&books).Error
	err = db.Debug().Select([]string{"id", "title", "age", "price"}).Find(&books).Error
	if err != nil {
		log.Fatal("============= QUERY ERROR =============")
	}
	println()
	fmt.Println("DATA BOOK :", books[0])
	println()
}

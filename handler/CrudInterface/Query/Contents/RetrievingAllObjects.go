package Contents

import (
	"fmt"
	"rest-api-golang/entity"
	"gorm.io/gorm"
)

func RetrievingAllObjects(db *gorm.DB, err any) {
	var books []entity.Book
	result := db.Debug().Find(&books)
	println()
	fmt.Println("DATA BOOKS :", books)
	println()
	fmt.Println("RESULT :", result.RowsAffected)
	println()
}
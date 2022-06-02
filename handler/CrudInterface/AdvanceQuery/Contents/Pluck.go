package Contents

import (
	"fmt"
	// "log"
	"rest-api-golang/entity"
	"gorm.io/gorm"
)

func Pluck(db *gorm.DB, err any) {
	// var ages []int
	// db.Debug().Find(&entity.User{}).Pluck("age", &ages)
	// fmt.Println("VALUE AGES :",ages)
	// println()


	// var names []string
	// db.Debug().Model(&entity.User{}).Pluck("name", &names)
	// fmt.Println("\nVALUE NAMES :",names)
	// println()


	// var id []int
	// db.Debug().Model(&entity.User{}).Pluck("id", &id)
	// fmt.Println("\nVALUE ID :",id)
	// println()


	// var names []string
	// db.Debug().Table("users").Pluck("name", &names)
	// fmt.Println("\nVALUE NAMES :",names)
	// println()


	// var names []string
	// db.Debug().Table("books").Pluck("title", &names)
	// fmt.Println("\nVALUE BOOKS :",names)
	// println()


	var books []entity.Book
	db.Debug().Select("title, description").Find(&books)
	fmt.Println("\nVALUE BOOKS :",books[1].Title)
	println()
}
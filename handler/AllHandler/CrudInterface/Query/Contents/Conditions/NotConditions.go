package Conditions

import (
	"fmt"
	"log"
	"rest-api-golang-v2/entity"

	"gorm.io/gorm"
)

func NotConditions(db *gorm.DB, err any) {
	// var books []entity.Book
	// err = db.Debug().Not("title", "ALI").First(&books).Error
	// if err != nil {
	// 	log.Fatal("============= QUERY ERROR =============")
	// }
	// println()
	// fmt.Println("DATA BOOK :",books)
	// println()

	// var books []entity.Book
	// err = db.Debug().Not("title", []string{"RIAND", "TRISULI"}).Find(&books).Error // NOT IN
	// if err != nil {
	// 	log.Fatal("============= QUERY ERROR =============")
	// }
	// println()
	// fmt.Println("DATA BOOK :",books)
	// println()

	// var books []entity.Book
	// err = db.Debug().Not([]int64{1,2,3}).Find(&books).Error // Not In slice of primary keys
	// if err != nil {
	// 	log.Fatal("============= QUERY ERROR =============")
	// }
	// println()
	// fmt.Println("DATA BOOK :",books)
	// println()

	// var books []entity.Book
	// err = db.Debug().Not("description LIKE ?", "%golang%").Find(&books).Error
	// if err != nil {
	// 	log.Fatal("============= QUERY ERROR =============")
	// }
	// println()
	// fmt.Println("DATA BOOK :",books)
	// println()

	// var books []entity.Book
	// err = db.Debug().Not("rating <= ?", 6).Find(&books).Error // PLAIN SQL
	// if err != nil {
	// 	log.Fatal("============= QUERY ERROR =============")
	// }
	// println()
	// fmt.Println("DATA BOOK :",books)
	// println()

	// var books []entity.Book
	// err = db.Debug().Not(entity.Book{Title: "RIAND"}).Not(entity.Book{Title: "HENDRS"}).Find(&books).Error // STRUCT
	// if err != nil {
	// 	log.Fatal("============= QUERY ERROR =============")
	// }
	// println()
	// fmt.Println("DATA BOOK :",books)
	// println()

	var books []entity.Book
	err = db.Debug().Not("description LIKE ?", "%golang%").Not("description LIKE ?", "%php%").Find(&books).Error
	if err != nil {
		log.Fatal("============= QUERY ERROR =============")
	}
	println()
	fmt.Println("DATA BOOK :", books)
	println()
}

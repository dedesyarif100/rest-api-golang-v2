package Conditions

import (
	"fmt"
	"log"
	"gorm.io/gorm"
	"rest-api-golang-v2/entity"
)

func StringConditions(db *gorm.DB, err any) {
	// var books []entity.Book
	// err = db.Debug().Where("title = ?", "RIAND").First(&books).Error
	// if err != nil {
	// 	log.Fatal("============= QUERY ERROR =============")
	// }
	// println()
	// fmt.Println("DATA BOOK :",books)
	// println()

	// var books []entity.Book
	// err = db.Debug().Where("title = ?", "RIAND").Find(&books).Error
	// if err != nil {
	// 	log.Fatal("============= QUERY ERROR =============")
	// }
	// println()
	// fmt.Println("DATA BOOK :",books)
	// println()

	// var books []entity.Book
	// err = db.Debug().Where("title != ?", "TRISULI").Find(&books).Error
	// if err != nil {
	// 	log.Fatal("============= QUERY ERROR =============")
	// }
	// println()
	// fmt.Println("DATA BOOK :",books)
	// println()

	// var books []entity.Book
	// err = db.Debug().Where("title IN (?)", []string{"TRISULI", "RIAND"}).Find(&books).Error
	// if err != nil {
	// 	log.Fatal("============= QUERY ERROR =============")
	// }
	// println()
	// fmt.Println("DATA BOOK :",books)
	// println()

	// var books []entity.Book
	// err = db.Debug().Where("description LIKE ?", "%golang%").Find(&books).Error
	// if err != nil {
	// 	log.Fatal("============= QUERY ERROR =============")
	// }
	// println()
	// fmt.Println("DATA BOOK :",books)
	// println()

	// var books []entity.Book
	// err = db.Debug().Where("price = ? AND rating = ?", "110000", "2").Find(&books).Error
	// if err != nil {
	// 	log.Fatal("============= QUERY ERROR =============")
	// }
	// println()
	// fmt.Println("DATA BOOK :",books)
	// println()

	// var books []entity.Book
	// err = db.Debug().Where("created_at > ?", "2022-05-27 22:00:00").Find(&books).Error
	// if err != nil {
	// 	log.Fatal("============= QUERY ERROR =============")
	// }
	// println()
	// fmt.Println("DATA BOOK :",books)
	// println()

	var books []entity.Book
	err = db.Debug().Where("created_at BETWEEN ? AND ?", "2022-06-17 16:48:36", "2022-07-15 19:04:56").Find(&books).Error
	if err != nil {
		log.Fatal("============= QUERY ERROR =============")
	}
	println()
	fmt.Println("DATA BOOK :", books)
	println()
}

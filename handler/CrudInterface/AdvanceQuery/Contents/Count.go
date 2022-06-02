package Contents

import (
	"fmt"
	// "log"
	// "rest-api-golang/entity"
	"gorm.io/gorm"
)

func Count(db *gorm.DB, err any) {
	// var count int64
	// var orders []entity.Order
	// db.Debug().Model(&orders).Count(&count)
	// fmt.Println("\nJUMLAH DATA ORDERS :",count)
	// println()


	// var count int64
	// var books []entity.Book
	// db.Debug().Where("description LIKE ?", "%vue%").Or("description LIKE ?", "%php%").Find(&books).Count(&count)
	// fmt.Println("\nJUMLAH =", count)
	// println()


	// var count int64
	// db.Debug().Model(&entity.Book{}).Where("description LIKE ?", "%buku%").Count(&count)
	// fmt.Println("\nJUMLAH DATA BOOKS YANG TERDAPAT DESKRIPSI PHP :",count)
	// println()


	// var count int64
	// db.Debug().Table("books").Count(&count)
	// fmt.Println("\nJUMLAH DATA BOOKS :",count)
	// println()


	// var count int64
	// db.Debug().Model(&entity.User{}).Distinct("gender").Count(&count)
	// fmt.Println("\nJUMLAH DATA BOOKS :",count)
	// println()


	var count int64
	db.Debug().Table("books").Select("count(distinct(created_at))").Count(&count)
	fmt.Println("\nJUMLAH DATA DISTICNT BOOKS :",count)
	println()
}
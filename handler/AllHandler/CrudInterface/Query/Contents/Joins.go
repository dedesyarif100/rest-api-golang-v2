package Contents

import (
	"fmt"
	"log"
	"rest-api-golang/entity"
	"gorm.io/gorm"
)

func Joins(db *gorm.DB, err any) {
	// var books []entity.Book
	// err = db.Debug().Model(&books).Select("books.title").Joins("JOIN orders ON orders.book_id = books.id").Scan(&books).Error
	// if err != nil {
	// 	log.Fatal("============= QUERY ERROR =============")
	// }
	// println()
	// fmt.Println("DATA BOOK :", books)
	// println()


	// var books []entity.Book
	// err = db.Debug().Table("books").Select("books.title").Joins("JOIN orders ON orders.book_id = books.id").Scan(&books).Error
	// if err != nil {
	// 	log.Fatal("============= QUERY ERROR =============")
	// }
	// println()
	// fmt.Println("DATA BOOK :", books)
	// println()
	// for _, data := range books {
	// 	fmt.Println("ID  	:",data)
	// }


	// multiple joins with parameter
		var books []entity.Book
		err = db.Debug().Joins("JOIN orders ON orders.book_id = books.id AND orders.state = ?", "UNKNOW").Joins("JOIN credit_cards ON credit_cards.book_id = books.id").Where("credit_cards.number = ?", "113").Find(&books).Error
		if err != nil {
			log.Fatal("============= QUERY ERROR =============")
		}
		println()
		fmt.Println("DATA BOOK :", books)
		println()
		

		// var books []entity.Book
		// err = db.Debug().Joins("JOIN orders ON orders.book_id = books.id AND orders.state = ?", "UNKNOW").Find(&books).Error
		// if err != nil {
		// 	log.Fatal("============= QUERY ERROR =============")
		// }
		// println()
		// fmt.Println("DATA BOOK :", books)
		// println()


		// var books []entity.Book
		// err = db.Debug().Joins("JOIN credit_cards ON credit_cards.book_id = books.id").Where("credit_cards.number = ?", "117").Find(&books).Error
		// if err != nil {
		// 	log.Fatal("============= QUERY ERROR =============")
		// }
		// println()
		// fmt.Println("DATA BOOK :", books)
		// println()
		
		
		// var books []entity.Book
		// err = db.Debug().Joins("JOIN credit_cards ON credit_cards.book_id = books.id AND credit_cards.number = ?", "116").Find(&books).Error
		// if err != nil {
		// 	log.Fatal("============= QUERY ERROR =============")
		// }
		// println()
		// fmt.Println("DATA BOOK :", books)
		// println()


	// Joins Preloading
		// var books []entity.Book
		// db.Debug().Joins("credit_cards").Find(&books)
		// println()
		// println(books)
		// println()

		// var users []entity.User
		// db.Joins("companies").Find(&users)
		// println()
		// println(users)
		// println()

		// var users []entity.User
		// db.Joins("companies", DB.Where(&entity.Company{Name: "COMPANY A"})).Find(&users)
		// println()
		// println(users)
		// println()
}
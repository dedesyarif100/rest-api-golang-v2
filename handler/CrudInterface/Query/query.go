package Query

import (
	// "errors"
	"fmt"
	"log"
	"rest-api-golang/entity"
	// "rest-api-golang/handler/CrudInterface/Query/Conditions"
	// "rest-api-golang/handler/CrudInterface/Query/GroupByAndHaving"
	// "rest-api-golang/handler/CrudInterface/Query/Joins"
	// "rest-api-golang/handler/CrudInterface/Query/LimitAndOffset"
	// "rest-api-golang/handler/CrudInterface/Query/Order"
	"rest-api-golang/handler/CrudInterface/Query/RetrievingAllObjects"
	"rest-api-golang/handler/CrudInterface/Query/RetrievingSingleObject"
	// "rest-api-golang/handler/CrudInterface/Query/Scan"
	"gorm.io/gorm"
)

func RetrievingSingleObject(db *gorm.DB, err any) {
	// var books []entity.Book
	// err = db.Debug().First(&books).Error
	// err = db.Debug().Take(&books).Error
	// err = db.Debug().Last(&books).Error
	// err = db.Debug().First(&books, 2).Error
	// err = db.Debug().Last(&books).Find(&books).Error
	// if err != nil {
		// log.Fatal("============= QUERY ERROR =============")
	// }
	// println()
	// fmt.Println("DATA BOOK :", books)
	// println()


	// var books []entity.Book
	// result := db.Debug().Find(&books)
	// println()
	// fmt.Println("COUNT OF RECORDS :",result.RowsAffected)
	// fmt.Println("ERROR :",result.Error)
	// fmt.Println(errors.Is(result.Error, gorm.ErrRecordNotFound))
	// println()


	// var books []entity.Book
	// result2 := map[string]interface{}{}
	// db.Model(&books).Find(&result2)
	// db.Table("books").Take(&result2)
	// println()
	// fmt.Println("DATA BOOKS :", books)
	// println()


	// var books []entity.Book
	// db.Debug().Find(&books, []int{1, 2, 3})
	// println()
	// fmt.Println("DATA BOOKS :", books)
	// println()


	// Retrieving objects with primary key
		// var books []entity.Book
		// db.Debug().First(&books, "2")
		// println()
		// fmt.Println("DATA BOOKS :", books)
		// println()


		// var books []entity.Book
		// db.Debug().Find(&books, []int{1, 2, 3})
		// println()
		// fmt.Println("DATA BOOKS :", books)
		// println()


		// val := entity.Book{ID: 4}
		// db.Debug().First(&val)
		// println()
		// fmt.Println("DATA BOOKS :", val)
		// println()


		val := []entity.Book{}
		db.Debug().Model(entity.Book{ID: 4}).First(&val)
		println()
		fmt.Println("DATA BOOKS :", val)
		println()
}

// func RetrievingAllObjects(db *gorm.DB, err any) {
// 	var books []entity.Book
// 	result := db.Debug().Find(&books)
// 	println()
// 	fmt.Println("DATA BOOKS :", books)
// 	println()
// 	fmt.Println("RESULT :", result.RowsAffected)
// 	println()
// }

func Conditions(db *gorm.DB, err any) {
	// String Conditions
		// StringConditions(db, err)

	// Struct & Map Conditions
		// StructAndMapConditions(db, err)

	// Not Conditions
		// NotConditions(db, err)

	// Or Conditions
		OrConditions(db, err)

	// Inline Condition
		// InlineConditions(db, err)
}

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
		err = db.Debug().Where("created_at BETWEEN ? AND ?", "2022-05-27 21:00:00", "2022-05-27 23:00:00").Find(&books).Error
		if err != nil {
			log.Fatal("============= QUERY ERROR =============")
		}
		println()
		fmt.Println("DATA BOOK :",books)
		println()
	}

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

	func SpecifyStructSearchFields() {

	}

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
		fmt.Println("DATA BOOK :",books)
		println()
	}

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
		fmt.Println("DATA BOOK :",books)
		println()
	}

	func OrConditions(db *gorm.DB, err any) {
		// var books []entity.Book
		// err = db.Debug().Where("title = ?", "HENDRS").Or("title = ?", "RIAND").Find(&books).Error
		// if err != nil {
		// 	log.Fatal("============= QUERY ERROR =============")
		// }
		// println()
		// fmt.Println("DATA BOOK :",books)
		// println()


		// var books []entity.Book
		// err = db.Debug().Where("description LIKE ?", "%php%").Or("description LIKE ?", "%golang%").Find(&books).Error
		// if err != nil {
		// 	log.Fatal("============= QUERY ERROR =============")
		// }
		// println()
		// fmt.Println("DATA BOOK :",books)
		// println()


		// var books []entity.Book
		// err = db.Debug().Where("title = 'RIAND'").Or(entity.Book{Title: "HENDRS"}).Find(&books).Error // STRUCT
		// if err != nil {
		// 	log.Fatal("============= QUERY ERROR =============")
		// }
		// println()
		// fmt.Println("DATA BOOK :",books)
		// println()


		var books []entity.Book
		err = db.Debug().Where("title = 'RIAND'").Or(map[string]interface{}{"title" : "HENDRS"}).Find(&books).Error // MAP
		if err != nil {
			log.Fatal("============= QUERY ERROR =============")
		}
		println()
		fmt.Println("DATA BOOK :",books)
		println()
	}

func SelectingSpecificFields() {

}

func Order(db *gorm.DB, err any) {
	// var books []entity.Book
	// err = db.Debug().Order("age desc").Find(&books).Error
	// if err != nil {
	// 	log.Fatal("============= QUERY ERROR =============")
	// }
	// println()
	// fmt.Println("DATA BOOK :",books)
	// println()


	// var books []entity.Book
	// err = db.Debug().Order("age desc").Order("title").Find(&books).Error
	// if err != nil {
	// 	log.Fatal("============= QUERY ERROR =============")
	// }
	// println()
	// fmt.Println("DATA BOOK :",books)
	// println()


	var books []entity.Book
	err = db.Debug().Order("age").Find(&books).Error
	if err != nil {
		log.Fatal("============= QUERY ERROR =============")
	}
	println()
	fmt.Println("DATA BOOK :",books)
	println()
}

func LimitAndOffset(db *gorm.DB, err any) {
	// var books []entity.Book
	// err = db.Debug().Limit(2).Find(&books).Error
	// if err != nil {
	// 	log.Fatal("============= QUERY ERROR =============")
	// }
	// println()
	// fmt.Println("DATA BOOK :",books)
	// println()
	

	var books []entity.Book
	err = db.Debug().Limit(2).Offset(3).Find(&books).Error
	if err != nil {
		log.Fatal("============= QUERY ERROR =============")
	}
	println()
	fmt.Println("DATA BOOK :",books)
	println()
}

func GroupByAndHaving(db *gorm.DB, err any) {
	var results []entity.Result
	err = db.Debug().Table("orders").Select("date(created_at) as date, sum(amount) as total").Group("date(created_at)").Having("sum(amount) >= ?", 20000).Scan(&results)
	if err != nil {
		log.Fatal("============= QUERY ERROR =============")
	}
	println()
	fmt.Println("DATA RESULTS :", results)
	println()
}

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
		// var books []entity.Book
		// err = db.Debug().Joins("JOIN orders ON orders.book_id = books.id AND orders.state = ?", "UNKNOW").Joins("JOIN credit_cards ON credit_cards.book_id = books.id").Where("credit_cards.number = ?", "113").Find(&books).Error
		// if err != nil {
		// 	log.Fatal("============= QUERY ERROR =============")
		// }
		// println()
		// fmt.Println("DATA BOOK :", books)
		// println()
		

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

func RunningQuery(db *gorm.DB, err any) {
	// var amount int
	// db.Debug().Raw("SELECT SUM(amount) FROM orders WHERE state = ?", "UNKNOW").Scan(&amount)
	// fmt.Println(amount)
	// QUERY
		// Retrieving a single object
			Retrievingsingleobject.RetrievingSingleObject(db, err)
		

		// Retrieving all objects
			RetrievingAllObjects.RetrievingAllObjects(db, err)

		// CONDITIONS
			Conditions(db, err)
			
		// Order
			// Order(db, err)

		// Limit & Offset
			// LimitAndOffset(db, err)

		// Group By & Having
			// GroupByAndHaving(db, err)

		// Joins
			// Joins(db, err)

		// Scan
			Scan(db, err)
}
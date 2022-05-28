package main

import (
	"fmt"
	"log"
	"rest-api-golang/book"
	"rest-api-golang/handler"
	"time"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

func main() {
	db, err := gorm.Open("mysql", "root:root@(127.0.0.1:43306)/rest-api-golang?charset=utf8&parseTime=True&loc=Local")
  	defer db.Close()

	if err != nil {
		log.Fatal("DB CONNECTION ERROR");
	}

	fmt.Println("DATABASE CONNECTION SUCCESS");

	// CRUD
	// CREATE DATA FROM MIGRATION ==============================================================================
		// TABLE BOOKS
			// db.AutoMigrate(&book.Book{}); // UNTUK MIGRATE COLUMN / TABEL
			// title := []string{"DEDE", "TRISULI", "HENDRS", "RIAND", "ALI", "MIRDAS"}
			// description := []string{
			// 	"buku php laravel",
			// 	"buku mysql golang",
			// 	"buku laravel mysql",
			// 	"buku golang php",
			// 	"php react",
			// 	"laravel vue",
			// }
			// price := []int{
			// 	110000,
			// 	115000,
			// 	120000,
			// 	125000,
			// 	130000,
			// 	135000,
			// }
			// rating := []int64{2, 4, 6, 8, 10, 12}
			// discount := []int64{12, 14, 16, 18, 20, 22}
			// // created_at := []string{
			// // 	"2022-05-27 20:04:36",
			// // 	"2022-05-27 21:04:36",
			// // 	"2022-05-27 22:04:36",
			// // 	"2022-05-27 23:04:36",
			// // 	"2022-05-28 00:04:36",
			// // 	"2022-05-28 01:04:36",
			// // }

			// for key, data := range title {
			// 	insertData := book.Book{}
			// 	insertData.Title = data
			// 	insertData.Description = description[key]
			// 	insertData.Price = price[key]
			// 	insertData.Rating = int(rating[key])
			// 	// insertData.CreatedAt = created_at[key]
			// 	insertData.Discount = int(discount[key])
			// 	err = db.Create(&insertData).Error
			// 	if err != nil {
			// 		fmt.Println("==========================")
			// 		fmt.Println("Error creating book record FROM MIGRATION")
			// 		fmt.Println("==========================")
			// 	}
			// }

		// TABLE ORDER
			// db.AutoMigrate(&book.Order{}); // UNTUK MIGRATE COLUMN / TABEL
			// amount := []int{10000, 20000, 30000}
			// state := []string{"PAID", "UNPAID", "UNKNOW"}
			// for key, data := range amount {
			// 	insertData := book.Order{}
			// 	insertData.Amount = data
			// 	insertData.State = state[key]
			// 	insertData.DeletedAt = time.Now()
			// 	err = db.Create(&insertData).Error
			// 	if err != nil {
			// 		fmt.Println("==========================")
			// 		fmt.Println("Error creating book record FROM MIGRATION")
			// 		fmt.Println("==========================")
			// 	}
			// }

	// insertData := book.Book{};
	// insertData.Title = "HENDRS"
	// insertData.Price = 110000
	// insertData.Discount = 17
	// insertData.Rating = 6
	// insertData.Description = "Buku 2"

	// err = db.Create(&insertData).Error
	// if err != nil {
	// 	fmt.Println("==========================")
	// 	fmt.Println("Error creating book record FROM MIGRATION")
	// 	fmt.Println("==========================")
	// }
	// CREATE DATA FROM MIGRATION ==============================================================================

	// GET DATA
	var books []book.Book
	// QUERY
		// err = db.Debug().First(&books).Error
		// err = db.Debug().Take(&books).Error
		// err = db.Debug().Last(&books).Error
		// err = db.Debug().First(&books, 3).Error
		// err = db.Debug().Last(&books).Find(&books).Error

	// WHERE
		// err = db.Debug().Where("title = ?", "HENDRS").First(&books).Error
		// err = db.Debug().Where("title = ?", "RIAND").Find(&books).Error
		// err = db.Debug().Where("title != ?", "TRISULI").Find(&books).Error
		// err = db.Debug().Where("title IN (?)", []string{"TRISULI", "RIAND"}).Find(&books).Error
		// err = db.Debug().Where("description LIKE ?", "%golang%").Find(&books).Error
		// err = db.Debug().Where("price = ? AND rating = ?", "110000", "2").Find(&books).Error
		// err = db.Debug().Where("created_at > ?", "2022-05-27 22:00:00").Find(&books).Error
		// err = db.Debug().Where("created_at BETWEEN ? AND ?", "2022-05-27 21:00:00", "2022-05-27 23:00:00").Find(&books).Error

	// Struct & Map
		// err = db.Debug().Where(map[string]interface{}{"title": "HENDRS", "price": "150000"}).Find(&books).Error // Map
		// err = db.Debug().Where([]int64{2, 4}).Find(&books).Error; // Slice of primary keys

	// NOT
		// err = db.Debug().Not("title", "ALI").First(&books).Error
		// err = db.Debug().Not("title", []string{"RIAND", "TRISULI"}).Find(&books).Error // NOT IN
		// err = db.Debug().Not([]int64{1,2,3}).Find(&books).Error // Not In slice of primary keys
		// err = db.Debug().Not("description LIKE ?", "%golang%").Find(&books).Error
		// err = db.Debug().Not("rating <= ?", 6).Find(&books).Error // PLAIN SQL
		// err = db.Debug().Not(book.Book{Title: "RIAND"}).Not(book.Book{Title: "HENDRS"}).Find(&books).Error // STRUCT
		// err = db.Debug().Not("description LIKE ?", "%golang%").Not("description LIKE ?", "%php%").Find(&books).Error

	// OR
		// err = db.Debug().Where("title = ?", "HENDRS").Or("title = ?", "RIAND").Find(&books).Error
		// err = db.Debug().Where("description LIKE ?", "%php%").Or("description LIKE ?", "%golang%").Find(&books).Error
		// err = db.Debug().Where("title = 'RIAND'").Or(book.Book{Title: "HENDRS"}).Find(&books).Error // STRUCT
		// err = db.Debug().Where("title = 'RIAND'").Or(map[string]interface{}{"title" : "HENDRS"}).Find(&books).Error // MAP

	// Inline Condition
		// err = db.Debug().First(&books).Error
		// Plain SQL
		// err = db.Debug().Find(&books, "title = ?", "TRISULI").Error
		// err = db.Debug().Find(&books, "title <> ? AND rating >= ?", "RIAND", 4).Error
		// struct
		// err = db.Debug().Find(&books, book.Book{Rating: 6}).Error
		// Map
		// err = db.Debug().Find(&books, map[string]any{"rating": 6}).Error

	// FirstOrInit
		// err = db.Debug().Where(book.Book{Title: "RIAND"}).FirstOrInit(&books).Error
		// err = db.Debug().FirstOrInit(&books, map[string]interface{}{"title": "RIAND"}).Error

	// Attrs
		// err = db.Debug().Where(book.Book{Title: "non_existing"}).Attrs(book.Book{Rating: 2}).FirstOrInit(&books).Error
		// err = db.Debug().Where(book.Book{Title: "DEDE"}).Attrs(book.Book{Rating: 2}).FirstOrInit(&books).Error

	// Assign
		// err = db.Debug().Where(book.Book{Title: "non_existing"}).Assign(book.Book{Rating: 4}).FirstOrInit(&books).Error
		// err = db.Debug().Where(book.Book{Title: "DEDE"}).Assign(book.Book{Rating: 111}).FirstOrInit(&books).Error

	// FirstOrCreate
		// err = db.Debug().FirstOrCreate(&books, book.Book{Title: "non_existing"}).Error
		// err = db.Debug().Where(book.Book{Title: "DEDE"}).FirstOrCreate(&books).Error

	// Advanced Query
		// SubQuery
			// err = db.Debug().Where()

	if err != nil {
		fmt.Println("==========================")
		fmt.Println("Error creating book record")
		fmt.Println("==========================")
	}

	// fmt.Println("ID 		:", books.ID)
	// fmt.Println("TITLE 		:", books.Title)
	// fmt.Println("DESCRIPTION 	:", books.Description)
	// fmt.Println("PRICE 		:", books.Price)
	// fmt.Println("RATING 		:", books.Rating)
	// fmt.Println("DISCOUNT 	:", books.Discount)
	// fmt.Println("BOOK OBJECT  	:", books)
	// fmt.Println("-----------------------")

	// GET ALL DATA -----------------------------
	for _, b := range books {
		fmt.Println("ID 		:", b.ID)
		fmt.Println("TITLE 		:", b.Title)
		fmt.Println("DESCRIPTION 	:", b.Description)
		fmt.Println("PRICE 		:", b.Price)
		fmt.Println("RATING 		:", b.Rating)
		fmt.Println("DISCOUNT 	:", b.Discount)
		fmt.Println("BOOK OBJECT  	:", b)
		fmt.Println("-----------------------")
	}
	// GET ALL DATA -----------------------------

	// db.AutoMigrate(&book.Assets{});

	router := gin.Default();

	v1 := router.Group("/v1");

	v1.GET("/", handler.RootHandler);
	v1.GET("/hello", handler.HelloHandler);
	v1.GET("/books/:id/:title", handler.BooksHandler);
	v1.GET("/query", handler.QueryHandler);
	v1.POST("/books", handler.BooksPostHandler);

	v2 := router.Group("/v2");

	v2.GET("/", handler.RootHandler);
	v2.GET("/hello", handler.HelloHandler);
	v2.GET("/books/:id/:title", handler.BooksHandler);
	v2.GET("/query", handler.QueryHandler);
	v2.POST("/books", handler.BooksPostHandler);

	router.Run(":8080") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

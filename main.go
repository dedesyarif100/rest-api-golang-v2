package main

import (
	"fmt"
	"log"
	"rest-api-golang/book"
	"rest-api-golang/handler"
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
	db.AutoMigrate(&book.Book{});

	// CRUD
	// CREATE DATA
	// insertData := book.Book{};
	// insertData.Title = "HENDRS"
	// insertData.Price = 110000
	// insertData.Discount = 17
	// insertData.Rating = 6
	// insertData.Description = "Buku 2"

	// err = db.Create(&insertData).Error
	// if err != nil {
	// 	fmt.Println("==========================")
	// 	fmt.Println("Error creating book record")
	// 	fmt.Println("==========================")
	// }

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
	err = db.Debug().Where("title IN (?)", []string{"TRISULI", "RIAND"}).Find(&books).Error
	// err = db.Debug().Where("description LIKE ?", "%golang%").Find(&books).Error
	// err = db.Debug().Where("price = ? AND rating = ?", "110000", "2").Find(&books).Error
	// err = db.Debug().Where("created_at > ?", "2022-05-27 22:00:00").Find(&books).Error
	// err = db.Debug().Where("created_at BETWEEN ? AND ?", "2022-05-27 21:00:00", "2022-05-27 23:00:00").Find(&books).Error

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


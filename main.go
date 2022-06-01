package main

import (
	// "errors"
	// "database/sql"
	"fmt"
	"log"

	// "net/http"
	// "rest-api-golang/book"
	// "rest-api-golang/handler"

	// "rest-api-golang/handler/CrudInterface/AdvanceQuery"
	// "rest-api-golang/handler/CrudInterface/Create"
	// "rest-api-golang/handler/CrudInterface/Delete"
	"rest-api-golang/handler/CrudInterface/Query"
	// "rest-api-golang/handler/CrudInterface/RawSqlAndSqlBuilder"

	// "rest-api-golang/migrations"
	"github.com/gin-gonic/gin"
	// _ "github.com/go-sql-driver/mysql"
	// "github.com/jinzhu/gorm"
	// "time"
	// "reflect"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	// db, err := gorm.Open("mysql", "root:root@(127.0.0.1:43306)/rest-api-golang?charset=utf8&parseTime=True&loc=Local")
  	// defer db.Close()
	// if err != nil {
	// 	log.Fatal("DB CONNECTION ERROR");
	// }
	// fmt.Println("DATABASE CONNECTION SUCCESS");

	dsn := "root:root@(127.0.0.1:43306)/rest-api-golang?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("DB CONNECTION ERROR");
	}
	fmt.Println("SUCCESS");
	// fmt.Println(reflect.TypeOf(err))
	// return;

	// AdvanceQuery.RunningQuery(db)
	// Create.RunningQuery(db)
	// Delete.RunningQuery(db)
	// RawSqlAndSqlBuilder.RunningQuery(db)

	// CRUD
	// CREATE DATA FROM MIGRATION ==============================================================================
		// var users []book.User
		// var books []book.Book
		// TABLE BOOKS
			// migrations.CreateBooksTable(db, err)

		// TABLE COMPANY
			// migrations.CreateCompanyTable(db, err)

		// TABLE USER
			// migrations.CreateUsersTable(db, err)

		// TABLE API USER
			// migrations.CreateApiUsersTable(db, err)

		// TABLE ORDER
			// migrations.CreateOrdersTable(db, err)

		// TABLE CREDIT CARDS
			// migrations.CreateCreditCardsTable(db, err)

		// TABLE PIZZAS
			// migrations.CreatePizzasTable(db, err)

		// TABLE RESULT
			// migrations.CreateResultsTable(db, err)
			// db.AutoMigrate(&book.Result{})

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
	// var books []book.Book
		// Select
			// err = db.Debug().Select("id, title").Find(&books).Error
			// err = db.Debug().Select([]string{"id", "title"}).Find(&books).Error
			// for _, b := range books {
			// 	fmt.Println(b.ID, b.Title)
			// }

	// CRUD INTERFACE
		// ** Create **

		// ** Query **
			Query.RunningQuery(db, err)

		// ** Advanced Query **

		// ** Update **

		// ** Delete **

		// ** Raw SQL & SQL Builder **
	
	// if err != nil {
	// 	fmt.Println("==========================")
	// 	fmt.Println("Error creating book record")
	// 	fmt.Println("==========================")
	// }

	// fmt.Println("ID 		:", books.ID)
	// fmt.Println("TITLE 		:", books.Title)
	// fmt.Println("DESCRIPTION 	:", books.Description)
	// fmt.Println("PRICE 		:", books.Price)
	// fmt.Println("RATING 		:", books.Rating)
	// fmt.Println("DISCOUNT 	:", books.Discount)
	// fmt.Println("BOOK OBJECT  	:", books)
	// fmt.Println("-----------------------")

	// GET ALL DATA BOOKS -----------------------------
		// fmt.Println("JUMLAHNYA ADALAH :",count)
		// for _, b := range books {
		// 	fmt.Println("ID 		:", b.ID)
		// 	fmt.Println("TITLE 		:", b.Title)
		// 	fmt.Println("AGE 		:", b.Age)
		// 	fmt.Println("DESCRIPTION 	:", b.Description)
		// 	fmt.Println("PRICE 		:", b.Price)
		// 	fmt.Println("RATING 		:", b.Rating)
		// 	fmt.Println("DISCOUNT 	:", b.Discount)
		// 	fmt.Println("BOOK OBJECT  	:", b)
		// 	fmt.Println("------------------------------------")
		// }
	// GET ALL DATA BOOKS -----------------------------

	// db.AutoMigrate(&book.Assets{});

	// LIST EXAMPLE API ============================================================================
	router := gin.Default();
	// 	v1 := router.Group("/v1");
	// 	v1.GET("/", handler.RootHandler);
	// 	v1.GET("/hello", handler.HelloHandler);
	// 	v1.GET("/books/:id/:title", handler.BooksHandler);
	// 	v1.GET("/query", handler.QueryHandler);
	// 	v1.POST("/books", handler.BooksPostHandler);

	// 	v2 := router.Group("/v2");
	// 	v2.GET("/", handler.RootHandler);
	// 	v2.GET("/hello", handler.HelloHandler);
	// 	v2.GET("/books/:id/:title", handler.BooksHandler);
	// 	v2.GET("/query", handler.QueryHandler);
	// 	v2.POST("/books", handler.BooksPostHandler);
	// LIST EXAMPLE API ============================================================================

	// manipulate := router.Group("/createExampleOne");
	// GET ALL DATA BOOK ============================================================================
		// manipulate.GET("/book", handler.GetAllBook);

	// CREATE DATA ============================================================================
		// manipulate.POST("/book", handler.CreateBook);

	// UPDATE DATA ============================================================================
		// manipulate.PUT("/book/:id", handler.UpdateBook);

	// DELETE DATA ============================================================================
		// manipulate.DELETE("/book/:id", handler.DeleteBook);

		router.Run(":8888") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
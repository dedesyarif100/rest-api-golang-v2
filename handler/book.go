package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"fmt"
	"log"
	"rest-api-golang/entity"
	"rest-api-golang/book"
	"github.com/go-playground/validator/v10"
	"github.com/jinzhu/gorm"
)

func RootHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"bio": "SOFTWARE ENGINEER",
		"message": "pong",
		"name": "DEDE SYARIFUDIN",
	});
}

func HelloHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"subtitle": "THIS IS A SUBTITLE",
		"title": "THIS IS A TITLE",
	});
}

func BooksHandler(c *gin.Context) {
	id := c.Param("id");
	fmt.Println("CEK ID :",id)
	title := c.Param(("title"));
	c.JSON(http.StatusOK, gin.H{
		"id": id,
		"title": title,
	});
}

func QueryHandler(c *gin.Context) {
	content := c.Query("content");
	title := c.Query("title");

	for i := 0; i <= 2; i++ {
		c.JSON(http.StatusOK, gin.H{
			"content": content,
			"title": title,
		},);
	}
}

func BooksPostHandler(c *gin.Context) {
	var bookInput book.BookInput
	
	err := c.ShouldBindJSON(&bookInput)
	// fmt.Println("CEK DATA :",bookInput.Title)
	if err != nil {
		fmt.Println("MASUK :", err);

		errorMessages := []string{}
		for _, e := range err.(validator.ValidationErrors) {
			msg := fmt.Sprintf("Error on field %s, condition: %s", e.Field(), e.ActualTag());
			errorMessages = append(errorMessages, msg);
			// c.JSON(http.StatusBadRequest, msg);
			// return;
		}

		c.JSON(http.StatusBadRequest, gin.H{
			"errors": errorMessages,
		});
		return;
	}

	c.JSON(http.StatusOK, gin.H{
		"title": 		bookInput.Title,
		"price": 		bookInput.Price,
		"sub_title": 	bookInput.SubTitle,
		"isactive": 	bookInput.IsActive,
	});
}

// CREATE UPDATE DELETE
// GET ALL BOOK
func GetAllBook(c *gin.Context) {
	db, err := gorm.Open("mysql", "root:root@(127.0.0.1:43306)/rest-api-golang?charset=utf8&parseTime=True&loc=Local")
  	defer db.Close()
	if err != nil {
		log.Fatal("DB CONNECTION ERROR");
	}
	fmt.Println("DATABASE CONNECTION SUCCESS");

	var books []entity.Book
	var count int
	db.Debug().Find(&books)
	db.Debug().Find(&books).Count(&count)

	// for _, val := range books {
	// 	c.JSON(http.StatusOK, gin.H {
	// 		"id"			: val.ID,
	// 		"title"			: val.Title,
	// 		"age"			: val.Age,
	// 		"description"	: val.Description,
	// 		"price"			: val.Price,
	// 		"rating"		: val.Rating,
	// 		"discount"		: val.Discount,
	// 		"created_at"	: val.CreatedAt,
	// 		"updated_at"	: val.UpdatedAt,
	// 	})
	// }

	for val := 0; val < count; val++ {
		c.JSON(http.StatusOK, gin.H {
			"id"			: books[val].ID,
			"title"			: books[val].Title,
			"age"			: books[val].Age,
			"description"	: books[val].Description,
			"price"			: books[val].Price,
			"rating"		: books[val].Rating,
			"discount"		: books[val].Discount,
			"created_at"	: books[val].CreatedAt,
			"updated_at"	: books[val].UpdatedAt,
		})
	}
}

// CREATE BOOK
func CreateBook(c *gin.Context) {
	db, err := gorm.Open("mysql", "root:root@(127.0.0.1:43306)/rest-api-golang?charset=utf8&parseTime=True&loc=Local")
  	defer db.Close()
	if err != nil {
		log.Fatal("DB CONNECTION ERROR");
	}
	fmt.Println("DATABASE CONNECTION SUCCESS");

	var book entity.Book
	c.ShouldBindJSON(&book)
	db.Create(&book)
}

// UPDATE BOOK
func UpdateBook(c *gin.Context) {
	db, err := gorm.Open("mysql", "root:root@(127.0.0.1:43306)/rest-api-golang?charset=utf8&parseTime=True&loc=Local")
  	defer db.Close()
	if err != nil {
		log.Fatal("DB CONNECTION ERROR");
	}
	fmt.Println("DATABASE CONNECTION SUCCESS");

	var books entity.Book
	id := c.Param("id");
	db.Debug().First(&books, id)
	c.ShouldBindJSON(&books)

	fmt.Println("CEK BOOK :",books)
	db.Save(&books)
	// books.Title = books.Title
}

// DELETE BOOK
func DeleteBook(c *gin.Context) {
	db, err := gorm.Open("mysql", "root:root@(127.0.0.1:43306)/rest-api-golang?charset=utf8&parseTime=True&loc=Local")
  	defer db.Close()
	if err != nil {
		log.Fatal("DB CONNECTION ERROR");
	}
	fmt.Println("DATABASE CONNECTION SUCCESS");

	var books entity.Book
	id := c.Param("id");
	db.Debug().Delete(&books, id)
}
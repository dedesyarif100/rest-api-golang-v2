package Api

import (
	"fmt"
	"log"
	"net/http"
	"gorm.io/gorm"
	"gorm.io/driver/mysql"
	"rest-api-golang/entity"
	"github.com/gin-gonic/gin"
)

// CREATE UPDATE DELETE
// GET ALL BOOK
func GetAllBook(c *gin.Context) {
	dsn := "root:root@(127.0.0.1:43306)/rest-api-golang?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("DB CONNECTION ERROR");
	}
	fmt.Println("DATABASE CONNECTION SUCCESS");

	var books []entity.Book
	var count int64
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

	var val int64
	for val = 0; val < count; val++ {
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
	dsn := "root:root@(127.0.0.1:43306)/rest-api-golang?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
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
	dsn := "root:root@(127.0.0.1:43306)/rest-api-golang?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
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
	dsn := "root:root@(127.0.0.1:43306)/rest-api-golang?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("DB CONNECTION ERROR");
	}
	fmt.Println("DATABASE CONNECTION SUCCESS");

	var books entity.Book
	id := c.Param("id");
	db.Debug().Delete(&books, id)
}
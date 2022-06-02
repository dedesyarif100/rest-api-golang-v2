package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"fmt"
	// "log"
	// "rest-api-golang/entity"
	"rest-api-golang/book"
	"github.com/go-playground/validator/v10"
	// "github.com/jinzhu/gorm"
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
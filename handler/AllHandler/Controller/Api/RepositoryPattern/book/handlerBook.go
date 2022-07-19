package book

import (
	"fmt"
	"strconv"
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type bookHandler struct {
	bookService Service
}

func NewBookHandler(bookService Service) *bookHandler {
	return &bookHandler{bookService}
}

func (handler *bookHandler) GetBooks(c *gin.Context) {
	books, err := handler.bookService.FindAll()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": err.Error(),
		})
		return
	}

	var booksResponse []BookResponse
	for _, book := range books {
		bookResponse := convertToBookResponse(book)
		booksResponse = append(booksResponse, bookResponse)
	}

	c.JSON(http.StatusOK, gin.H{
		"data": booksResponse,
	})
}


func (handler *bookHandler) GetBook(c *gin.Context) {
	idString := c.Param("id")
	id, _ := strconv.Atoi(idString)
	book, err := handler.bookService.FindById(int(id))
	
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": err.Error(),
		})
	} else {
		bookResponse := convertToBookResponse(book)
		c.JSON(http.StatusOK, gin.H{
			"data": bookResponse,
		})
	}
}


func (handler *bookHandler) CreateBook(c *gin.Context) {
	var bookRequest BookRequest
	err := c.ShouldBindJSON(&bookRequest)

	// fmt.Println("CEK DATA :",bookRequest.Title)
	if err != nil {
		fmt.Println("MASUK :", err);

		errorMessages := []string{}
		for _, e := range err.(validator.ValidationErrors) {
			msg := fmt.Sprintf("Error on field %s, condition: %s", e.Field(), e.ActualTag());
			errorMessages = append(errorMessages, msg);
		}

		c.JSON(http.StatusBadRequest, gin.H{
			"errors": errorMessages,
		});
		return;
	}

	book, err := handler.bookService.Create(bookRequest)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": convertToBookResponse(book),
	});
}


func (handler *bookHandler) UpdateBook(c *gin.Context) {
	var bookRequest BookRequest
	err := c.ShouldBindJSON(&bookRequest)

	// fmt.Println("CEK DATA :",bookRequest.Title)
	if err != nil {
		fmt.Println("MASUK :", err);

		errorMessages := []string{}
		for _, e := range err.(validator.ValidationErrors) {
			msg := fmt.Sprintf("Error on field %s, condition: %s", e.Field(), e.ActualTag());
			errorMessages = append(errorMessages, msg);
		}

		c.JSON(http.StatusBadRequest, gin.H{
			"errors": errorMessages,
		});
		return;
	}

	idString := c.Param("id")
	id, _ := strconv.Atoi(idString)
	fmt.Println("UPDATE MASUK :", id)

	book, err := handler.bookService.Update(c, int(id), bookRequest)

	fmt.Println("MASUK HANDLER UPDATE :", book)
}


func (handler *bookHandler) DeleteBook(c *gin.Context) {
	idString := c.Param("id")
	id, _ := strconv.Atoi(idString)
	// fmt.Println("DELETE MASUK :", id)

	book, err := handler.bookService.Delete(int(id))

	fmt.Println("MASUK HANDLER BOOK :", err)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": convertToBookResponse(book),
	});
}


func convertToBookResponse(b Book) BookResponse {
	return BookResponse{
		ID			: b.ID,
		Title		: b.Title,
		Age			: b.Age,
		Price		: b.Price,
		Description	: b.Description,
		Rating		: b.Rating,
		Discount	: b.Discount,
	}
}

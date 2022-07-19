package book

import (
	"fmt"
	"net/http"
	"github.com/gin-gonic/gin"
)

type Service interface {
	FindAll() ([]Book, error)
	FindById(ID int) (Book, error)
	Create(bookRequest BookRequest) (Book, error)
	Update(c *gin.Context, ID int, bookRequest BookRequest) (Book, error)
	Delete(ID int) (Book, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) FindAll() ([]Book, error) {
	books, err := s.repository.FindAll()
	return books, err
}

func (s *service) FindById(ID int) (Book, error) {
	book, err := s.repository.FindById(ID)
	fmt.Println("MASUK SERVICE BOOK :", err)
	return book, err

	// fmt.Println("MASUK SERVICE BOOK :", err)
}


func (s *service) Create(bookRequest BookRequest) (Book, error) {
	age, _ := bookRequest.Age.Int64()
	price, _ :=	bookRequest.Price.Int64()
	rating, _ :=	bookRequest.Rating.Int64()
	discount, _ :=	bookRequest.Discount.Int64()
	fmt.Println("MASUK BOOK REQUEST :")

	book := Book{
		Title: bookRequest.Title,
		Age: int(age),
		Price: int(price),
		Description: bookRequest.Description,
		Rating: int(rating),
		Discount: int(discount),
	}

	newBook, err := s.repository.Create(book)
	return newBook, err
}


func (s *service) Update(c *gin.Context, ID int, bookRequest BookRequest) (Book, error) {
	book, err := s.repository.FindById(ID)

	if book.ID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": err.Error(),
		})
		return book, nil
	} else {
		age, _ := bookRequest.Age.Int64()
		price, _ :=	bookRequest.Price.Int64()
		rating, _ :=	bookRequest.Rating.Int64()
		discount, _ :=	bookRequest.Discount.Int64()
		
		book.Title 			= bookRequest.Title
		book.Age			= int(age)
		book.Price 			= int(price)
		book.Description 	= bookRequest.Description
		book.Rating 		= int(rating)
		book.Discount 		= int(discount)
		
		newBook, err := s.repository.Update(book)
		fmt.Println("MASUK SERVICE UPDATE : ", book.ID)

		c.JSON(http.StatusOK, gin.H{
			"data": convertToBookResponse(book),
		});
		return newBook, err
	}
}


func (s *service) Delete(ID int) (Book, error) {
	book, _ := s.repository.FindById(ID)
	newBook, err := s.repository.Delete(book)
	fmt.Println("MASUK SERVICE BOOK :", err)
	return newBook, err
}
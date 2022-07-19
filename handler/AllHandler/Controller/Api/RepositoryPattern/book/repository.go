package book

import (
	"fmt"

	"gorm.io/gorm"
)

type Repository interface {
	FindAll() ([]Book, error)
	FindById(ID int) (Book, error)
	Create(book Book) (Book, error)
	Update(book Book) (Book, error)
	Delete(book Book) (Book, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindAll() ([]Book, error) {
	var books []Book
	err := r.db.Find(&books).Error
	return books, err
}

func (r *repository) FindById(ID int) (Book, error) {
	var book Book
	err := r.db.First(&book, ID).Error
	fmt.Println("FIND BY ID, IN REPOSITORY :", err)
	return book, err
}

func (r *repository) Create(book Book) (Book, error) {
	err := r.db.Create(&book).Error
	return book, err
}

func (r *repository) Update(book Book) (Book, error) {
	fmt.Println("MASUK REPOSITORY UPDATE")
	err := r.db.Save(&book).Error
	// return book, nil
	return book, err
}

func (r *repository) Delete(book Book) (Book, error) {
	err := r.db.Delete(&book).Error
	fmt.Println("MASUK REPOSITORY BOOK :", err)
	return book, err
}
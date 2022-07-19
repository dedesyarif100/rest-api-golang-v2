package book

import (
	"encoding/json"
)

type BookRequest struct {
	Title 		string 		`json:"title" binding:"required"`
	Age			json.Number	`json:"age" binding:"required,number"`
	Price 		json.Number	`json:"price" binding:"required,number"`
	Description	string		`json:"description" binding:"required"`
	Rating		json.Number	`json:"rating" binding:"required,number"`
	Discount	json.Number	`json:"discount" binding:"required,number"`
}
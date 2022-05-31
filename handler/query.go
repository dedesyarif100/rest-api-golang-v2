package handler

import (
	"fmt"
	"gorm.io/gorm"
)

func RunningQuery(db *gorm.DB) {
	var amount int
	db.Debug().Raw("SELECT SUM(amount) FROM orders WHERE state = ?", "UNKNOW").Scan(&amount)
	fmt.Println(amount)
}
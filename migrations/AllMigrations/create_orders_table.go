package AllMigrations

import (
	"fmt"
	"rest-api-golang-v2/entity"
	"time"

	"gorm.io/gorm"
)

func CreateOrdersTable(db *gorm.DB, err any) {
	var users []entity.User
	var books []entity.Book
	db.AutoMigrate(&entity.Order{}) // UNTUK MIGRATE COLUMN / TABEL
	amount := []int{10000, 20000, 30000, 40000, 50000, 60000}
	state := []string{"PAID", "UNPAID", "UNKNOW", "PAID", "UNPAID", "UNKNOW"}
	for key, data := range amount {
		insertData := entity.Order{}
		db.Debug().Find(&users)
		insertData.UserId = users[key].ID
		db.Debug().Find(&books)
		insertData.BookId = books[key].ID
		insertData.Amount = data
		insertData.State = state[key]
		insertData.DeletedAt = time.Now()
		err = db.Create(&insertData).Error
		if err != nil {
			fmt.Println("==========================")
			fmt.Println("Error creating ORDERS record FROM MIGRATION")
			fmt.Println("==========================")
		}
	}
}

package migrations

import (
	"rest-api-golang/entity"
	"gorm.io/gorm"
	"fmt"
)

func CreateCreditCardsTable(db *gorm.DB, err any) {
	var books []entity.Book
	db.AutoMigrate(&entity.CreditCard{})
	nameCards := []string{"BRI", "BNI", "BCA", "MANDIRI", "CIMB", "DANAMON", "BANK JATIM"}
	numberCards := []int{111, 112, 113, 114, 115, 116, 117}
	for key, data := range nameCards {
		insertData := entity.CreditCard{}
		db.Debug().Find(&books)
		insertData.BookId = books[key].ID
		insertData.Name = data
		insertData.Number = numberCards[key]
		err = db.Create(&insertData).Error
		if err != nil {
			fmt.Println("==========================")
			fmt.Println("Error creating CREDIT CARDS record FROM MIGRATION")
			fmt.Println("==========================")
		}
	}
}
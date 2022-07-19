package Contents

import (
	"fmt"
	// "log"
	"rest-api-golang-v2/entity"

	"gorm.io/gorm"
)

func FirstOrInit(db *gorm.DB, err any) {
	var books []entity.Book
	db.Debug().FirstOrInit(&books, entity.Book{Title: "DEDE"})
	err = db.Debug().Where(entity.Book{Title: "RIAND"}).FirstOrInit(&books).Error
	err = db.Debug().FirstOrInit(&books, map[string]interface{}{"title": "RIAND"}).Error
	db.Debug().Where(entity.Book{Title: "DEDE"}).Attrs(entity.Book{Age: 22}).FirstOrInit(&books)
	db.Debug().Where(entity.Book{Title: "DEDE"}).Attrs("age", 22).FirstOrInit(&books)
	println()
	for _, data := range books {
		fmt.Println("---------------------------------")
		fmt.Println("ID  		:", data.ID)
		fmt.Println("TITLE	  	:", data.Title)
		fmt.Println("AGE  		:", data.Age)
		fmt.Println("DESCRIPTION  	:", data.Description)
		fmt.Println("PRICE  		:", data.Price)
		fmt.Println("RATING  	:", data.Rating)
		fmt.Println("DISCOUNT  	:", data.Discount)
		fmt.Println("---------------------------------")
	}
	println()

	// // Attrs
	// 	err = db.Debug().Where(entity.Book{Title: "non_existing"}).Attrs(entity.Book{Rating: 2}).FirstOrInit(&books).Error
	// 	err = db.Debug().Where(entity.Book{Title: "DEDE"}).Attrs(entity.Book{Rating: 2}).FirstOrInit(&books).Error

	// // Assign
	// 	err = db.Debug().Where(entity.Book{Title: "non_existing"}).Assign(entity.Book{Rating: 4}).FirstOrInit(&books).Error
	// 	err = db.Debug().Where(entity.Book{Title: "DEDE"}).Assign(entity.Book{Rating: 111}).FirstOrInit(&books).Error
}

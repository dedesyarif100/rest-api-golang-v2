package Contents

import (
	"fmt"
	// "log"
	"rest-api-golang/entity"
	"gorm.io/gorm"
)

func FirstOrCreate(db *gorm.DB, err any) {
	var books []entity.Book
	// db.Debug().FirstOrCreate(&books, entity.Book{Title: "non_existing"})
	err = db.Debug().Where(entity.Book{Title: "DEDE"}).FirstOrCreate(&books).Error
	// err = db.Debug().Where(entity.Book{Title: "DEDE"}).Attrs(entity.Book{Price: 2}).FirstOrCreate(&books).Error
	println()
	for _, data := range books {
		fmt.Println("---------------------------------")
		fmt.Println("ID  		:",data.ID)
		fmt.Println("TITLE	  	:",data.Title)
		fmt.Println("AGE  		:",data.Age)
		fmt.Println("DESCRIPTION  	:",data.Description)
		fmt.Println("PRICE  		:",data.Price)
		fmt.Println("RATING  	:",data.Rating)
		fmt.Println("DISCOUNT  	:",data.Discount)
		fmt.Println("---------------------------------")
	}
	println()
}
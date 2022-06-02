package Contents

import (
	"fmt"
	"rest-api-golang/entity"
	"gorm.io/gorm"
)

func RetrievingSingleObject(db *gorm.DB, err any) {
	// var books []entity.Book
	// err = db.Debug().First(&books).Error
	// err = db.Debug().Take(&books).Error
	// err = db.Debug().Last(&books).Error
	// err = db.Debug().First(&books, 2).Error
	// err = db.Debug().Last(&books).Find(&books).Error
	// if err != nil {
		// log.Fatal("============= QUERY ERROR =============")
	// }
	// println()
	// fmt.Println("DATA BOOK :", books)
	// println()


	// var books []entity.Book
	// result := db.Debug().Find(&books)
	// println()
	// fmt.Println("COUNT OF RECORDS :",result.RowsAffected)
	// fmt.Println("ERROR :",result.Error)
	// fmt.Println(errors.Is(result.Error, gorm.ErrRecordNotFound))
	// println()


	// var books []entity.Book
	// result2 := map[string]interface{}{}
	// db.Model(&books).Find(&result2)
	// db.Table("books").Take(&result2)
	// println()
	// fmt.Println("DATA BOOKS :", books)
	// println()


	// var books []entity.Book
	// db.Debug().Find(&books, []int{1, 2, 3})
	// println()
	// fmt.Println("DATA BOOKS :", books)
	// println()


	// Retrieving objects with primary key
		// var books []entity.Book
		// db.Debug().First(&books, "2")
		// println()
		// fmt.Println("DATA BOOKS :", books)
		// println()


		// var books []entity.Book
		// db.Debug().Find(&books, []int{1, 2, 3})
		// println()
		// fmt.Println("DATA BOOKS :", books)
		// println()


		// val := entity.Book{ID: 4}
		// db.Debug().First(&val)
		// println()
		// fmt.Println("DATA BOOKS :", val)
		// println()


		val := []entity.Book{}
		db.Debug().Model(entity.Book{ID: 4}).First(&val)
		println()
		fmt.Println("DATA BOOKS :", val)
		println()
}
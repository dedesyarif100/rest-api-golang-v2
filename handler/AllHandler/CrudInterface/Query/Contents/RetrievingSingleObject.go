package Contents

import (
	"fmt"
	"log"
	// "os"
	// "errors"
	"rest-api-golang-v2/entity"
	"gorm.io/gorm"
)

func RetrievingSingleObject(db *gorm.DB, err any) {
	var books []entity.Book
	err = db.Debug().First(&books).Error
	err = db.Debug().Take(&books).Error
	err = db.Debug().Last(&books).Error
	err = db.Debug().First(&books, 4).Error
	err = db.Debug().Last(&books).Find(&books).Error
	if err != nil {
		fmt.Println("ERROR :", err)
		log.Fatal("============= QUERY ERROR =============")
	}
	println()
	fmt.Println("DATA BOOK :", books)
	println()

	// var books []entity.Book
	// result := db.Debug().Find(&books)
	// println()
	// fmt.Println("COUNT OF RECORDS :",result.RowsAffected)
	// fmt.Println("ERROR :",result.Error)
	// fmt.Println(errors.Is(result.Error, gorm.ErrRecordNotFound))
	// println()

	// var books []entity.Book
	// result2 := map[string]interface{}{}
	// db.Debug().Model(&books).Find(&result2)
	// db.Debug().Table("books").Take(&result2)
	// println()
	// fmt.Println("DATA BOOKS :", result2)
	// println()

	// var books []entity.Book
	// db.Debug().Find(&books, []int{1, 2, 4})
	// println()
	// fmt.Println("DATA BOOKS :", books)
	// println()

	// // Retrieving objects with primary key
	// var books []entity.Book
	// db.Debug().First(&books, "2")
	// println()
	// fmt.Println("DATA BOOKS :", books)
	// println()

	// var books []entity.Book
	// db.Debug().Find(&books, []int{1, 2, 4})
	// println()
	// fmt.Println("DATA BOOKS :", books)
	// println()

	// val := entity.Book{ID: 4}
	// err = db.Debug().First(&val).Error
	// if err != nil {
	// 	fmt.Println("ERROR :", err)
	// 	os.Exit(1)
	// }
	// println()
	// fmt.Println("DATA BOOKS :", val)
	// println()

	// books := []entity.Book{}
	// db.Debug().First(&books, "id = ?", 5)
	// println()
	// fmt.Println("DATA BOOKS :", books)
	// println()

	// books := []entity.Book{}
	// db.Debug().First(&books, "title = ?", "ali")
	// println()
	// fmt.Println("DATA BOOKS :", books)
	// println()

	// books := []entity.Book{}
	// db.Debug().Find(&books, "title = ?", "ali")
	// println()
	// fmt.Println("DATA BOOKS :", books)
	// println()

	// books := []entity.Book{}
	// err = db.Debug().First(&books, "title = ?", "tes").Error
	// if err != nil {
	// 	fmt.Println("ERROR :", err)
	// 	os.Exit(1)
	// }
	// println()
	// fmt.Println("DATA BOOKS :", books)
	// println()

	// var result = entity.Book{ID: 5}
	// db.Debug().First(&result)
	// println()
	// fmt.Println("DATA BOOKS :", result)
	// println()

	// --------------------------------------------------------------------------
	// var result []entity.Book
	// db.Debug().Model(entity.Book{}).First(&result)
	// db.Debug().Model(entity.Book{}).First(&result, 4)
	// db.Debug().Model(entity.Book{}).First(&result, "id = ?", 4)
	// db.Debug().Model(entity.Book{}).First(&result, "title = ?", "ali")

	// db.Debug().Model(entity.Book{}).Find(&result)
	// db.Debug().Model(entity.Book{}).Find(&result, 4)
	// db.Debug().Model(entity.Book{}).Find(&result, []int{1, 2, 4})
	// db.Debug().Model(entity.Book{}).Find(&result, "title = ?", "ali")

	// db.Debug().Model(entity.Book{}).Take(&result)
	// db.Debug().Model(entity.Book{}).Take(&result, 4)
	// db.Debug().Model(entity.Book{}).Take(&result, "id = ?", 4)
	// db.Debug().Model(entity.Book{}).Take(&result, "title = ?", "ali")

	// db.Debug().Model(entity.Book{}).Last(&result)
	// db.Debug().Model(entity.Book{}).Last(&result, 4)
	// db.Debug().Model(entity.Book{}).Last(&result, "id = ?", 4)
	// db.Debug().Model(entity.Book{}).Last(&result, "title = ?", "ali")



	// db.Debug().Table("books").First(&result)
	// db.Debug().Table("books").First(&result, 4)
	// db.Debug().Table("books").First(&result, "id = ?", 4)
	// db.Debug().Table("books").First(&result, "title = ?", "ali")

	// db.Debug().Table("books").Find(&result)
	// db.Debug().Table("books").Find(&result, 4)
	// db.Debug().Table("books").Find(&result, []int{1, 2, 4})
	// db.Debug().Table("books").Find(&result, "title = ?", "ali")

	// db.Debug().Table("books").Take(&result)
	// db.Debug().Table("books").Take(&result, 4)
	// db.Debug().Table("books").Take(&result, "id = ?", 4)
	// db.Debug().Table("books").Take(&result, "title = ?", "ali")

	// db.Debug().Table("books").Last(&result)
	// db.Debug().Table("books").Last(&result, 4)
	// db.Debug().Table("books").Last(&result, "id = ?", 4)
	// db.Debug().Table("books").Last(&result, "title = ?", "ali")

	// println()
	// fmt.Println("DATA BOOKS :", result)
	// println()
	// --------------------------------------------------------------------------
}

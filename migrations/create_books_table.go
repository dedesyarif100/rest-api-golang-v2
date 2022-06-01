package migrations

import (
	"fmt"
	"rest-api-golang/entity"
	"gorm.io/gorm"
)

func CreateBooksTable(db *gorm.DB, err any) {
	db.AutoMigrate(&entity.Book{}); // UNTUK MIGRATE COLUMN / TABEL
	title := []string{"DEDE", "TRISULI", "HENDRS", "RIAND", "ALI", "MIRDAS", "MA'ARIF"}
	age := []int{22, 25, 23, 26, 21, 20, 24}
	description := []string{
		"buku php laravel",
		"buku mysql golang",
		"buku laravel mysql",
		"buku golang php",
		"php react",
		"laravel vue",
		"vue react",
	}
	price := []int{
		110000,
		115000,
		120000,
		125000,
		130000,
		135000,
		140000,
	}
	rating := []int64{2, 4, 6, 8, 10, 12, 14}
	discount := []int64{12, 14, 16, 18, 20, 22, 24}
	// created_at := []string{
	// 	"2022-05-27 20:04:36",
	// 	"2022-05-27 21:04:36",
	// 	"2022-05-27 22:04:36",
	// 	"2022-05-27 23:04:36",
	// 	"2022-05-28 00:04:36",
	// 	"2022-05-28 01:04:36",
	// }

	for key, data := range title {
		insertData := entity.Book{}
		insertData.Title = data
		insertData.Age = age[key]
		insertData.Description = description[key]
		insertData.Price = price[key]
		insertData.Rating = int(rating[key])
		// insertData.CreatedAt = created_at[key]
		insertData.Discount = int(discount[key])
		err = db.Create(&insertData).Error
		if err != nil {
			fmt.Println("==========================")
			fmt.Println("Error creating book record FROM MIGRATION")
			fmt.Println("==========================")
		}
	}
}
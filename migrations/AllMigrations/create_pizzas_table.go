package AllMigrations

import (
	"fmt"
	"rest-api-golang/entity"
	"gorm.io/gorm"
)

func CreatePizzasTable(db *gorm.DB, err any) {
	db.AutoMigrate(&entity.Pizza{})
	namePizza := []string{"pepperoni", "hawaiian", "hawaiian"}
	sizePizza := []string{"small", "medium", "xlarge"}
	for key, data := range namePizza {
		insertData := entity.Pizza{}
		insertData.Pizza = data
		insertData.Size = sizePizza[key]
		err = db.Create(&insertData).Error
		if err != nil {
			fmt.Println("==========================")
			fmt.Println("Error creating CREDIT CARDS record FROM MIGRATION")
			fmt.Println("==========================")
		}
	}
}
package migrations

import (
	"fmt"
	"rest-api-golang/entity"
	"gorm.io/gorm"
)

func CreateCompanyTable(db *gorm.DB, err any) {
	db.AutoMigrate(&entity.Company{}); // UNTUK MIGRATE COLUMN / TABEL
	name := []string{"COMPANY A", "COMPANY B", "COMPANY C", "COMPANY D", "COMPANY E", "COMPANY F"}
	for _, data := range name {
		insertData := entity.Company{}
		insertData.Name = data
		err = db.Create(&insertData).Error
		if err != nil {
			fmt.Println("==========================")
			fmt.Println("Error creating book record FROM MIGRATION")
			fmt.Println("==========================")
		}
	}
}
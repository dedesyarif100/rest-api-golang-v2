package AllMigrations

import (
	"fmt"
	"rest-api-golang/entity"
	"gorm.io/gorm"
)

func CreateApiUsersTable(db *gorm.DB, err any) {
	// var ApiUser []book.ApiUser
	db.AutoMigrate(&entity.ApiUser{})
	nameApiUser := []string{"API USER A", "API USER B", "API USER C", "API USER D", "API USER E", "API USER F"}
	for _, data := range nameApiUser {
		insertData := entity.ApiUser{}
		insertData.Name = data
		err = db.Create(&insertData).Error
		if err != nil {
			fmt.Println("==========================")
			fmt.Println("Error creating APIUSERS record FROM MIGRATION")
			fmt.Println("==========================")
		}
	}
}
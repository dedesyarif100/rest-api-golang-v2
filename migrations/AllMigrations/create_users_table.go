package AllMigrations

import (
	"fmt"
	"rest-api-golang-v2/entity"
	"gorm.io/gorm"
)

func CreateUsersTable(db *gorm.DB, err any) {
	var company []entity.Company
	db.AutoMigrate(&entity.User{})
	nameUser := []string{"USER A", "USER B", "USER C", "USER D", "USER E", "USER F"}
	gender := []string{"MALE", "FEMALE", "MALE", "FEMALE", "MALE", "FEMALE"}
	role := []string{"superadmin", "superadmin", "admin", "admin", "user", "user"}
	// ageUser := []int{22, 22}
	for key, data := range nameUser {
		insertData := entity.User{}
		db.Debug().Find(&company)
		insertData.CompanyId = company[key].ID
		insertData.Name = data
		insertData.Age = 22
		insertData.Gender = gender[key]
		insertData.Role = role[key]
		err = db.Create(&insertData).Error
		if err != nil {
			fmt.Println("==========================")
			fmt.Println("Error creating USERS record FROM MIGRATION")
			fmt.Println("==========================")
		}
	}
}

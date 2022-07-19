package Contents

import (
	"fmt"
	"log"
	"rest-api-golang-v2/entity"

	"gorm.io/gorm"
)

func SmartSelectFields(db *gorm.DB, err any) {
	var users []entity.User
	var apiUsers entity.ApiUser
	err = db.Debug().Model(&users).Limit(3).Find(&apiUsers).Error
	if err != nil {
		log.Fatal("============= QUERY ERROR =============")
	}
	println()
	fmt.Println("DATA USERS :", users)
	println()
}

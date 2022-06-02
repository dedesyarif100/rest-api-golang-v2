package Contents

import (
	"fmt"
	// "log"
	"rest-api-golang/entity"
	"gorm.io/gorm"
)

func InWithMultipleColumns(db *gorm.DB, err any) {
	var users []entity.User
	db.Debug().Where("(name, age, role) IN ?", [][]interface{}{{"USER A", 22, "superadmin"}, {"USER C", 22, "admin"}}).Find(&users)
	println()
	for _, data := range users {
		fmt.Println("----------------------------")
		fmt.Println("ID 		:",data.ID)
		fmt.Println("NAME 		:",data.Name)
		fmt.Println("AGE 		:",data.Age)
		fmt.Println("GENDER 		:",data.Gender)
		fmt.Println("ROLE 		:",data.Role)
		fmt.Println("----------------------------")
	}
	println()
}
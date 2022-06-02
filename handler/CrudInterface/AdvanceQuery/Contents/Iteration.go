package Contents

import (
	"fmt"
	// "log"
	"rest-api-golang/entity"
	"gorm.io/gorm"
)

func Iteration(db *gorm.DB, err any) {
	rows, err := db.Debug().Model(&entity.User{}).Where("name = ?", "USER A").Rows()
	defer rows.Close()
	println()
	for rows.Next() {
		var users []entity.User
		db.ScanRows(rows, &users)
		for _, data := range users {
			fmt.Println("---------------------------------")
			fmt.Println("ID  		:",data.ID)
			fmt.Println("NAME	  	:",data.Name)
			fmt.Println("AGE  		:",data.Age)
			fmt.Println("GENDER  	:",data.Gender)
			fmt.Println("ROLE  		:",data.Role)
			fmt.Println("---------------------------------")
		}
	}
	println()
}
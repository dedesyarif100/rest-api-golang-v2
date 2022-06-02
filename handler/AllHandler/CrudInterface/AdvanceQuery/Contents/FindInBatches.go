package Contents

import (
	"fmt"
	// "log"
	"rest-api-golang/entity"
	"gorm.io/gorm"
)

func FindInBatches(db *gorm.DB, err any) {
	var users []entity.User
	println()
	result := db.Debug().Where("name = ?", "USER A").FindInBatches(&users, 100, func(tx *gorm.DB, batch int) error {
		for _, data := range users {
			fmt.Println("---------------------------------")
			fmt.Println("ID  		:",data.ID)
			fmt.Println("NAME	  	:",data.Name)
			fmt.Println("AGE  		:",data.Age)
			fmt.Println("GENDER  	:",data.Gender)
			fmt.Println("ROLE  		:",data.Role)
			fmt.Println("---------------------------------")
		}

		tx.Save(&users)
		// fmt.Println(tx.RowsAffected)
		// fmt.Println(batch)
		return nil
	})
	println()
	fmt.Println("RESULT ERROR 		:", result.Error)        // returned error
	fmt.Println("RESULT ROWS AFFECTED	:", result.RowsAffected) // processed records count in all batches
	println()
}
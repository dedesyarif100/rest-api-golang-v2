package Contents

import (
	"fmt"
	// "log"
	"rest-api-golang/entity"
	"gorm.io/gorm"
)

func SubQuery(db *gorm.DB, err any) {
	// var orders []entity.Order
	// err = db.Debug().Find(&orders).Error
	// println()
	// fmt.Println("DATA : >>>>>>>>>>>>>>>>>>>>")
	// for _, order := range orders {
	// 	fmt.Println("-------------------------")
	// 	fmt.Println("ID		:", order.ID)
	// 	fmt.Println("AMOUNT		:", order.Amount)
	// 	fmt.Println("STATE		:", order.State)
	// 	fmt.Println("-------------------------")
	// }
	// println()


	// var orders []entity.Order
	// err = db.Debug().Where("amount > ?", db.Table("orders").Select("AVG(amount)").Where("state = ?", "PAID").SubQuery()).Find(&orders).Error
	// db.Debug().Where("amount > (?)", db.Table("orders").Select("AVG(amount)")).Find(&orders)
	// fmt.Println("DATA : >>>>>>>>>>>>>>>>>>>>")
	// // GET ALL DATA ORDERS -----------------------------
	// 	for _, order := range orders {
	// 		fmt.Println("ID 		:", order.ID)
	// 		fmt.Println("AMOUNT 		:", order.Amount)
	// 		fmt.Println("STATE 		:", order.State)
	// 		fmt.Println("-----------------------")
	// 	}
	// // GET ALL DATA ORDERS -----------------------------


	var users []entity.User
	subQuery := db.Debug().Select("AVG(age)").Where("name LIKE ?", "%USER%").Table("users")
	db.Debug().Select("AVG(age) as umur").Group("name").Having("AVG(umur) > (?)", subQuery).Find(&users)
	fmt.Println("DATA : >>>>>>>>>>>>>>>>>>>>")
	// GET ALL DATA USERS -----------------------------
		for _, user := range users {
			fmt.Println("ID 		:", user.ID)
			fmt.Println("NAME 		:", user.Name)
			fmt.Println("AGE 		:", user.Age)
			fmt.Println("GENDER 	:", user.Gender)
			fmt.Println("-----------------------")
		}
	// GET ALL DATA USERS -----------------------------
}
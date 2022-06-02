package RawSqlAndSqlBuilder

import "gorm.io/gorm"

func RawSqlAndSqlBuilder(db *gorm.DB, err any) {
	// Raw SQL & SQL Builder
		// var credit_cards []book.CreditCard
		// db.Debug().Raw("SELECT * FROM credit_cards WHERE name = ?", "BRI").First(&credit_cards)
		// for _, data := range credit_cards {
		// 	fmt.Println("----------------------------")
		// 	fmt.Println("ID 		:",data.ID)
		// 	fmt.Println("BOOK ID 	:",data.BookId)
		// 	fmt.Println("NAME 		:",data.Name)
		// 	fmt.Println("NUMBER 		:",data.Number)
		// 	fmt.Println("----------------------------")
		// }

		// var amount int
		// db.Debug().Raw("SELECT SUM(amount) FROM orders WHERE state = ?", "UNKNOW").Scan(&amount)
		// fmt.Println(amount)
		// var tes = "tes"
		

		// var users []book.User
		// db.Raw("UPDATE users SET name = ? WHERE age = ? RETURNING id, name", "jinzhu", 20).Scan(&users)
		// fmt.Println(users)

		// Exec with Raw SQL
			// db.Exec("DROP TABLE users")
			// fmt.Println("SUCCESS")
}
package AdvanceQuery

import "gorm.io/gorm"

func RunningQuery(db *gorm.DB) {
	// ADVANCE QUERY
		// Smart Select Fields
			// var users []book.User
			// var apiUsers book.ApiUser
			// db.Debug().Model(&users).Limit(3).Find(&apiUsers)
			// fmt.Println(users)

		// SubQuery
			// var orders []book.Order
			// err = db.Debug().Find(&orders).Error
			// for _, order := range orders {
			// 	fmt.Println("ID		:", order.ID)
			// 	fmt.Println("AMOUNT		:", order.Amount)
			// 	fmt.Println("STATE		:", order.State)
			// 	fmt.Println("-------------------------")
			// }

			// err = db.Debug().Where("amount > ?", db.Table("orders").Select("AVG(amount)").Where("state = ?", "PAID").SubQuery()).Find(&orders).Error
			// db.Debug().Where("amount > (?)", db.Table("orders").Select("AVG(amount)")).Find(&orders)
			// GET ALL DATA ORDERS -----------------------------
				// for _, order := range orders {
				// 	fmt.Println("ID 		:", order.ID)
				// 	fmt.Println("AMOUNT 		:", order.Amount)
				// 	fmt.Println("STATE 		:", order.State)
				// 	fmt.Println("-----------------------")
				// }
			// GET ALL DATA ORDERS -----------------------------

			// var users []book.User
			// subQuery := db.Debug().Select("AVG(age)").Where("name LIKE ?", "%USER%").Table("users")
			// db.Debug().Select("AVG(age) as umur").Group("name").Having("AVG(umur) > (?)", subQuery).Find(&users)
			// GET ALL DATA USERS -----------------------------
				// for _, user := range users {
				// 	fmt.Println("ID 		:", user.ID)
				// 	fmt.Println("NAME 		:", user.Name)
				// 	fmt.Println("AGE 		:", user.Age)
				// 	fmt.Println("GENDER 	:", user.Gender)
				// 	fmt.Println("-----------------------")
				// }
			// GET ALL DATA USERS -----------------------------

		// Group Conditions
			// var pizzas []book.Pizza
			// db.Debug().Where(
			// 	db.Where("pizza = ?", "pepperoni").Where(db.Where("size = ?", "small").Or("size = ?", "medium")),
			// ).Or(
			// 	db.Where("pizza = ?", "hawaiian").Where("size = ?", "xlarge"),
			// ).Find(&pizzas)
			// println()
			// for _, data := range pizzas {
			// 	fmt.Println("-----------------------")
			// 	fmt.Println("ID 		:",data.ID)
			// 	fmt.Println("PIZZA 		:",data.Pizza)
			// 	fmt.Println("SIZE 		:",data.Size)
			// 	fmt.Println("-----------------------")
			// }

		// IN with multiple columns
			// var users []book.User
			// db.Debug().Where("(name, age, role) IN ?", [][]interface{}{{"USER A", 22, "superadmin"}, {"USER C", 22, "admin"}}).Find(&users)
			// println()
			// for _, data := range users {
			// 	fmt.Println("----------------------------")
			// 	fmt.Println("ID 		:",data.ID)
			// 	fmt.Println("NAME 		:",data.Name)
			// 	fmt.Println("AGE 		:",data.Age)
			// 	fmt.Println("GENDER 		:",data.Gender)
			// 	fmt.Println("ROLE 		:",data.Role)
			// 	fmt.Println("----------------------------")
			// }

		// Named Argument
			// var users []book.User
			
			// db.Debug().Where("name = @nameuser OR name = @nameuser", sql.Named("nameuser", "USER B")).Find(&users)

			// db.Debug().Where("name = @nameuser AND age = @ageuser AND role = @roleuser", 
			// 	sql.Named("nameuser", "USER A"), 
			// 	sql.Named("ageuser", 22),
			// 	sql.Named("roleuser", "superadmin"),
			// ).Find(&users)
			
			// db.Debug().Where("name = @nameuser AND age = @ageuser AND role = @roleuser", 
			// 	map[string]interface{}{
			// 		"nameuser"	: "USER B",
			// 		"ageuser"	: 22,
			// 		"roleuser"	: "superadmin",
			// 	},
			// ).First(&users)
			// println()
			// for _, data := range users {
			// 	fmt.Println("----------------------------")
			// 	fmt.Println("ID 		:",data.ID)
			// 	fmt.Println("NAME 		:",data.Name)
			// 	fmt.Println("AGE 		:",data.Age)
			// 	fmt.Println("GENDER 		:",data.Gender)
			// 	fmt.Println("ROLE 		:",data.Role)
			// 	fmt.Println("----------------------------")
			// }

		// Find To Map
			// var users []book.User
			// result := map[string]interface{}{}
			// db.Debug().Model(&users).First(&result, "id = ?", 1)
			// // fmt.Println(result)
			// fmt.Println("----------------------------")
			// for key, data := range result {
			// 	fmt.Println(key,"  :",data)
			// }
			// fmt.Println("----------------------------")

			// var results []map[string]interface{}
			// db.Debug().Table("users").Find(&results)
			// for _, result := range results {
			// 	fmt.Println("-----------------------------------------------------")
			// 	for key, data := range result {
			// 		fmt.Println(key,"  :",data)
			// 	}
			// 	fmt.Println("-----------------------------------------------------")
			// }

		// FirstOrInit
			// var books []book.Book
			// db.Debug().FirstOrInit(&books, book.Book{Title: "DEDE"})
			// err = db.Debug().Where(book.Book{Title: "RIAND"}).FirstOrInit(&books).Error
			// err = db.Debug().FirstOrInit(&books, map[string]interface{}{"title": "RIAND"}).Error
			// db.Debug().Where(book.Book{Title: "DEDE"}).Attrs(book.Book{Age: 22}).FirstOrInit(&books)
			// db.Debug().Where(book.Book{Title: "DEDE"}).Attrs("age", 22).FirstOrInit(&books)
			// for _, data := range books {
			// 	fmt.Println("---------------------------------")
			// 	fmt.Println("ID  		:",data.ID)
			// 	fmt.Println("TITLE	  	:",data.Title)
			// 	fmt.Println("AGE  		:",data.Age)
			// 	fmt.Println("DESCRIPTION  	:",data.Description)
			// 	fmt.Println("PRICE  		:",data.Price)
			// 	fmt.Println("RATING  	:",data.Rating)
			// 	fmt.Println("DISCOUNT  	:",data.Discount)
			// 	fmt.Println("---------------------------------")
			// }

			// Attrs
				// err = db.Debug().Where(book.Book{Title: "non_existing"}).Attrs(book.Book{Rating: 2}).FirstOrInit(&books).Error
				// err = db.Debug().Where(book.Book{Title: "DEDE"}).Attrs(book.Book{Rating: 2}).FirstOrInit(&books).Error

			// Assign
				// err = db.Debug().Where(book.Book{Title: "non_existing"}).Assign(book.Book{Rating: 4}).FirstOrInit(&books).Error
				// err = db.Debug().Where(book.Book{Title: "DEDE"}).Assign(book.Book{Rating: 111}).FirstOrInit(&books).Error

		// FirstOrCreate
			// var books []book.Book
			// db.Debug().FirstOrCreate(&books, book.Book{Title: "non_existing"})
			// err = db.Debug().Where(book.Book{Title: "DEDE"}).FirstOrCreate(&books).Error
			// err = db.Debug().Where(book.Book{Title: "DEDE"}).Attrs(book.Book{Price: 2}).FirstOrCreate(&books).Error
			// for _, data := range books {
			// 	fmt.Println("---------------------------------")
			// 	fmt.Println("ID  		:",data.ID)
			// 	fmt.Println("TITLE	  	:",data.Title)
			// 	fmt.Println("AGE  		:",data.Age)
			// 	fmt.Println("DESCRIPTION  	:",data.Description)
			// 	fmt.Println("PRICE  		:",data.Price)
			// 	fmt.Println("RATING  	:",data.Rating)
			// 	fmt.Println("DISCOUNT  	:",data.Discount)
			// 	fmt.Println("---------------------------------")
			// }

		// Iteration
			// rows, err := db.Debug().Model(&book.User{}).Where("name = ?", "USER A").Rows()
			// defer rows.Close()
			// for rows.Next() {
			// 	var users []book.User
			// 	db.ScanRows(rows, &users)
			// 	for _, data := range users {
			// 		fmt.Println("---------------------------------")
			// 		fmt.Println("ID  		:",data.ID)
			// 		fmt.Println("NAME	  	:",data.Name)
			// 		fmt.Println("AGE  		:",data.Age)
			// 		fmt.Println("GENDER  	:",data.Gender)
			// 		fmt.Println("ROLE  		:",data.Role)
			// 		fmt.Println("---------------------------------")
			// 	}
			// }

		// FindInBatches
			// var users []book.User
			// result := db.Debug().Where("name = ?", "USER A").FindInBatches(&users, 100, func(tx *gorm.DB, batch int) error {
			// 	for _, data := range users {
			// 		fmt.Println("---------------------------------")
			// 		fmt.Println("ID  		:",data.ID)
			// 		fmt.Println("NAME	  	:",data.Name)
			// 		fmt.Println("AGE  		:",data.Age)
			// 		fmt.Println("GENDER  	:",data.Gender)
			// 		fmt.Println("ROLE  		:",data.Role)
			// 		fmt.Println("---------------------------------")
			// 	}

			// 	tx.Save(&users)
			// 	// fmt.Println(tx.RowsAffected)
			// 	// fmt.Println(batch)
			// 	return nil
			// })
			// fmt.Println("RESULT ERROR 		:", result.Error)        // returned error
			// fmt.Println("RESULT ROWS AFFECTED	:", result.RowsAffected) // processed records count in all batches

		// Count
			// var count int64
			// var orders []book.Order
			// db.Debug().Model(&orders).Count(&count)
			// fmt.Println("JUMLAH DATA ORDERS :",count)

			// var count int64
			// var books []book.Book
			// db.Debug().Where("description LIKE ?", "%vue%").Or("description LIKE ?", "%php%").Find(&books).Count(&count)
			// fmt.Println("JUMLAH =", count)

			// var count int64
			// db.Debug().Model(&book.Book{}).Where("description LIKE ?", "%buku%").Count(&count)
			// fmt.Println("JUMLAH DATA BOOKS YANG TERDAPAT DESKRIPSI PHP :",count)

			// var count int64
			// db.Debug().Table("books").Count(&count)
			// fmt.Println("JUMLAH DATA BOOKS :",count)

			// var count int64
			// db.Debug().Model(&book.User{}).Distinct("gender").Count(&count)
			// fmt.Println("JUMLAH DATA BOOKS :",count)

			// var count int64
			// db.Debug().Table("books").Select("count(distinct(created_at))").Count(&count)
			// fmt.Println("JUMLAH DATA DISTICNT BOOKS :",count)

		// Pluck
			// var ages []int
			// db.Debug().Find(&book.User{}).Pluck("age", &ages)
			// fmt.Println("VALUE AGES :",ages)

			// var names []string
			// db.Debug().Model(&book.User{}).Pluck("name", &names)
			// fmt.Println("VALUE NAMES :",names)

			// var id []int
			// db.Debug().Model(&book.User{}).Pluck("id", &id)
			// fmt.Println("VALUE ID :",id)

			// var names []string
			// db.Debug().Table("users").Pluck("name", &names)
			// fmt.Println("VALUE NAMES :",names)

			// db.Debug().Table("books").Pluck("title", &names)
			// fmt.Println("VALUE BOOKS :",names)

			// var books []book.Book
			// db.Debug().Select("title, description").Find(&books)
			// fmt.Println("VALUE BOOKS :",books[1].Title)
}
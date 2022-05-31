package main

import (
	// "errors"
	// "database/sql"
	"fmt"
	"log"

	// "net/http"
	// "rest-api-golang/book"
	"rest-api-golang/handler"

	"github.com/gin-gonic/gin"
	// _ "github.com/go-sql-driver/mysql"
	// "github.com/jinzhu/gorm"
	// "time"
	// "reflect"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	// db, err := gorm.Open("mysql", "root:root@(127.0.0.1:43306)/rest-api-golang?charset=utf8&parseTime=True&loc=Local")
  	// defer db.Close()
	// if err != nil {
	// 	log.Fatal("DB CONNECTION ERROR");
	// }
	// fmt.Println("DATABASE CONNECTION SUCCESS");

	dsn := "root:root@(127.0.0.1:43306)/rest-api-golang?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("DB CONNECTION ERROR");
	}
	fmt.Println("SUCCESS");

	// CRUD
	// CREATE DATA FROM MIGRATION ==============================================================================
		// var users []book.User
		// var books []book.Book
		// TABLE BOOKS
			// db.AutoMigrate(&book.Book{}); // UNTUK MIGRATE COLUMN / TABEL
			// title := []string{"DEDE", "TRISULI", "HENDRS", "RIAND", "ALI", "MIRDAS", "MA'ARIF"}
			// age := []int{22, 25, 23, 26, 21, 20, 24}
			// description := []string{
			// 	"buku php laravel",
			// 	"buku mysql golang",
			// 	"buku laravel mysql",
			// 	"buku golang php",
			// 	"php react",
			// 	"laravel vue",
			// 	"vue react",
			// }
			// price := []int{
			// 	110000,
			// 	115000,
			// 	120000,
			// 	125000,
			// 	130000,
			// 	135000,
			// 	140000,
			// }
			// rating := []int64{2, 4, 6, 8, 10, 12, 14}
			// discount := []int64{12, 14, 16, 18, 20, 22, 24}
			// // created_at := []string{
			// // 	"2022-05-27 20:04:36",
			// // 	"2022-05-27 21:04:36",
			// // 	"2022-05-27 22:04:36",
			// // 	"2022-05-27 23:04:36",
			// // 	"2022-05-28 00:04:36",
			// // 	"2022-05-28 01:04:36",
			// // }

			// for key, data := range title {
			// 	insertData := book.Book{}
			// 	insertData.Title = data
			// 	insertData.Age = age[key]
			// 	insertData.Description = description[key]
			// 	insertData.Price = price[key]
			// 	insertData.Rating = int(rating[key])
			// 	// insertData.CreatedAt = created_at[key]
			// 	insertData.Discount = int(discount[key])
			// 	err = db.Create(&insertData).Error
			// 	if err != nil {
			// 		fmt.Println("==========================")
			// 		fmt.Println("Error creating book record FROM MIGRATION")
			// 		fmt.Println("==========================")
			// 	}
			// }

		// TABLE USER
			// db.AutoMigrate(&book.User{});
			// nameUser := []string{"USER A", "USER B", "USER C", "USER D", "USER E", "USER F"}
			// gender := []string{"MALE", "FEMALE", "MALE", "FEMALE", "MALE", "FEMALE"}
			// role := []string{"superadmin", "superadmin", "admin", "admin", "user", "user"}
			// // ageUser := []int{22, 22}
			// for key, data := range nameUser {
			// 	insertData := book.User{}
			// 	insertData.Name = data
			// 	insertData.Age = 22
			// 	insertData.Gender = gender[key]
			// 	insertData.Role = role[key]
			// 	err = db.Create(&insertData).Error
			// 	if err != nil {
			// 		fmt.Println("==========================")
			// 		fmt.Println("Error creating USERS record FROM MIGRATION")
			// 		fmt.Println("==========================")
			// 	}
			// }

		// TABLE API USER
			// var ApiUser []book.ApiUser
			// db.AutoMigrate(&book.ApiUser{})
			// nameApiUser := []string{"API USER A", "API USER B", "API USER C", "API USER D", "API USER E", "API USER F"}
			// for _, data := range nameApiUser {
			// 	insertData := book.ApiUser{}
			// 	insertData.Name = data
			// 	err = db.Create(&insertData).Error
			// 	if err != nil {
			// 		fmt.Println("==========================")
			// 		fmt.Println("Error creating APIUSERS record FROM MIGRATION")
			// 		fmt.Println("==========================")
			// 	}
			// }

		// TABLE ORDER
			// db.AutoMigrate(&book.Order{}); // UNTUK MIGRATE COLUMN / TABEL
			// amount := []int{10000, 20000, 30000, 40000, 50000, 60000}
			// state := []string{"PAID", "UNPAID", "UNKNOW", "PAID", "UNPAID", "UNKNOW"}
			// for key, data := range amount {
			// 	insertData := book.Order{}
			// 	db.Debug().Find(&users)
			// 	insertData.UserId = users[key].ID
			// 	db.Debug().Find(&books)
			// 	insertData.BookId = books[key].ID
			// 	insertData.Amount = data
			// 	insertData.State = state[key]
			// 	insertData.DeletedAt = time.Now()
			// 	err = db.Create(&insertData).Error
			// 	if err != nil {
			// 		fmt.Println("==========================")
			// 		fmt.Println("Error creating ORDERS record FROM MIGRATION")
			// 		fmt.Println("==========================")
			// 	}
			// }

		// TABLE CREDIT CARDS
			// db.AutoMigrate(&book.CreditCard{})
			// nameCards := []string{"BRI", "BNI", "BCA", "MANDIRI", "CIMB", "DANAMON", "BANK JATIM"}
			// numberCards := []int{111, 112, 113, 114, 115, 116, 117}
			// for key, data := range nameCards {
			// 	insertData := book.CreditCard{}
			// 	db.Debug().Find(&books)
			// 	insertData.BookId = books[key].ID
			// 	insertData.Name = data
			// 	insertData.Number = numberCards[key]
			// 	err = db.Create(&insertData).Error
			// 	if err != nil {
			// 		fmt.Println("==========================")
			// 		fmt.Println("Error creating CREDIT CARDS record FROM MIGRATION")
			// 		fmt.Println("==========================")
			// 	}
			// }

		// TABLE PIZZAS
			// db.AutoMigrate(&book.Pizza{})
			// namePizza := []string{"pepperoni", "hawaiian", "hawaiian"}
			// sizePizza := []string{"small", "medium", "xlarge"}
			// for key, data := range namePizza {
			// 	insertData := book.Pizza{}
			// 	insertData.Pizza = data
			// 	insertData.Size = sizePizza[key]
			// 	err = db.Create(&insertData).Error
			// 	if err != nil {
			// 		fmt.Println("==========================")
			// 		fmt.Println("Error creating CREDIT CARDS record FROM MIGRATION")
			// 		fmt.Println("==========================")
			// 	}
			// }

		// TABLE RESULT
			// db.AutoMigrate(&book.Result{})

	// insertData := book.Book{};
	// insertData.Title = "HENDRS"
	// insertData.Price = 110000
	// insertData.Discount = 17
	// insertData.Rating = 6
	// insertData.Description = "Buku 2"

	// err = db.Create(&insertData).Error
	// if err != nil {
	// 	fmt.Println("==========================")
	// 	fmt.Println("Error creating book record FROM MIGRATION")
	// 	fmt.Println("==========================")
	// }
	// CREATE DATA FROM MIGRATION ==============================================================================

	// GET DATA
	// var books []book.Book
	// QUERY
		// Retrieving a single object
			// err = db.Debug().First(&books).Error
			// err = db.Debug().Take(&books).Error
			// err = db.Debug().Last(&books).Error
			// err = db.Debug().First(&books, 3).Error
			// err = db.Debug().Last(&books).Find(&books).Error

			// result := db.Debug().Find(&books)
			// fmt.Println("COUNT OF RECORDS :",result.RowsAffected)
			// fmt.Println("ERROR :",result.Error)
			// fmt.Println(errors.Is(result.Error, gorm.ErrRecordNotFound))

			// result2 := map[string]interface{}{}
			// db.Model(&books).Find(&result2)
			// db.Table("books").Take(&result2)
			// fmt.Println("DATA BOOKS :", books)

			// db.Debug().Find(&books, []int{1, 2, 3})
			// fmt.Println("DATA BOOKS :", books)

			// db.First(&book.Book{ID: 2})
			// fmt.Println("DATA BOOKS :", books)
			// fmt.Println("DATA BOOKS :", books)

			// Retrieving objects with primary key
				// db.Debug().First(&books, "2")
				// fmt.Println("DATA BOOKS :", books)

				// db.Debug().Find(&books, []int{1, 2, 3})
				// fmt.Println("DATA BOOKS :", books)

				// val := book.Book{ID: 2}
				// db.Debug().First(&val)
				// fmt.Println("DATA BOOKS :", val)

				// val := book.Book{}
				// db.Debug().Model(book.Book{ID: 2}).First(&val)
				// fmt.Println("DATA BOOKS :", val)

			// Retrieving all objects
				// result := db.Debug().Find(&books) 
				// fmt.Println("DATA BOOKS :", books)
				// println()
				// fmt.Println("RESULT :", result.RowsAffected)

				
			// type num struct {
			// 	ID int
			// }
			// val := time.Now()
			// fmt.Println("CEK TYPE : ", reflect.TypeOf(val))

	// String Conditions
		// err = db.Debug().Where("title = ?", "HENDRS").First(&books).Error
		// err = db.Debug().Where("title = ?", "RIAND").Find(&books).Error
		// err = db.Debug().Where("title != ?", "TRISULI").Find(&books).Error
		// err = db.Debug().Where("title IN (?)", []string{"TRISULI", "RIAND"}).Find(&books).Error
		// err = db.Debug().Where("description LIKE ?", "%golang%").Find(&books).Error
		// err = db.Debug().Where("price = ? AND rating = ?", "110000", "2").Find(&books).Error
		// err = db.Debug().Where("created_at > ?", "2022-05-27 22:00:00").Find(&books).Error
		// err = db.Debug().Where("created_at BETWEEN ? AND ?", "2022-05-27 21:00:00", "2022-05-27 23:00:00").Find(&books).Error

	// Struct & Map Conditions
		// err = db.Debug().Where(&book.Book{Title: "DEDE", Age: 22}).Find(&books).Error
		// fmt.Println(books)
		// err = db.Debug().Where(map[string]interface{}{"title": "HENDRS", "price": "150000"}).Find(&books).Error // Map
		// err = db.Debug().Where([]int64{2, 4}).Find(&books).Error; // Slice of primary keys

	// Not Conditions
		// err = db.Debug().Not("title", "ALI").First(&books).Error
		// err = db.Debug().Not("title", []string{"RIAND", "TRISULI"}).Find(&books).Error // NOT IN
		// err = db.Debug().Not([]int64{1,2,3}).Find(&books).Error // Not In slice of primary keys
		// err = db.Debug().Not("description LIKE ?", "%golang%").Find(&books).Error
		// err = db.Debug().Not("rating <= ?", 6).Find(&books).Error // PLAIN SQL
		// err = db.Debug().Not(book.Book{Title: "RIAND"}).Not(book.Book{Title: "HENDRS"}).Find(&books).Error // STRUCT
		// err = db.Debug().Not("description LIKE ?", "%golang%").Not("description LIKE ?", "%php%").Find(&books).Error

	// Or Conditions
		// err = db.Debug().Where("title = ?", "HENDRS").Or("title = ?", "RIAND").Find(&books).Error
		// err = db.Debug().Where("description LIKE ?", "%php%").Or("description LIKE ?", "%golang%").Find(&books).Error
		// err = db.Debug().Where("title = 'RIAND'").Or(book.Book{Title: "HENDRS"}).Find(&books).Error // STRUCT
		// err = db.Debug().Where("title = 'RIAND'").Or(map[string]interface{}{"title" : "HENDRS"}).Find(&books).Error // MAP

	// Inline Condition
		// err = db.Debug().First(&books).Error
		// Plain SQL
		// err = db.Debug().Find(&books, "title = ?", "TRISULI").Error
		// err = db.Debug().Find(&books, "title <> ? AND rating >= ?", "RIAND", 4).Error
		// struct
		// err = db.Debug().Find(&books, book.Book{Rating: 6}).Error
		// Map
		// err = db.Debug().Find(&books, map[string]any{"rating": 6}).Error

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

	// ADVANCE QUERY
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

		// Select
			// err = db.Debug().Select("id, title").Find(&books).Error
			// err = db.Debug().Select([]string{"id", "title"}).Find(&books).Error
			// for _, b := range books {
			// 	fmt.Println(b.ID, b.Title)
			// }

		// Order
			// err = db.Debug().Order("age desc").Find(&books).Error
			// err = db.Debug().Order("age desc").Order("title").Find(&books).Error
			// err = db.Debug().Order("age", true).Find(&books).Error

		// Limit & Offset
		
			// err = db.Debug().Limit(2).Find(&books).Error
			// err = db.Debug().Limit(2).Offset(3).Find(&books).Error

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

		// Group By & Having			
			// var results []book.Result
			// db.Debug().Table("orders").Select("date(created_at) as date, sum(amount) as total").Group("date(created_at)").Having("sum(amount) >= ?", 20000).Scan(&results)
			// println()
			// fmt.Println("DATA RESULTS :", results)
			// println()

		// Joins
			// db.Debug().Model(&books).Select("books.title").Joins("JOIN orders ON orders.book_id = books.id").Scan(&books)

			// db.Debug().Table("books").Select("books.title, orders.state").Joins("JOIN orders ON orders.book_id = books.id").Scan(&books)
			
			// multiple joins with parameter
				// db.Debug().Joins("JOIN orders ON orders.book_id = books.id AND orders.state = ?", "UNKNOW").Joins("JOIN credit_cards ON credit_cards.book_id = books.id").Where("credit_cards.number = ?", "113").Find(&books)
				// db.Debug().Joins("JOIN orders ON orders.book_id = books.id AND orders.state = ?", "UNKNOW").Find(&books)
				// db.Debug().Joins("JOIN credit_cards ON credit_cards.book_id = books.id").Where("credit_cards.number = ?", "117").Find(&books)
				// db.Debug().Joins("JOIN credit_cards ON credit_cards.book_id = books.id AND credit_cards.number = ?", "116").Find(&books)
				// println()
				// fmt.Println("DATA BOOKS :", books)
				// println()

		// Joins Preloading

			// println()
			// println( db.Debug().Joins("credit_cards").Find(&books) )

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

		// Scan
			// db.Debug().Table("books").Select("title").Where("title = ?", "dede").Scan(&books)
			// db.Debug().Raw("SELECT title, age FROM books WHERE title = ?", "RIAND").Scan(&books)
			// fmt.Println(books)

		// Smart Select Fields
			// var users []book.User
			// var apiUsers book.ApiUser
			// db.Debug().Model(&users).Limit(3).Find(&apiUsers)
			// fmt.Println(users)

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
		handler.RunningQuery(db)
		

		// var users []book.User
		// db.Raw("UPDATE users SET name = ? WHERE age = ? RETURNING id, name", "jinzhu", 20).Scan(&users)
		// fmt.Println(users)

	// Exec with Raw SQL
		// db.Exec("DROP TABLE users")
		// fmt.Println("SUCCESS")


	// if err != nil {
	// 	fmt.Println("==========================")
	// 	fmt.Println("Error creating book record")
	// 	fmt.Println("==========================")
	// }

	// fmt.Println("ID 		:", books.ID)
	// fmt.Println("TITLE 		:", books.Title)
	// fmt.Println("DESCRIPTION 	:", books.Description)
	// fmt.Println("PRICE 		:", books.Price)
	// fmt.Println("RATING 		:", books.Rating)
	// fmt.Println("DISCOUNT 	:", books.Discount)
	// fmt.Println("BOOK OBJECT  	:", books)
	// fmt.Println("-----------------------")

	// GET ALL DATA BOOKS -----------------------------
			// fmt.Println("JUMLAHNYA ADALAH :",count)
		// for _, b := range books {
		// 	fmt.Println("ID 		:", b.ID)
		// 	fmt.Println("TITLE 		:", b.Title)
		// 	fmt.Println("AGE 		:", b.Age)
		// 	fmt.Println("DESCRIPTION 	:", b.Description)
		// 	fmt.Println("PRICE 		:", b.Price)
		// 	fmt.Println("RATING 		:", b.Rating)
		// 	fmt.Println("DISCOUNT 	:", b.Discount)
		// 	fmt.Println("BOOK OBJECT  	:", b)
		// 	fmt.Println("------------------------------------")
		// }
	// GET ALL DATA BOOKS -----------------------------

	// db.AutoMigrate(&book.Assets{});

	// LIST EXAMPLE API ============================================================================
	router := gin.Default();
	// 	v1 := router.Group("/v1");
	// 	v1.GET("/", handler.RootHandler);
	// 	v1.GET("/hello", handler.HelloHandler);
	// 	v1.GET("/books/:id/:title", handler.BooksHandler);
	// 	v1.GET("/query", handler.QueryHandler);
	// 	v1.POST("/books", handler.BooksPostHandler);

	// 	v2 := router.Group("/v2");
	// 	v2.GET("/", handler.RootHandler);
	// 	v2.GET("/hello", handler.HelloHandler);
	// 	v2.GET("/books/:id/:title", handler.BooksHandler);
	// 	v2.GET("/query", handler.QueryHandler);
	// 	v2.POST("/books", handler.BooksPostHandler);
	// LIST EXAMPLE API ============================================================================

	// manipulate := router.Group("/createExampleOne");
	// GET ALL DATA BOOK ============================================================================
		// manipulate.GET("/book", handler.GetAllBook);

	// CREATE DATA ============================================================================
		// manipulate.POST("/book", handler.CreateBook);

	// UPDATE DATA ============================================================================
		// manipulate.PUT("/book/:id", handler.UpdateBook);

	// DELETE DATA ============================================================================
		// manipulate.DELETE("/book/:id", handler.DeleteBook);

		router.Run(":8080") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
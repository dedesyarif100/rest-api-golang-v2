package start

import (
	"fmt"
	"log"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func GetConnection() {
	dsn := "root:root@(127.0.0.1:43306)/rest-api-golang-v2?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("DB CONNECTION ERROR")
	}
	fmt.Println("SUCCESS")

	// <<<<<<<<<<<<< START EXECUTE QUERY >>>>>>>>>>>>
	StartExecuteQuery(db, err)
	// <<<<<<<<<<<<< START EXECUTE QUERY >>>>>>>>>>>>
}

package migrations

import (
	"rest-api-golang-v2/migrations/AllMigrations"
	"gorm.io/gorm"
)

func Migrations(db *gorm.DB, err any) {
	// TABLE BOOKS
	AllMigrations.CreateBooksTable(db, err)

	// TABLE COMPANY
	AllMigrations.CreateCompanyTable(db, err)

	// TABLE USER
	AllMigrations.CreateUsersTable(db, err)

	// TABLE API USER
	AllMigrations.CreateApiUsersTable(db, err)

	// TABLE ORDER
	AllMigrations.CreateOrdersTable(db, err)

	// TABLE CREDIT CARDS
	AllMigrations.CreateCreditCardsTable(db, err)

	// TABLE PIZZAS
	AllMigrations.CreatePizzasTable(db, err)

	// TABLE RESULT
	AllMigrations.CreateResultsTable(db, err)
}

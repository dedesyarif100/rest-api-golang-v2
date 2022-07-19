package start

import (
	"rest-api-golang-v2/handler"
	// "rest-api-golang-v2/migrations"
	"rest-api-golang-v2/router"
	"gorm.io/gorm"
)

func StartExecuteQuery(db *gorm.DB, err any) {
	// CREATE DATA FROM MIGRATION ==============================================================================
		// migrations.Migrations(db, err)
	// CREATE DATA FROM MIGRATION ==============================================================================

	// HANDLER ==============================================================================
		handler.Handler(db, err)
	// HANDLER ==============================================================================

	// START ROUTER ============================================================================
		router.Router(db, err)
	// START ROUTER ============================================================================
}

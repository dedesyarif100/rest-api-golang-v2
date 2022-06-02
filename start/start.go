package start

import (
	"gorm.io/gorm"
	"rest-api-golang/router"
	"rest-api-golang/migrations"
	"rest-api-golang/handler"
	"github.com/gin-gonic/gin"
)

func StartExecuteQuery(db *gorm.DB, err any) {
	// CREATE DATA FROM MIGRATION ==============================================================================
		migrations.Migrations(db, err)
	// CREATE DATA FROM MIGRATION ==============================================================================


	// HANDLER ==============================================================================
		handler.Handler(db, err)
	// HANDLER ==============================================================================


	// START ROUTER ============================================================================
	route := gin.Default();
	router.Router(route)
	route.Run(":9999") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
	// START ROUTER ============================================================================
}
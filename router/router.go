package router

import (
	"rest-api-golang-v2/router/AllRouter"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Router(db *gorm.DB, err any) {
	route := gin.Default()
	AllRouter.Books(route, db, err)
	// AllRouter.Companies(route)
	route.Run(":9999") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

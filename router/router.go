package router

import (
	"github.com/gin-gonic/gin"
	"rest-api-golang/router/AllRouter"
)

func Router(route *gin.Engine) {
	AllRouter.Books(route)
}
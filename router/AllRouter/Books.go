package AllRouter

import (
	"github.com/gin-gonic/gin"
	"rest-api-golang/handler/AllHandler/Controller/Api"
)

func Books(route *gin.Engine) {
	// v1 := route.Group("/v1");
	// 	v1.GET("/", handler.RootHandler);
	// 	v1.GET("/hello", handler.HelloHandler);
	// 	v1.GET("/books/:id/:title", handler.BooksHandler);
	// 	v1.GET("/query", handler.QueryHandler);
	// 	v1.POST("/books", handler.BooksPostHandler);

	// v2 := route.Group("/v2");
	// 	v2.GET("/", handler.RootHandler);
	// 	v2.GET("/hello", handler.HelloHandler);
	// 	v2.GET("/books/:id/:title", handler.BooksHandler);
	// 	v2.GET("/query", handler.QueryHandler);
	// 	v2.POST("/books", handler.BooksPostHandler);
	// LIST EXAMPLE API ============================================================================

	manipulate := route.Group("/createExampleOne");
	// // GET ALL DATA BOOK ============================================================================
		manipulate.GET("/book", Api.GetAllBook);

	// // CREATE DATA ============================================================================
		manipulate.POST("/book", Api.CreateBook);

	// // UPDATE DATA ============================================================================
		manipulate.PUT("/book/:id", Api.UpdateBook);

	// // DELETE DATA ============================================================================
		manipulate.DELETE("/book/:id", Api.DeleteBook);
}
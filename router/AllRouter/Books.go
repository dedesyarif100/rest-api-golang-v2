package AllRouter

import (
	// "rest-api-golang-v2/handler/AllHandler"
	// "rest-api-golang-v2/handler/AllHandler/Controller/Api"
	"rest-api-golang-v2/handler/AllHandler/Controller/Api/RepositoryPattern/book"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Books(route *gin.Engine, db *gorm.DB, err any) {
	// EXAMPLE USING ROUTES ----------------------------------------------
	// v1 := route.Group("/v2")
	// v1.GET("/", AllHandler.RootHandler)
	// v1.GET("/hello", AllHandler.HelloHandler)
	// v1.GET("/books/:id/:title", AllHandler.BooksHandler)
	// v1.GET("/query", AllHandler.QueryHandler)
	// v1.POST("/books", AllHandler.BooksPostHandler)

	// v2 := route.Group("/v2");
	// 	v2.GET("/", handler.RootHandler);
	// 	v2.GET("/hello", handler.HelloHandler);
	// 	v2.GET("/books/:id/:title", handler.BooksHandler);
	// 	v2.GET("/query", handler.QueryHandler);
	// 	v2.POST("/books", handler.BooksPostHandler);
	// EXAMPLE USING ROUTES ----------------------------------------------

	// ROUTING BOOK ============================================================================
		v2 := route.Group("/v2")
		handler := v2.Group("/handlerBook")
		bookRepository := book.NewRepository(db)
		bookService := book.NewService(bookRepository)
		bookHandler := book.NewBookHandler(bookService)

		handler.GET("/books", bookHandler.GetBooks)
		handler.GET("/book/:id", bookHandler.GetBook)
		handler.POST("/book", bookHandler.CreateBook)
		handler.PUT("/book/:id", bookHandler.UpdateBook)
		handler.DELETE("/book/:id", bookHandler.DeleteBook)
}
package routes

func (rc *RouteConfig) BookRoute() {
	api := rc.App.Group("/book", rc.AuthMiddleware.Auth)
	api.GET("", rc.BookHandler.GetBook)
	api.GET("/:id", rc.BookHandler.GetBookById)
	api.POST("/create", rc.BookHandler.CreateBook)
	api.POST("/update/:id", rc.BookHandler.UpdateBook)
	api.DELETE("/delete/:id", rc.BookHandler.DeleteBook)
}

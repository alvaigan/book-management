package routes

func (rc *RouteConfig) BookRoute() {
	api := rc.App.Group("/book", rc.AuthMiddleware.Auth)
	api.GET("", rc.Handler.GetBook)
	api.GET("/data/:id", rc.Handler.GetBookById)
	api.POST("/create", rc.Handler.CreateBook)
	api.POST("/update/:id", rc.Handler.UpdateBook)
	api.DELETE("/delete/:id", rc.Handler.DeleteBook)
}

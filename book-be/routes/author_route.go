package routes

func (rc *RouteConfig) AuthorRoute() {
	api := rc.App.Group("/author", rc.AuthMiddleware.Auth)
	api.GET("", rc.AuthorHandler.GetAuthor)
	api.GET("/:id", rc.AuthorHandler.GetAuthorById)
	api.POST("/create", rc.AuthorHandler.CreateAuthor)
	api.POST("/update/:id", rc.AuthorHandler.UpdateAuthor)
	api.DELETE("/delete/:id", rc.AuthorHandler.DeleteAuthor)
}

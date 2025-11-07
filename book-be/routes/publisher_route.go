package routes

func (rc *RouteConfig) PublisherRoute() {
	api := rc.App.Group("/publisher")
	api.GET("/", rc.Handler.Login)
	api.GET("/data/:id", rc.Handler.Login)
	api.POST("/create", rc.Handler.Register)
	api.POST("/update/:id", rc.Handler.Register)
	api.DELETE("/delete/:id", rc.Handler.Register)
}

package routes

func (rc *RouteConfig) PublisherRoute() {
	api := rc.App.Group("/publisher")
	api.GET("", rc.PublisherHandler.GetPublisher)
	api.GET("/:id", rc.PublisherHandler.GetPublisherById)
	api.POST("/create", rc.PublisherHandler.CreatePublisher)
	api.POST("/update/:id", rc.PublisherHandler.UpdatePublisher)
	api.DELETE("/delete/:id", rc.PublisherHandler.DeletePublisher)
}

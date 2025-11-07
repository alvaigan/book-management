package routes

func (rc *RouteConfig) AuthRoute() {
	api := rc.App.Group("/auth")
	api.POST("/login", rc.Handler.Login)
	api.POST("/register", rc.Handler.Register)
}

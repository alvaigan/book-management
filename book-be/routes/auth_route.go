package routes

func (rc *RouteConfig) AuthRoute() {
	api := rc.App.Group("/auth")
	api.POST("/login", rc.AuthHandler.Login)
	api.POST("/register", rc.AuthHandler.Register)
}

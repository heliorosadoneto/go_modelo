package router

import (
	"git/handlers"
	"git/middleware"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func Inicializar() {
	router := gin.Default()

	store := cookie.NewStore([]byte("segredo"))
	router.Use(sessions.Sessions("sessao", store))

	router.POST("/login", handlers.LoginHandler)
	router.POST("/register", handlers.RegisterHandler)
	router.POST("/logout", handlers.LogoutHandler)

	auth := router.Group("/dashboard")
	auth.Use(middleware.AuthRequired())
	auth.GET("/", handlers.DashboardHandler)

	router.Run(":8080") // listen and serve on
}

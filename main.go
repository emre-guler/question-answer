package main

import (
	"os"

	"github.com/emre-guler/question-answer/globals"
	middleweare "github.com/emre-guler/question-answer/middleware"
	routes "github.com/emre-guler/question-answer/routes"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"

	"github.com/gin-gonic/gin"
)

var port string = os.Getenv("PORT")

func main() {
	router := gin.Default()

	router.Static("/statics", "./statics")
	router.LoadHTMLGlob("./templates/*.gohtml")

	router.Use(sessions.Sessions("sessions", cookie.NewStore(globals.Secret)))

	public := router.Group("/")
	routes.PublicRoutes(public)

	private := router.Group("/")
	private.Use(middleweare.AuthRequired)
	routes.PrivateRoutes(private)

	router.Run(("localhost:" + port))
}

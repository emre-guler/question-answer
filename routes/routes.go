package routes

import (
	"github.com/emre-guler/question-answer/controllers"
	"github.com/gin-gonic/gin"
)

func PublicRoutes(g *gin.RouterGroup) {
	g.GET("/login", controllers.LoginGetHandler())
	g.GET("/callback", controllers.CallbackGetHandler())
}

func PrivateRoutes(g *gin.RouterGroup) {

}

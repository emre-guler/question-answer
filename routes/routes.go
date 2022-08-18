package routes

import (
	appHandlers "github.com/emre-guler/question-answer/controllers/appHandlers"
	loginHandlers "github.com/emre-guler/question-answer/controllers/loginHandlers"
	"github.com/gin-gonic/gin"
)

func PublicRoutes(g *gin.RouterGroup) {
	g.GET("/login", loginHandlers.LoginGetHandler())
	g.GET("/callback", loginHandlers.CallbackGetHandler())
}

func PrivateRoutes(g *gin.RouterGroup) {
	g.GET("/main", appHandlers.MainGetHandler())
	g.POST("/main", appHandlers.MainPostHandler())
}

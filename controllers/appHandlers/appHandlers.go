package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func MainGetHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "main.gohtml", nil)
	}
}

func MainPostHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {

	}
}

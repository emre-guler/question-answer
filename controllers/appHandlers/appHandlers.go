package controllers

import (
	"log"

	"github.com/gin-gonic/gin"
)

func AppGetHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		log.Println("Hello World!")
	}
}

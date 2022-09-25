package middleware

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func SetCustomDefaultMiddleware(router *gin.Engine) *gin.Engine {
	router.Use(RequestLogger)
	return router
}

func setMiddleware(router *gin.Engine, middleware gin.HandlerFunc) *gin.Engine {
	router.Use(middleware)
	return router
}

func RequestLogger(c *gin.Context) {
	fmt.Println("logger running")
}

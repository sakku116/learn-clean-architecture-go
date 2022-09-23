package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func SetMiddleware(router *gin.Engine) *gin.Engine {
	router.Use(RequestLoggerMiddleware)

	return router
}

func RequestLoggerMiddleware(c *gin.Context) {
	fmt.Println("logger running")
}

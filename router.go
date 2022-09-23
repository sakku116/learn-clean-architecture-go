package main

import "github.com/gin-gonic/gin"

func SetRouter(router *gin.Engine) *gin.Engine {
	router.GET("ping/", TestHandler)

	return router
}

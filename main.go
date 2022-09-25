package main

import (
	"learn-clean-arc/middleware"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	middleware.SetCustomDefaultMiddleware(router)
	SetRouter(router)

	router.Run()
}

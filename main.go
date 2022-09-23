package main

import "github.com/gin-gonic/gin"

func main() {
	router := gin.Default()

	SetMiddleware(router)
	SetRouter(router)

	router.Run()
}

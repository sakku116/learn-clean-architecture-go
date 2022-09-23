package main

import "github.com/gin-gonic/gin"

func TestHandler(c *gin.Context) {
	c.IndentedJSON(200, gin.H{
		"message": "Pong",
	})
}

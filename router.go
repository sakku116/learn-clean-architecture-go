package main

import (
	"learn-clean-arc/handler"
	"learn-clean-arc/service"

	"github.com/gin-gonic/gin"
)

func SetRouter(router *gin.Engine) *gin.Engine {
	router.GET("ping/", func(c *gin.Context) {
		c.IndentedJSON(200, gin.H{
			"message": "Pong",
		})
	})

	product_api := router.Group("api/")
	{
		product_handler := handler.Product{
			Service: service.Product{},
		}
		product_api.GET("products/", product_handler.GetAll)
		product_api.GET("products/:id", product_handler.GetByID)
		product_api.POST("products/", product_handler.Create)
		product_api.PUT("products/:id", product_handler.Update)
		product_api.DELETE("products/:id", product_handler.Delete)
	}

	return router
}

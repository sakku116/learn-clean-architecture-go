package handler

import (
	"learn-clean-arc/dto"

	"github.com/gin-gonic/gin"
)

type Product struct {
}

func (h *Product) GetAll(c *gin.Context) {
}

func (h *Product) GetByID(c *gin.Context) {
	// id := c.Param("id")

}

func (h *Product) Create(c *gin.Context) {
	// get input
	var input dto.ProductBodyRequest
	c.Bind(&input)

}

func (h *Product) Update(c *gin.Context) {
	// id := c.Param("id")

}

func (h *Product) Delete(c *gin.Context) {
	// id := c.Param("id")

}

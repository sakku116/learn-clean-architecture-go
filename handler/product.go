package handler

import (
	"learn-clean-arc/dto"
	"learn-clean-arc/service"

	"github.com/gin-gonic/gin"
)

type Product struct {
	Service service.Product
}

func (h *Product) GetAll(c *gin.Context) {
	// products =
}

func (h *Product) GetByID(c *gin.Context) {
	_ = c.Param("id")

}

func (h *Product) Create(c *gin.Context) {
	// get input
	var input dto.ProductBodyRequest
	c.Bind(&input)

}

func (h *Product) Update(c *gin.Context) {
	_ = c.Param("id")

}

func (h *Product) Delete(c *gin.Context) {
	_ = c.Param("id")

}

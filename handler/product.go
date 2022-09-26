package handler

import (
	"learn-clean-arc/dto"
	"learn-clean-arc/response"
	"learn-clean-arc/service"

	"github.com/gin-gonic/gin"
)

type Product struct {
	Service service.Product
}

// ------------------- method ----------------------

func (h *Product) GetAll(c *gin.Context) {
	// products =
	response.SendMessage(c, "product handler => getAll")
}

func (h *Product) GetByID(c *gin.Context) {
	_ = c.Param("id")

	response.SendMessage(c, "product handler => getByID")
}

func (h *Product) Create(c *gin.Context) {
	// get input
	var input dto.ProductBodyRequest
	c.Bind(&input)

	response.SendMessage(c, "product handler => Create")
}

func (h *Product) Update(c *gin.Context) {
	_ = c.Param("id")

	response.SendMessage(c, "product handler => Update")
}

func (h *Product) Delete(c *gin.Context) {
	_ = c.Param("id")

	response.SendMessage(c, "product handler => Delete")
}

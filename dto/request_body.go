package dto

type ProductBodyRequest struct {
	Name  string `json:"name" form:"name"`
	Price int    `json:"price" form:"price"`
}

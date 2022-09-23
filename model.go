package main

type Product struct {
	Name      string
	Price     int
	CreatedAt string
	Sold      bool
}

func (model *Product) TableName() string {
	return "product"
}

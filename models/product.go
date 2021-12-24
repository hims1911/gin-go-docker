package models

type Product struct {
	ID    uint   `form:"id" json:"id"`
	Code  string `form:"code" json:"code"`
	Price uint   `form:"price" json:"price"`
}

type CreateProduct struct {
	Code  string `form:"code" json:"code"`
	Price uint   `form:"price" json:"price"`
}

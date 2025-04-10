package models

type Product struct {
	Id         uint64   `json:"id" db:"id" `
	Name       string   `json:"name" db:"name" binding:"required"`
	Quantity   uint32   `json:"quantity" db:"quantity" binding:"required"`
	Price      uint32   `json:"price" db:"price" binding:"required"`
	CategoryID uint64   `db:"category_id" json:"category_id" binding:"required"`
	Category   Category `json:"category" db:"category"`
}

type Category struct {
	Id   uint64 `db:"id" json:"id"`
	Name string `db:"name" json:"name"`
}

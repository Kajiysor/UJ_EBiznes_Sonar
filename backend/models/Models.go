package models

type Product struct {
	ID    uint   `gorm:"primary_key"`
	Name  string
	CategoryID uint
	Description string
	Price float64 
}

type Category struct {
	ID    uint   `gorm:"primary_key"`
	Name  string
}

type CartItem struct {
	ID    uint   `gorm:"primary_key"`
	ProductID uint `json:"product_id"`
	Quantity uint
}

type Confirmation struct {
	Name string
	Address string
	City string
	Zip string
	Delivery string
	CardNo string
	CardExp string
	CVV string
}
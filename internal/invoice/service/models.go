package service

var mapProductID = map[string]ProductDetail{
	"product1": {
		Name:   "Clothing Item 1",
		Amount: 500000,
		Notes:  "This is a beautiful piece of clothing.",
	},
	"product2": {
		Name:   "Clothing Item 2",
		Amount: 300000,
		Notes:  "Another stylish clothing choice.",
	},
}

type ProductDetail struct {
	Name   string
	Amount float32
	Notes  string
}

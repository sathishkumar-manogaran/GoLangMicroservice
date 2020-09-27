package data

import (
	"encoding/json"
	"io"
	"time"
)

// Product pojo / domain class for coffee shop
type Product struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float32 `json:"price"`
	SKU         string  `json:"sku"`
	CreatedOn   string  `json:"-"`
	UpdatedOn   string  `json:"-"`
	DeletedOn   string  `json:"-"`
}

// Creating 'type' for this Product so we can create func / method to expose
type Products []*Product

// Method created to expose and this is encapsulated object
func (p *Products) ToJSON(w io.Writer) error {
	encodedValue := json.NewEncoder(w)
	return encodedValue.Encode(p)
}

func GetProducts() Products {
	return products
}

var products = []*Product{
	&Product{
		ID:          1,
		Name:        "Latte",
		Description: "Forthy Milk Coffee",
		Price:       3.45,
		SKU:         "abc123",
		CreatedOn:   time.Now().UTC().String(),
		UpdatedOn:   time.Now().UTC().String(),
	},
	&Product{
		ID:          2,
		Name:        "Espresso",
		Description: "Short and Strong Coffee Without Milk",
		Price:       5.68,
		SKU:         "xyz123",
		CreatedOn:   time.Now().UTC().String(),
		UpdatedOn:   time.Now().UTC().String(),
	},
}

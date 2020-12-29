package data

import (
	"encoding/json"
	"io"
	"time"
)

type Product struct {
	ID          int     `json: "id"` // we tell json to lower case the id
	Name        string  `json: "name"`
	Description string  `json: "description"`
	Price       float32 `json: "price"`
	SKU         string  `json: "sku"`
	CreatedOn   string  `json: "-"` //we tell json not to print these lines with `json: "-"`
	UpdateOn    string  `json: "-"`
	DeleteOn    string  `json: "-"`
}
type Products []Product

func (p *Products) ToJ(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(p)
}
func GetProducts() Products {
	return ProductList
}

var ProductList = []*Product{
	&Product{
		ID:          1,
		Name:        "Cappucgino",
		Description: "Caco et du lait",
		Price:       2.45,
		SKU:         "abc323",
		CreatedOn:   time.Now().UTC().String(),
		UpdateOn:    time.Now().UTC().String(),
		DeleteOn:    time.Now().UTC().String(),
	},
	&Product{
		ID:          2,
		Name:        "Espresso",
		Description: "Café long avec du lait",
		Price:       2.50,
		SKU:         "fjd34",
		CreatedOn:   time.Now().UTC().String(),
		UpdateOn:    time.Now().UTC().String(),
		DeleteOn:    time.Now().UTC().String(),
	},
}

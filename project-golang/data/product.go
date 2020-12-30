package data

import (
	"encoding/json"
	"fmt"
	"io"
	"time"
)

type Product struct {
	ID          int     `json:"id"` // we tell json to lower case the id
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float32 `json:"price"`
	SKU         string  `json:"sku"`
	CreatedOn   string  `json:"-"` //we tell json not to print these lines with `json: "-"`
	UpdateOn    string  `json:"-"`
	DeleteOn    string  `json:"-"`
}
type Products []*Product

//the function  marshall  the struct
func (p *Products) ToJ(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(p)
}
func (p *Product) FromJ(r io.Reader) error {
	e := json.NewDecoder(r)
	return e.Decode(p)
}
func GetProducts() Products {
	return productList
}
func AddProducts(p *Product) {
	p.ID = GetNextId()
	productList = append(productList, p)
}

func GetNextId() int {
	lp := productList[len(productList)-1]
	return lp.ID + 1
}
func UpdateProducts(id int, p *Product) error {
	_, pos, err := findProduct(id)
	if err != nil {
		return err
	}
	p.ID = id
	productList[pos] = p
	return nil
}

var ErrorProductNotfound = fmt.Errorf("Product not found")

func findProduct(id int) (*Product, int, error) {
	for i, p := range productList {
		if p.ID == id {
			return p, i, nil
		}
	}
	return nil, -1, ErrorProductNotfound
}

var productList = []*Product{
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

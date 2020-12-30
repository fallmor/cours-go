package handlers

import (
	"log"
	"net/http"

	"github.com/fallmor/cours-go/project-golang/data"
)

type Products struct {
	l *log.Logger
}

func NewProducts(l *log.Logger) *Products {
	return &Products{l}
}

func (p *Products) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		p.GetProducts(rw, r)
		return
	}
	if r.Method == http.MethodPost {
		p.addProduct(rw, r)
	}
	rw.WriteHeader(http.StatusMethodNotAllowed)

	// fetch the products from the datastore

	// serialize the list to JSON

}
func (p *Products) GetProducts(rw http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle GET Products")

	// fetch the products from the datastore
	lp := data.GetProducts()

	// serialize the list to JSON
	err := lp.ToJ(rw)
	if err != nil {
		http.Error(rw, "Unable to marshal json", http.StatusInternalServerError)
	}
}
func (p *Products) addProduct(rw http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle Post Products")
	prod := &data.Product{}
	err := prod.FromJ(r.Body)
	if err != nil {
		http.Error(rw, "Unable to use Post Method", http.StatusBadRequest)
	}
	p.l.Printf("Prod: %#v", prod)
	data.AddProducts(prod)
}

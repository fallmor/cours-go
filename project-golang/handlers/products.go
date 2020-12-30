package handlers

import (
	"log"
	"net/http"
	"regexp"
	"strconv"

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
	if r.Method == http.MethodPut {
		reg := regexp.MustCompile(`/([0-9]+)`)
		g := reg.FindAllStringSubmatch(r.URL.Path, -1)
		if len(g) != 1 {
			p.l.Println("Invalid URI request more than one ID")
			http.Error(rw, "Invalid URI", http.StatusBadRequest)
		}
		if len(g[0]) != 2 {
			p.l.Println("Invalid URI request more than one capture group")
			http.Error(rw, "Invalid URI", http.StatusBadRequest)
		}
		idString := g[0][1]
		id, err := strconv.Atoi(idString)
		if err != nil {
			p.l.Println("Invalid URI unable to convert to numer", idString)
			http.Error(rw, "Invalid URI", http.StatusBadRequest)
			return
		}
		p.l.Println("Update request pass ", id)
		p.UpdateProducts(id, rw, r)
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
func (p *Products) UpdateProducts(id int, rw http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle Post Products")
	prod := &data.Product{}
	err := prod.FromJ(r.Body)
	if err != nil {
		http.Error(rw, "Unable to use Post Method", http.StatusBadRequest)
	}
	p.l.Printf("Prod: %#v", prod)
	err = data.UpdateProducts(id, prod)
	if err == data.ErrorProductNotfound {
		http.Error(rw, "Product Not found", http.StatusNotFound)
		return
	}
	if err != nil {
		http.Error(rw, "Product not found", http.StatusInternalServerError)
		return
	}
}

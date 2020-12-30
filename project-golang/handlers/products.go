package handlers

import (
	"context"
	"log"
	"net/http"
	"strconv"

	"github.com/fallmor/cours-go/project-golang/data"
	"github.com/gorilla/mux"
)

type Products struct {
	l *log.Logger
}

func NewProducts(l *log.Logger) *Products {
	return &Products{l}
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
func (p *Products) AddProduct(rw http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle Post Products")
	prod := r.Context().Value(KeyProdruct{}).(data.Product)
	data.AddProducts(&prod)
}
func (p *Products) UpdateProducts(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(rw, "Unable to convert Id", http.StatusBadRequest)
	}

	p.l.Println("Handle Post Products")
	prod := r.Context().Value(KeyProdruct{}).(data.Product)
	err = data.UpdateProducts(id, &prod)
	if err == data.ErrorProductNotfound {
		http.Error(rw, "Product Not found", http.StatusNotFound)
		return
	}
	if err != nil {
		http.Error(rw, "Product not found", http.StatusInternalServerError)
		return
	}
}

type KeyProdruct struct{}

func (p Products) MiddlewareProdutcsValidation(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		prod := data.Product{}
		err := prod.FromJ(r.Body)
		if err != nil {
			http.Error(rw, "Unable to use Post Method", http.StatusBadRequest)
			return
		}
		ctx := context.WithValue(r.Context(), KeyProdruct{}, prod)
		req := r.WithContext(ctx)
		next.ServeHTTP(rw, req)
	})
}

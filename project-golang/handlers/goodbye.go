package handlers

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type GoodBye struct {
	l *log.Logger
}

func NewGoodBye(l *log.Logger) *GoodBye {
	return &GoodBye{l}
}

func (G *GoodBye) ServeHTTP(wr http.ResponseWriter, r *http.Request) {
	G.l.Println("Handle Goddbye Request")
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		G.l.Println("Error", err)
		http.Error(wr, "An error occurs", http.StatusBadGateway)
	}

	fmt.Fprintf(wr, "GoodBye %s", b)
}

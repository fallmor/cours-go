package handlers

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Hello struct {
}

func (h *Hello) httpServe(wr http.ResponseWriter, r *http.Request) {
	log.Println("Hello from Mor")
	d, err := ioutil.ReadAll(r.Body)
	//log.Printf("data %s\n", d) //affiche les requetes passées avec curl
	fmt.Fprintf(wr, "Hello %s", d) //affiche les reponses avec le wr qui represente http.ResponseWriter
	if err != nil {
		//wr.WriteHeader(http.StatusBadGateway) //Go permet de gerer les codes http telque 404 de manière très simple avec le package http
		//wr.Write([]byte("Ooops"))             // le message qui sera affiché avec le status code
		http.Error(wr, "Ooops", http.StatusBadGateway)
		return
	}
}

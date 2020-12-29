package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/fallmor/cours-go/project-golang/handlers"
)

func main() {

	/*http.HandleFunc("/", func(wr http.ResponseWriter, r *http.Request) {
	 		//cette fonction va recevoir les requestes http (HandleFunc qui prend deux arguments le chemin, et une fonction )
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
		})
		http.HandleFunc("/Bye", func(http.ResponseWriter, *http.Request) {
			log.Println("ByeBye from Mor")
		})*/ //le code handler est géré maintenant par le fichier hello.go qui est dans le dossier handler

	l := log.New(os.Stdout, "product-api", log.LstdFlags)
	//create handler
	ph := handlers.NewProducts(l)
	gg := handlers.NewGoodBye(l)
	sm := http.NewServeMux()
	sm.Handle("/", ph)
	sm.Handle("/goodbye", gg)
	s := http.Server{
		Addr:         ":9090",
		Handler:      sm,
		IdleTimeout:  120 * time.Second,
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 1 * time.Second,
	}
	go func() {
		err := s.ListenAndServe() //accepte toutes les connexions sur le port 9090. En quelque sorte on crée un serveur http
		if err != nil {
			l.Fatal(err)
		}
	}()
	sigchan := make(chan os.Signal)
	signal.Notify(sigchan, os.Interrupt)
	signal.Notify(sigchan, os.Kill)
	sig := <-sigchan
	l.Println("Terminating the server", sig)
	tc, _ := context.WithTimeout(context.Background(), 30*time.Second)
	s.Shutdown(tc)
}

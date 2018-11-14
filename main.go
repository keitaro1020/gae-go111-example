package main

import (
	"fmt"
	"github.com/go-chi/chi"
	"github.com/keitaro1020/gae-go111-example/handler"
	"log"
	"net/http"
	"os"
)

func main() {

	r := chi.NewRouter()

	r.Get("/", handler.HelloHandler)
	r.Route("/api", func(r chi.Router){

	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Printf("Defaulting to port %s", port)
	}

	log.Printf("Listening on port %s", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), r))
}

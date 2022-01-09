package main

import (
	"log"
	"net/http"

	"github.com/apaarshrm39/Go-Resume/pkg/handlers"
)

func main() {

	mux := http.NewServeMux()

	mux.HandleFunc("/", handlers.Home)

	log.Fatal(http.ListenAndServe(":9000", mux))
}

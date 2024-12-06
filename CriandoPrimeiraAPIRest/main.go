package main

import (
	"PrimeiraAPIRest/routes"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	// Configura as rotas de clientes
	routes.ClienteRoutes(router)

	log.Println("Servidor iniciado na porta 8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}

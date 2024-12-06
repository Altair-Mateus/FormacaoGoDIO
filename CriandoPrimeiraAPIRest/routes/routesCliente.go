package routes

import (
	"PrimeiraAPIRest/handlers"
	"net/http"

	"github.com/gorilla/mux"
)

// Configura as rotas relacionadas a clientes
func ClienteRoutes(router *mux.Router) {
	router.HandleFunc("/clientes", handlers.GetClientes).Methods(http.MethodGet)
	router.HandleFunc("/clientes/{id}", handlers.GetCliente).Methods(http.MethodGet)
	router.HandleFunc("/clientes", handlers.CreateCliente).Methods(http.MethodPost)
	router.HandleFunc("/clientes/{id}", handlers.DeleteCliente).Methods(http.MethodDelete)
}

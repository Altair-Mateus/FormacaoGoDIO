package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	models "PrimeiraAPIRest/models"

	"github.com/gorilla/mux"
)

var cliente []models.Cliente

// Retorna todos os clientes
func GetClientes(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	encoder := json.NewEncoder(w)
	encoder.SetIndent("", "  ") // Adiciona indentação
	encoder.Encode(cliente)
}

// Retorna um cliente específico por ID
func GetCliente(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	// Converte a id para inteiro
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "ID inválida", http.StatusBadRequest)
		return
	}

	// localiza o cliente
	for _, item := range cliente {
		if item.Id == id {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusCreated)
			encoder := json.NewEncoder(w)
			encoder.SetIndent("", "  ") // Adiciona indentação
			encoder.Encode(item)
			return
		}
	}

	http.Error(w, "Cliente não encontrado", http.StatusNotFound)
}

// Adiociona um novo cliente
func CreateCliente(w http.ResponseWriter, r *http.Request) {
	var novoCliente models.Cliente

	// Decodificando o body da requisição para a struct
	err := json.NewDecoder(r.Body).Decode(&novoCliente)
	if err != nil {
		http.Error(w, "Dados inválidos para adicionar novo cliente", http.StatusBadRequest)
		return
	}

	novoCliente.Id = len(cliente) + 1
	cliente = append(cliente, novoCliente)

	// Retornando o cliente e o status de criado
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	encoder := json.NewEncoder(w)
	encoder.SetIndent("", "  ") // Adiciona indentação
	encoder.Encode(novoCliente)
}

// Remove o cliente de acordo com a ID
func DeleteCliente(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	// Converte a id para inteiro
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "Id inválido", http.StatusBadRequest)
		return
	}

	// Localiza o cliente
	for index, cli := range cliente {
		if cli.Id == id {
			cliente = append(cliente[:index], cliente[index+1:]...)
			w.WriteHeader(http.StatusNoContent)
			return
		}
	}

	http.Error(w, "Cliente não encontrado", http.StatusNotFound)
}

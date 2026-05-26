package routes

import (
	"myapi/internal/handlers"

	"github.com/gorilla/mux"
)

func SetupCategoriaRoutes(r *mux.Router) {
	r.HandleFunc("/categorias", handlers.ListCategorias).Methods("GET")
	r.HandleFunc("/categorias/{id}", handlers.GetCategoria).Methods("GET")
	r.HandleFunc("/categorias", handlers.CreateCategoria).Methods("POST")
	r.HandleFunc("/categorias", handlers.UpdateCategoria).Methods("PUT")
	r.HandleFunc("/categorias/{id}", handlers.DeleteCategoria).Methods("DELETE")
}

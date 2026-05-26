package main

import (
	"log"
	"myapi/internal/config"
	"myapi/internal/routes"
	"net/http"
)

func main() {
	config.ConectDatabase()   // config: banco de dados
	r := routes.SetupRoutes() // config: rotas

	log.Println("Servidor rodando na porta 8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}

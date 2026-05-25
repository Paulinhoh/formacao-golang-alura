package main

import (
	"api-rest/database"
	"api-rest/routes"
	"fmt"
)

func main() {
	database.ConectarComBancoDeDados()

	fmt.Println("Iniciando o servidor REST com Go")
	routes.HandleResquest()
}

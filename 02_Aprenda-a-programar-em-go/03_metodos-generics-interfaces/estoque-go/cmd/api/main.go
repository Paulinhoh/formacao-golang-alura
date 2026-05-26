package main

import (
	"estoque-go/internal/models"
	"estoque-go/internal/services"
	"fmt"
)

func main() {
	fmt.Println("Sistema de Estoque")

	estoque := services.NewEstoque()
	items := []models.Item{
		{ID: 1, Name: "Fone", Quantity: 5, Price: 100},
		{ID: 2, Name: "Camiseta", Quantity: 1, Price: 55.99},
		{ID: 3, Name: "Mouse", Quantity: 2, Price: 12.99},
	}

	for _, item := range items {
		if err := estoque.AddItem(item, "Paulinho"); err != nil {
			fmt.Println(err)
		}
	}

	for _, item := range estoque.ListItems() {
		fmt.Printf("ID: %d | Name: %s | Quabtidade: %d | Preço: %.2f\n", item.ID, item.Name, item.Quantity, item.Price)
	}
	fmt.Println("Valor total do estoque:", estoque.CalculateTotalCost())

	if err := estoque.DeleteItem(3, 2, "Paulinho"); err != nil {
		fmt.Println(err)
	}

	logs := estoque.ViewAuditLog()
	for _, log := range logs {
		fmt.Printf("[%s] Ação: %s - Usuario: %s - Item ID: %d - Quantidade: %d - Motivo: %s\n", log.Timestamp.Format("02/01/2006-15:04:05"), log.Action, log.User, log.ItemID, log.Quantity, log.Reason)
	}

	itemBySearch, err := services.FindBy(items, func(item models.Item) bool {
		return item.Price > 40
	})
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Item encontrado:", itemBySearch)

	alura := services.Fornecedor{
		CNPJ:    "123456",
		Contato: "11970707070",
		Cidade:  "São Paulo",
	}
	fmt.Print(alura.GetInfo())
	if alura.VerificarDisponibilidade(10, 15) {
		fmt.Println("Possui disponibilidade")
	} else {
		fmt.Println("Não possui disponibilidade")
	}
}

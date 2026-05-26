package services

import (
	"estoque-go/internal/models"
	"fmt"
	"strconv"
	"time"
)

type estoque struct {
	items map[string]models.Item
	logs  []models.Log
}

func NewEstoque() *estoque {
	return &estoque{
		items: make(map[string]models.Item),
		logs:  []models.Log{},
	}
}

func (e *estoque) AddItem(item models.Item, user string) error {
	if item.Quantity <= 0 {
		return fmt.Errorf("erro ao adicionar o item: [ID:%d] possui uma quantidade inválida (zero ou negativa)", item.ID)
	}

	existingItem, exists := e.items[strconv.Itoa(item.ID)]
	if exists {
		item.Quantity += existingItem.Quantity
	}

	e.items[strconv.Itoa(item.ID)] = item
	e.logs = append(e.logs, models.Log{
		Timestamp: time.Now(),
		Action:    "Entrada de estoque",
		User:      user,
		ItemID:    item.ID,
		Quantity:  item.Quantity,
		Reason:    "Adicionando novos itens no estoque",
	})
	return nil
}

func (e *estoque) ListItems() []models.Item {
	var itemList []models.Item
	for _, item := range e.items {
		itemList = append(itemList, item)
	}

	return itemList
}

func (e *estoque) ViewAuditLog() []models.Log {
	return e.logs
}

func (e *estoque) CalculateTotalCost() float64 {
	var totCost float64

	for _, item := range e.items {
		totCost += float64(item.Quantity) * item.Price
	}

	return totCost
}

func (e *estoque) DeleteItem(itemID int, quantity int, user string) error {
	// Valide a existência do item no estoque
	existingItem, exists := e.items[strconv.Itoa(itemID)]
	if !exists {
		return fmt.Errorf("erro ao remover item: [ID:%d] não existe no estoque", itemID)
	}

	// Verifique se a quantidade a ser removida é válida
	if quantity <= 0 {
		return fmt.Errorf("erro ao remover item: quantidade inválida (zero ou negativa) para [ID:%d]", itemID)
	}

	// Verifique se a quantidade disponível no estoque é suficiente
	if existingItem.Quantity < quantity {
		return fmt.Errorf("erro ao remover item: estoque insuficiente para [ID:%d]. Disponível: %d, Solicitado: %d", itemID, existingItem.Quantity, quantity)
	}

	// Atualize o estoque removendo a quantidade informada
	existingItem.Quantity -= quantity
	if existingItem.Quantity == 0 {
		// Remova o item completamente se a quantidade for zero
		delete(e.items, strconv.Itoa(itemID))
	} else {
		// Atualize a quantidade do item no estoque
		e.items[strconv.Itoa(itemID)] = existingItem
	}

	// Registre a operação em um log
	e.logs = append(e.logs, models.Log{
		Timestamp: time.Now(),
		Action:    "Saída de estoque",
		User:      user,
		ItemID:    itemID,
		Quantity:  quantity,
		Reason:    "Removendo itens do estoque",
	})

	return nil
}

func FindBy[T any](data []T, comparator func(T) bool) ([]T, error) {
	var result []T
	for _, v := range data {
		if comparator(v) {
			result = append(result, v)
		}
	}
	if len(result) == 0 {
		return nil, fmt.Errorf("nenhum item foi encontrado")
	}
	return result, nil
}

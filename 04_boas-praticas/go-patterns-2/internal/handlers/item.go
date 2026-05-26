package handlers

import (
	"encoding/json"
	"myapi/internal/repositories"
	"myapi/internal/services"
	"myapi/internal/utils"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// ListItens - Lista todos os itens
func ListItems(w http.ResponseWriter, r *http.Request) {
	repository := repositories.NewItemRepository()
	items, err := repository.ListAll()
	if err != nil {
		utils.RespondWithError(w, "Erro ao listar os itens", http.StatusNotFound)
		return
	}
	if err := json.NewEncoder(w).Encode(items); err != nil {
		utils.RespondWithError(w, "Erro ao codificar os itens", http.StatusInternalServerError)
	}
}

// GetItem - Busca um item por ID (via rota: /item/{id})
func GetItem(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr := vars["id"]

	if idStr == "" {
		utils.RespondWithError(w, "ID não fornecido", http.StatusBadRequest)
		return
	}
	id, err := strconv.Atoi(idStr)
	if err != nil {
		utils.RespondWithError(w, "ID inválido", http.StatusBadRequest)
		return
	}

	repository := repositories.NewItemRepository()
	item, err := repository.GetByID(id)
	if err != nil {
		utils.RespondWithError(w, "Item não encontrado", http.StatusNotFound)
		return
	}

	if err := json.NewEncoder(w).Encode(item); err != nil {
		utils.RespondWithError(w, "Erro ao codificar os itens", http.StatusInternalServerError)
	}
}

// GetItemByCode - Busca um item pelo campo "codigo" (via rota: /item/codigo/{codigo})
func GetItemByCode(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	code := vars["codigo"]

	if code == "" {
		utils.RespondWithError(w, "Código não fornecido", http.StatusBadRequest)
		return
	}

	repository := repositories.NewItemRepository()
	item, err := repository.GetByCode(code)
	if err != nil {
		utils.RespondWithError(w, "Item não encontrado", http.StatusNotFound)
		return
	}

	if err := json.NewEncoder(w).Encode(item); err != nil {
		utils.RespondWithError(w, "Erro ao codificar os itens", http.StatusInternalServerError)
	}
}

// CreateItem - Cria um novo item (envie JSON via POST)
func CreateItem(w http.ResponseWriter, r *http.Request) {
	item, err := services.DecodeAndValidaItem(r)
	if err != nil {
		utils.RespondWithError(w, err.Error(), http.StatusBadRequest)
		return
	}

	repository := repositories.NewItemRepository()
	createdItem, err := repository.Create(item)
	if err != nil {
		utils.RespondWithError(w, "Erro ao criar o item", http.StatusInternalServerError)
		return
	}

	if err := json.NewEncoder(w).Encode(createdItem); err != nil {
		utils.RespondWithError(w, "Erro ao codificar os itens", http.StatusInternalServerError)
	}
}

// UpdateItem - Atualiza um item existente (envie JSON via PUT, com o campo id preenchido)
func UpdateItem(w http.ResponseWriter, r *http.Request) {
	item, err := services.DecodeAndValidaItem(r)
	if err != nil {
		utils.RespondWithError(w, err.Error(), http.StatusBadRequest)
		return
	}

	repository := repositories.NewItemRepository()
	if err := repository.Update(item); err != nil {
		utils.RespondWithError(w, "Erro ao atualizar o item", http.StatusInternalServerError)
		return
	}

	if err := json.NewEncoder(w).Encode(item); err != nil {
		utils.RespondWithError(w, "Erro ao codificar os itens", http.StatusInternalServerError)
	}
}

// DeleteItem - Deleta um item por ID (via rota: /item/{id})
func DeleteItem(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr := vars["id"]

	if idStr == "" {
		utils.RespondWithError(w, "ID não fornecido", http.StatusBadRequest)
		return
	}
	id, err := strconv.Atoi(idStr)
	if err != nil {
		utils.RespondWithError(w, "ID inválido", http.StatusBadRequest)
		return
	}

	repository := repositories.NewItemRepository()
	if err := repository.Delete(id); err != nil {
		utils.RespondWithError(w, "Erro ao deletar o item", http.StatusInternalServerError)
		return
	}
	if _, err := w.Write([]byte("Item deletado com sucesso")); err != nil {
		utils.RespondWithError(w, err.Error(), http.StatusInternalServerError)
	}
}

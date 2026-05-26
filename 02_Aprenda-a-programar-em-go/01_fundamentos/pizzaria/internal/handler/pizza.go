package handler

import (
	"net/http"
	"pizzaria/internal/data"
	"pizzaria/internal/models"
	"pizzaria/internal/service"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetPizzas(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"pizzas": data.Pizzas,
	})
}

func PostPizzas(ctx *gin.Context) {
	var newPizza models.Pizza
	if err := ctx.ShouldBindJSON(&newPizza); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"erro": err.Error()})
		return
	}

	if err := service.ValidatePizzaPrice(&newPizza); err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"erro": err.Error()})
		return
	}

	newPizza.ID = len(data.Pizzas) + 1
	data.Pizzas = append(data.Pizzas, newPizza)
	data.SavePizza()
	ctx.JSON(http.StatusCreated, newPizza)
}

func GetPizzasByID(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"erro": err.Error()})
		return
	}

	for _, p := range data.Pizzas {
		if p.ID == id {
			ctx.JSON(http.StatusOK, p)
			return
		}
	}
	ctx.JSON(http.StatusNotFound, gin.H{"message": "pizza not found"})
}

func DeletePizzaByID(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"erro": err.Error()})
		return
	}

	for i, p := range data.Pizzas {
		if p.ID == id {
			data.Pizzas = append(data.Pizzas[:i], data.Pizzas[i+1:]...)
			data.SavePizza()
			ctx.JSON(http.StatusOK, gin.H{"message": "pizza deletada"})
			return
		}
	}
	ctx.JSON(http.StatusNotFound, gin.H{"message": "pizza not found"})
}

func UpdatePizzaByID(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"erro": err.Error()})
		return
	}

	var updatedPizza models.Pizza
	if err := ctx.ShouldBindJSON(&updatedPizza); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"erro": err.Error()})
		return
	}

	if err := service.ValidatePizzaPrice(&updatedPizza); err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"erro": err.Error()})
		return
	}

	for i, p := range data.Pizzas {
		if p.ID == id {
			data.Pizzas[i] = updatedPizza
			data.Pizzas[i].ID = id
			data.SavePizza()
			ctx.JSON(http.StatusOK, data.Pizzas[i])
			return
		}
	}
	ctx.JSON(http.StatusNotFound, gin.H{"message": "pizza not found"})
}

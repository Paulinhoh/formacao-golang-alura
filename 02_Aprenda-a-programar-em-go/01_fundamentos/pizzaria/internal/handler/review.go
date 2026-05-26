package handler

import (
	"net/http"
	"pizzaria/internal/data"
	"pizzaria/internal/models"
	"pizzaria/internal/service"
	"strconv"

	"github.com/gin-gonic/gin"
)

func PostReview(ctx *gin.Context) {
	pizzaID, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"erro": err.Error()})
		return
	}

	var newReview models.Review
	if err := ctx.ShouldBindJSON(&newReview); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"erro": err.Error()})
		return
	}

	if err := service.VaidateReviewRating(&newReview); err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"erro": err.Error()})
		return
	}

	for i, p := range data.Pizzas {
		if p.ID == pizzaID {
			p.Review = append(p.Review, newReview)
			data.Pizzas[i] = p
			data.SavePizza()
			ctx.JSON(http.StatusCreated, p)
			return
		}
	}

	ctx.JSON(http.StatusNotFound, gin.H{"message": "pizza not found"})
}

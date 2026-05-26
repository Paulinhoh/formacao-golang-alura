package validators

import (
	"errors"
	"myapi/internal/models"
)

func ValidateItem(item *models.Item) error {
	if item.Preco <= 0 {
		return errors.New("o preço não pode ser negativo")
	}

	if item.Quantidade < 0 {
		return errors.New("quantidade deve ser maior que zero")
	}

	if len(item.Codigo) != 6 {
		return errors.New("o código precisa ter 6 caracteres")
	}

	return nil
}

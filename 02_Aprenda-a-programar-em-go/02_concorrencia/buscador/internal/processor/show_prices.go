package processor

import (
	"buscador/internal/models"
	"fmt"
)

func ShowPriceAVG(priceChannel <-chan models.PriceDetail, done chan<- bool) {
	var totPrice float64
	countPrice := 0.0
	for price := range priceChannel {
		totPrice += price.Value
		countPrice++
		avgPrice := totPrice / countPrice

		fmt.Printf("[%s] | Store: %s | R$ %.2f | Preço médio até agr: R$ %.2f\n", price.TimeStamp.Format("02-Jan-2006 15:04:05"), price.StoreName, price.Value, avgPrice)
	}

	done <- true
}

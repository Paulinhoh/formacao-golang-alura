package main

import (
	"buscador/internal/fetcher"
	"buscador/internal/models"
	"buscador/internal/processor"
	"fmt"
	"time"
)

func main() {
	start := time.Now()

	priceChannel := make(chan models.PriceDetail, 4) // 4 -> buffer
	done := make(chan bool)

	go fetcher.FetchPrices(priceChannel)
	go processor.ShowPriceAVG(priceChannel, done)

	<-done // trava a execução do showPriceAVG até o valor true seja setado no channel done

	fmt.Printf("Tempo total: %v", time.Since(start))
}

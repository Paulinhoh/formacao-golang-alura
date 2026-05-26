package fetcher

import (
	"buscador/internal/models"
	"math/rand/v2"
	"sync"
	"time"
)

func FetchPrices(priceChannel chan<- models.PriceDetail) {
	var wg sync.WaitGroup
	wg.Add(4)

	go func() {
		defer wg.Done()
		priceChannel <- FetchePriceFromSite01()
	}()
	go func() {
		defer wg.Done()
		priceChannel <- FetchePriceFromSite02()
	}()
	go func() {
		defer wg.Done()
		priceChannel <- FetchePriceFromSite03()
	}()

	go func() {
		defer wg.Done()
		FetchAndSendMultiplePrices(priceChannel)
	}()

	wg.Wait()
	close(priceChannel)
}

// buscar preços de diferentes sites
func FetchePriceFromSite01() models.PriceDetail {
	time.Sleep(3 * time.Second)
	return models.PriceDetail{
		StoreName: "A",
		Value:     rand.Float64() * 100,
		TimeStamp: time.Now(),
	}
}

func FetchePriceFromSite02() models.PriceDetail {
	time.Sleep(1 * time.Second)
	return models.PriceDetail{
		StoreName: "B",
		Value:     rand.Float64() * 100,
		TimeStamp: time.Now(),
	}
}

func FetchePriceFromSite03() models.PriceDetail {
	time.Sleep(2 * time.Second)
	return models.PriceDetail{
		StoreName: "C",
		Value:     rand.Float64() * 100,
		TimeStamp: time.Now(),
	}
}

func FetchAndSendMultiplePrices(priceChannel chan<- models.PriceDetail) {
	time.Sleep(6 * time.Second)
	prices := []float64{
		rand.Float64() * 100,
		rand.Float64() * 100,
		rand.Float64() * 100,
	}

	for _, price := range prices {
		priceChannel <- models.PriceDetail{
			StoreName: "D",
			Value:     price,
			TimeStamp: time.Now(),
		}
	}
}

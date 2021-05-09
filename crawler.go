package main

import (
	"encoding/csv"
	"log"
	"os"

	"github.com/gocolly/colly"
)

func main() {
	fileName := "cryptocoinmarketcap.csv"
	file, err := os.Create(fileName)
	if err != nil {
		log.Fatalf("Cannot create file %q: %s\n", fileName, err)
		return
	}
	defer file.Close()
	writer := csv.NewWriter(file)
	defer writer.Flush()

	headers := []string{"Name", "Symbol", "Market Cap (USD)", "Price (USD)", "Circulating Supply (USD)", "Volume (24h)", "Change (1h)", "Change (24h)", "Change (7d)"}

	// setting header
	writer.Write(headers)

	cCrawler := colly.NewCollector()

	cCrawler.OnHTML("tbody tr", func(e *colly.HTMLElement) {
		data := []string{
			e.ChildText(".cmc-table__column-name"),
			e.ChildText(".cmc-table__cell--sort-by__symbol"),
			e.ChildText(".cmc-table__cell--sort-by__market-cap"),
			e.ChildText(".cmc-table__cell--sort-by__price"),
			e.ChildText(".cmc-table__cell--sort-by__circulating-supply"),
			e.ChildText(".cmc-table__cell--sort-by__volume-24-h"),
			e.ChildText(".cmc-table__cell--sort-by__percent-change-1-h"),
			e.ChildText(".cmc-table__cell--sort-by__percent-change-24-h"),
			e.ChildText(".cmc-table__cell--sort-by__percent-change-7-d"),
		}

		writer.Write(data)
	})

	cCrawler.Visit("https://coinmarketcap.com/all/views/all/")

	log.Printf("Scraping finished, check file %q for results\n", fileName)
}

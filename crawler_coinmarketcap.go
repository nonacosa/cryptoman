package main

import (
	"github.com/gocolly/colly"
)

type CrawlerCallback func([]string)

const CoinHost string = "https://coinmarketcap.com"
const AllApi string = "/all/views/all/"

const FieldSymbol string = ".cmc-table__cell--sort-by__symbol"
const FieldPrice string = ".cmc-table__cell--sort-by__price"
const FieldVolume24 string = ".cmc-table__cell--sort-by__volume-24-h"
const FieldMarketCap string = ".cmc-table__cell--sort-by__market-cap .sc-1ow4cwt-0"
const FieldChange1H string = ".cmc-table__cell--sort-by__percent-change-1-h"
const FieldChange24H string = ".cmc-table__cell--sort-by__percent-change-24-h"
const FieldChange7D string = ".cmc-table__cell--sort-by__percent-change-7-d"

func PullByLine(f CrawlerCallback) {
	// Instantiate default collector
	c := colly.NewCollector()

	count := 0
	c.OnHTML("tbody tr", func(e *colly.HTMLElement) {
		count++
		// only fetch top 20
		if count <= 20 {
			f([]string{
				e.ChildText(FieldSymbol),
				e.ChildText(FieldPrice),
				e.ChildText(FieldVolume24),
				e.ChildText(FieldMarketCap),
				AligningMinus(e.ChildText(FieldChange1H)),
				AligningMinus(e.ChildText(FieldChange24H)),
				AligningMinus(e.ChildText(FieldChange7D)),
			})
		}
	})
	c.Visit(CoinHost + AllApi)
}

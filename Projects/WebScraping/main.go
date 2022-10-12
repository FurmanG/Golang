package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/fatih/camelcase"
	"github.com/gocolly/colly"
	log "github.com/sirupsen/logrus"
)

var prossessnum uint = 0

func main() {
	// Stay same for all
	c := colly.NewCollector()
	c.OnError(func(_ *colly.Response, err error) {
		fmt.Println("Error:", err.Error())
	})

	c.OnResponse(func(r *colly.Response) {
		log.Info("The URL is response well: ", r.Request.URL)
	})

	// =================================listicle-slide-dek
	// Change from here:

	c.OnHTML("div.list", func(e *colly.HTMLElement) {
		log.Info("============Dogs===================")
		HtmlText := e.DOM
		dogname := HtmlText.Find("a").Text()
		dogname = strings.ReplaceAll(dogname, "See Details >", "")
		dogname = strings.TrimSpace(dogname)
		log.Info("-- dogname --> ", dogname)

		Lifespan := HtmlText.Find("span").Text()
		Lifespan = strings.ReplaceAll(Lifespan, "yearsPopularityHypoallergenicIntelligenceSee Details >Rank...HypoallergenicRank", "")
		log.Info("-- Lifespan --> ", Lifespan)

		Popularity := HtmlText.Find("div.pop").Text()
		Popularity = strings.ReplaceAll(Popularity, "Popularity", "")
		log.Info("-- Popularity --> ", Popularity)

		Hypoallergenic := HtmlText.Find("div.hyp").Text()
		Hypoallergenic = strings.ReplaceAll(Hypoallergenic, "Hypoallergenic", "")
		log.Info("-- Hypoallergenic --> ", Hypoallergenic)

		prices := HtmlText.Find("div.right-b").Text()
		prices = strings.ReplaceAll(prices, "See Details >", "")
		prices = strings.ReplaceAll(prices, "Prices: ", "")
		log.Info("-- prices --> ", prices)

		Temperament := HtmlText.Find("div.list-3 p").Text()
		split := camelcase.Split(Temperament)
		log.Info("-- Temperament --> ", split)

		prossessnum++

	})

	c.Visit("https://www.dogbreedslist.info/all-dog-breeds/")

	for i := 2; i <= 8; i++ {
		c.Visit("https://www.dogbreedslist.info/all-dog-breeds/page-" + strconv.Itoa(i) + ".html")
	}

	log.Info("==========Number of dog breed prossed =============================")
	log.Info(prossessnum)

}

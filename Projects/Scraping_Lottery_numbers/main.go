package main

import (
	"fmt"
	num "lottery/Utile"
	"strconv"

	"github.com/gocolly/colly"
)

var pl = fmt.Println
var prossessnum uint = 0

func main() {

	// Creating a collector
	c := colly.NewCollector()

	// Set error handler
	c.OnError(func(_ *colly.Response, err error) {
		fmt.Println("Error:", err.Error())
	})

	// Checking URL response.
	c.OnResponse(func(r *colly.Response) {
		pl("The URL is response well: ", r.Request.URL)
	})

	// Scraping the web tag <tr>
	c.OnHTML("tr", func(e *colly.HTMLElement) {
		q := e.DOM
		// Scraping one string, of the 5 winning numbers.
		winNumsStr := q.Find("li.ball").Text()
		winNumb := new(num.WinNum)
		winNumb.WinNumbers(winNumsStr)

		// Scraping the powerball number.
		powerballstr := q.Find("li.powerball").Text()
		// Scraping the date of game.
		date := q.Find("a").Text()

		winNumb.WinPerMoth(powerballstr, date)

		// Counting the number of games processed
		prossessnum++

	})

	// Scraping all pages, each Html page is a whole year games, start 2010 unit now.
	for i := 2010; i <= 2022; i++ {
		c.Visit("https://www.lottonumbers.com/powerball-results-" + strconv.Itoa(i))
	}

	pl("=======================================")
	pl(num.WinPerMonth)
	pl("=======================================")
	pl(prossessnum)
	pl("=======================================")
	pl(num.PowerballPerMoth)
	pl("=======================================")

	var num69 [69]int

	for i := 0; i <= 68; i++ {
		for j := 0; j <= 11; j++ {
			num69[i] = num69[i] + num.WinPerMonth[j][i]
		}

	}
	pl(num69)

}

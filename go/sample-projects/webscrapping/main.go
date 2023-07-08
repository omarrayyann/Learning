package main

import (
	"fmt"

	"github.com/gocolly/colly"
)

type Quote struct {
	Quote  string
	Author string
}

func main() {
	var quotes []Quote
	c := colly.NewCollector(colly.AllowedDomains("quotes.toscrape.com"))

	c.OnRequest(func(r *colly.Request) {
		r.Headers.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/15.1 Safari/605.1.15")
		fmt.Println("Visiting: ", r.URL)
	})

	c.OnResponse(func(r *colly.Response) {
		fmt.Println("Status Code: ", r.StatusCode)
	})

	c.OnError(func(r *colly.Response, err error) {
		fmt.Println("Error: ", err.Error())
	})

	c.OnHTML(".quote", func(h *colly.HTMLElement) {
		div := h.DOM
		quote := div.Find(".text").Text()
		author := div.Find(".author").Text()
		q := Quote{Quote: quote, Author: author}
		quotes = append(quotes, q)
	})

	c.Visit("http://quotes.toscrape.com")

	fmt.Println(quotes)

}

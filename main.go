package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"

	"github.com/gocolly/colly"
)

//Quotes nos permite acceder a los quotes de la página
type Quotes struct {
	QUOTE string
	TITLE string
	TAGS  []string
}

func main() {
	citas := make([]Quotes, 0)

	c := colly.NewCollector(
		colly.AllowedDomains("quotes.toscrape.com", "www.quotes.toscrape.com"),
	)

	c.OnRequest(func(r *colly.Request) {
		r.Headers.Set("content-type", "application/json; charset=utf-8")
		fmt.Println("Visiting", r.URL)
	})

	// Find and visit next page links
	c.OnHTML(`.next a[href]`, func(e *colly.HTMLElement) {
		link := e.Attr("href")
		// Print link
		fmt.Printf("Link found: %s\n", link)
		e.Request.Visit(link)
	})

	// Find quotes with colly
	c.OnHTML("div[class=quote]", func(e *colly.HTMLElement) {
		log.Println("Quote quote", e.Request.URL)

		quote := e.ChildText("span.text")
		if quote == "" {
			log.Println("No quote found", e.Request.URL)
		}
		//fmt.Printf("Quote: %s \n ", quote)

		cita := Quotes{
			QUOTE: quote,
			TITLE: e.ChildText("small.author"),
			TAGS:  e.ChildTexts("div.tags > a"),
		}

		citas = append(citas, cita)
		jsonData, err := json.Marshal(citas)
		if err != nil {
			return
		}
		ioutil.WriteFile("quotes.json", jsonData, 0644)
	})

	c.Visit("http://quotes.toscrape.com/")

}

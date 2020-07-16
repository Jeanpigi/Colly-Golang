package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"

	"github.com/gocolly/colly"
)

//Quotes is a struct what save data will receive
type Quotes struct {
	QUOTE string
	TITLE string
	TAGS  []string
}

func main() {
	citas := make([]Quotes, 0)
	//Instantiate default collector colly framework
	c := colly.NewCollector(
		colly.AllowedDomains("quotes.toscrape.com", "www.quotes.toscrape.com"),
	)

	//Before making a request print "Visiting ..."
	c.OnRequest(func(r *colly.Request) {
		r.Headers.Set("content-type", "application/json; charset=utf-8")
		fmt.Println("Visiting", r.URL)
	})

	// Find quotes with colly
	c.OnHTML("div[class=quote]", func(e *colly.HTMLElement) {
		log.Println("Quote quote", e.Request.URL)

		cita := Quotes{
			QUOTE: e.ChildText("span.text"),     // On every a span element which has text on his class call callback
			TITLE: e.ChildText("small.author"),  // On every a small element which has author on his class call callback
			TAGS:  e.ChildTexts("div.tags > a"), // On every a div which has a call callback
		}

		citas = append(citas, cita)
		//This is way for turning data received to json
		jsonData, err := json.Marshal(citas)
		if err != nil {
			return
		}
		ioutil.WriteFile("quotes.json", jsonData, 0644)
	})

	// Find and visit next page links
	c.OnHTML(`.next a[href]`, func(e *colly.HTMLElement) {
		link := e.Attr("href")

		//In case do not find a link
		if link == "" {
			log.Println("No link found", e.Request.URL)
		}
		// Print link
		fmt.Printf("Link found: %s\n", link)
		e.Request.Visit(link)
	})

	//Start scraping on http://quotes.toscrape.com/
	c.Visit("http://quotes.toscrape.com/")

}

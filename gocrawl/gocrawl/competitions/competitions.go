package competitions

import (
	"fmt"
	"github.com/geziyor/geziyor"
	"github.com/geziyor/geziyor/client"
	"githug.com/PuerkitoBio/goquery"
)

//TODO: come up with a strategy to handle maps that are created by the parse functions
//WARN: make sure to read and understand the extract option for extracting data and determine whether we can use it to minimise allocatins

// returns a map of leagues names and links to pages where we can get data
// about the competition(leagues)
func GetLeagues() {
	scraper := geziyor.NewGeziyor(&geziyor.Options{
		AllowedDomains:    []string{"https:www.//sofascore.com"},
		StartURLs:         []string{"https://www.sofascore.com"},
		RobotsTxtDisabled: false,
		ParseFunc:         legueParse,
	})
	scraper.Start()

}

// returns a map of cup names and links to pages where we can get data
// about the competition(cups)
func GetCups() {
	scraper := geziyor.NewGeziyor(&geziyor.Options{
		AllowedDomains:    []string{"https:www.//sofascore.com"},
		StartURLs:         []string{"https://www.sofascore.com"},
		RobotsTxtDisabled: false,
		ParseFunc:         cupParse,
	})
	scraper.Start()

}

//WARN:handle errors that might occur when sending requests e.g an invalid response code

func leagueParse(g *geziyor.Geziyor, r *client.Response) {
	//parse html here and extract data
	var leagues map[string]string
	leagues = make(map[string]string)
	r.HTMLDoc.Find("a").Each(func(_ int, g *goquery.Selection) {
		//get link text for competitions and add them to leagues map

	})
}

func cupParse(g *geziyor.Geziyor, r *client.Response) {
	//parse html here and extract data
	var cups map[string]string
	cups = make(map[string]string)
	r.HTMLDoc.Find("insert selector here").Each(func(i int, s *goquery.Selection) {
		//get link and put it into the cups map
	})
}

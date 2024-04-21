package main

import (
	"fmt"
	"github.com/JESSE-SOTERIA/gocrawl/stats.git"
	"github.com/gocolly/colly"
	"strings"
)

func main() {
	domain := "https://www.sofascore.com"
	competitions := make(map[string]string)
	var wantedComps = []string{"premier-league", "italy", "laliga", "bundesliga", "champions-league", "europa", "brazil"}
	c := colly.NewCollector()
	colly.AllowedDomains(domain)
	c.OnHTML("a", func(e *colly.HTMLElement) {
		for _, league := range wantedComps {
			link := e.Attr("href")
			if strings.Contains(link, league) {
				//solve the id thing for links to leagues
				competitions[league] = link
			}
		}

	})
	c.Visit(domain)

	fmt.Println(competitions)

	c.OnHTML("a", func(e *colly.HTMLElement) {
		link := e.Attr("href")
		if strings.Contains(link, "team") {
			fmt.Println(link)
		}
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println(r.URL.String())
	})

	//print the team links for every league competition
	for _, comp := range competitions {
		c.Visit(comp)
	}
}

func makeLeagueTable(link string, scrapper *colly.Collector) {
	scrapper.OnHTML("", func(element *colly.HTMLElement) {

	})
}

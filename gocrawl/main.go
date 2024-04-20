package main

import (
	"fmt"
	"github.com/gocolly/colly"
	"strings"
)

func main() {
	domain := "https://www.sofascore.com"
	competitions := make(map[string]string)
	//names should correspond to the actual names in thee links in sofascore
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

	c.Visit(competitions["europa"])

}

// returns teams in the league in terms of the links to the actual data for the team(maybe should return a map of team name and a link to the team data???)
// this map will be used to rescrape the data for teams on an interval.
//func getLeague(string) map[string]string {

//}

//the league pages have the same structure
// we can have a single function that gets data from a league page when called
//this should be called for every single league
//the league links should be in a map
//the function should be called for every member of the map.

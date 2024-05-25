package main

import (
	"fmt"
	"github.com/JESSE-SOTERIA/gocrawl/competitions"
	"github.com/PuerkitoBio/goquery"
	"github.com/geziyor/geziyor"
	"github.com/geziyor/geziyor/client"
)

//WARN: make find a way to map data from webite into the database
//make the database api so it could be used by the statistics package
//make stats package so that it could be used by main website.

func main() {
	//get competitions package initializes comp maps with comp names as keys
	//and links to pages as values
	fmt.Println("this is the js rendered scraper")
	leagues, err := competitions.GetLeagues()

	//TODO: make sure to handle the error better: give more informative mesage
	if err != nil {
		fmt.Println("there has been an error getting leagues")
	}
	cups, err := competitions.Getcups()
	//TODO: make sure to handle the error better: give more informative mesage
	if err != nil {
		fmt.Println("there has been an error getting cups")
	}

	//iterate through leagues and cups to get competition  stats
	//do the samee to get comp teams and
	//iterate through tems to  get team data incl players
}

package stats

import (
	"fmt"
)

// these are the type that dictate the data structures for storing a teams statistics
type leagueStats struct {
	won      string
	lost     string
	played   string
	goals    string
	conceded string
	points   int
}

type team struct {
	name         string
	competitions []string
	leagueStats
}

func print() {
	fmt.Println("this is the stats package")
}

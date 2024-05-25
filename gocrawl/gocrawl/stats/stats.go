package stats

import ()

// these are the type that dictate the data structures for storing a teams statistics
type LeagueStats struct {
	Won      int
	Lost     int
	Played   int
	Goals    int
	Conceded int
	Points   int
}
type player struct {
	name        string
	team        string
	nationality string
	avgrating   float32
	played      int
	goals       int
	assists     int
	position    string
	started     int
}
type defender struct {
	tackles     int
	cleansheets int
	player
}
type game struct {
	comp    string
	home    string
	away    string
	score   string
	scorers []string
	asists  []string
}

type Team struct {
	Name         string
	Competitions []string
	LeagueStats
}
type LeagueTable struct {
	Name            string
	Teams           []string
	ClSpots         []string
	EruopaSpots     []string
	RelegationSpots []string
}

// same type for uefa and europa because of the same format
type EuropeLeagueGroup struct {
	name  string
	teams []Team
}

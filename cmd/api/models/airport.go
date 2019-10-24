package models

import ( 
	"strconv"
)

// A COMPLETER QUAND REDIS OK
type Airport struct {
	Id string
}

type Airports []Airport

// -----
// /!\ EDITER LE CODE - CODE DE TEST
func NewAirport(i int) Airport{
	index := "ZZZ"+strconv.Itoa(i)
	return Airport{Id:index}
}

// /!\ EDITER LE CODE - CODE DE TEST
func ListAllAirport() *Airports {
	var airports Airports

	for i := 0; i < 10; i++ {
		var a Airport
		a = NewAirport(i)

		airports = append(airports, a)
	}
	return &airports
}

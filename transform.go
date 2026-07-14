package main

func BuildConcerts(dl map[string][]string) []Concert {

	var Concertslice []Concert
	for key, value := range dl {

		newConcert := Concert{}

		newConcert.Location = key
		newConcert.Dates = value

		Concertslice = append(Concertslice, newConcert)
	}
	return Concertslice
}

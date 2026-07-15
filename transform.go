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

func MergeArtistsWithRelations(artists []Artist, relations []Relation) []Artist {

	RelationByID := map[int]Relation{}

	for _, relation := range relations {

		RelationByID[relation.ID] = relation
	}

	sliceArtist := []Artist{}

	for _, artist := range artists {

		NewArtist := Artist{}

		matchedRelation := RelationByID[artist.ID]

		mapper := matchedRelation.DatesLocations
		Concertslice := BuildConcerts(mapper)

		NewArtist.Concerts = Concertslice

		sliceArtist = append(artist.Concerts, New)
	}
	return sliceArtist
}

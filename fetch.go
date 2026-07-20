package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

//return a slice of artist
func FetchArtists() ([]Artist, error) {

	resp, err := http.Get("https://groupietrackers.herokuapp.com/api/artists")

	if err != nil {
		return nil, fmt.Errorf("error: %w", err)
	}
	
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("something went wrong: %d", resp.StatusCode)
	}
	fmt.Println(http.ContentType)
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)

	if err != nil {
		return nil, fmt.Errorf("error reading body: %w", err)
	}

	var artist []Artist

	err = json.Unmarshal(data, &artist)

	if err != nil {
		return nil, fmt.Errorf("error while Unmarshaling: %w", err)
	}

	return artist, nil
}

func FetchRelations() ([]Relation, error) {

	resp, err := http.Get("https://groupietrackers.herokuapp.com/api/relation")

	if err != nil {
		return nil, fmt.Errorf("error: %w", err)
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("something went wrong: %d", resp.StatusCode)
	}

	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)

	if err != nil {
		return nil, fmt.Errorf("error reading body: %w", err)
	}

	var Wrapindx WrapIndx

	err = json.Unmarshal(data, &Wrapindx)

	if err != nil {
		return nil, fmt.Errorf("error while Unmarshaling: %w", err)
	}

	//Wrapindx.Index locate the value []Relation 
	sliceOfRelation := Wrapindx.Index
	return sliceOfRelation, nil
}

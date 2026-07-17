package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func FetchArtists() ([]Artist, error) {

	resp, err := http.Get("https://groupietrackers.herokuapp.com/api/artists")

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

	var artist []Artist

	err = json.Unmarshal(data, &artist)

	if err != nil {
		return nil, fmt.Errorf("error while Unmarshaling: %w", err)
	}

	return artist, nil
}

func FetchRelations() ([]Relation, error) {

	resp, err := http.Get("https://groupietrackers.herokuapp.com/api/relations")

	if err != nil {
		return nil, fmt.Errorf("error: %w", err)
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("something went wrong: %w", resp.StatusCode)
	}

	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)

	if err != nil {
		return nil, fmt.Errorf("error reading body: %w", err)
	}

	var relation []Relation

	err = json.Unmarshal(data, &relation)

	if err != nil {
		return nil, fmt.Errorf("error while Unmarshaling: %w", err)
	}

	return relation, nil
}

package main

import (
	"fmt"
	"io"
	"net/http"
)

func FetchArtists() ([]Artist, error) {

	resp, err := http.Get("https://groupietrackers.herokuapp.com/api/artists")

	if err != nil {
		return nil, fmt.Errorf("%w", err)
	}

	data, err := io.ReadAll(resp.Body)

	if err != nil {
		return nil, fmt.Errorf("%w", err)
	}

	defer resp.Body.Close()

	fmt.Println(string(data))

	return nil, nil
}



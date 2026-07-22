package main


import(
	"net/http"
	"fmt"
)

var Artists []Artist

func main(){

sliceArtists, err := FetchArtists()

if err != nil{
	return fmt.Errorf("error fetching artist: %w",  err)
}

sliceRelation, err := FetchRelations()

if err != nil{
	return fmt.Errorf("error fetching relation: %w", err)
}

Artists = MergeArtistsWithRelations(sliceArtist, sliceRelation)

	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/artist/", artistHandler)
	http.HandleFunc("/search", searchHandler)

	fmt.Println("server is live on http://localhost:8080")

	err := http.ListenAndServe(":8080", nil)

	if err != nil{
		fmt.Println(err)
		return 
	}
}
package main


import(
	"net/http"
	"fmt"
)

func main(){


	// http.HandleFunc("/", homeHandler)
	// http.HandleFunc("/artist/", artistHandler)
	// http.HandleFunc("/search", searchHandler)
	fmt.Print(FetchArtists())
	fmt.Println("server is live on http://localhost:8080")

	err := http.ListenAndServe(":8080", nil)

	if err != nil{
		fmt.Println(err)
		return 
	}
}
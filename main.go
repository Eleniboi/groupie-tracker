package main


import(
	"net/http"
	"fmt"
)

func main(){


	http.HandleFunc("/", homeHandler)
	fmt.Println("server is live on http://localhost:8080")

	err := http.ListenAndServe(":8080", nil)

	if err != nil{
		fmt.Println(err)
		return 
	}
}
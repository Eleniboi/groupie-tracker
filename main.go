package main 


import(
	"fmt"
)


type Concert struct{

	Location string
	Dates []string
}

type Artist struct{

	ID int	`json:id`
	Name string  `json:name`
	Image string
	CreationDate int
	FirstAlbum string
	Concerts []string
}

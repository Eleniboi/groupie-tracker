package main



type Concert struct {
	Location string
	Dates    []string
}

type Artist struct {
	ID           int      `json:"id"`
	Name         string   `json:"name"`
	Image        string   `json:"image"`
	CreationDate int      `json:"creationDate"`
	FirstAlbum   string   `json:"firstAlbum"`
	Members      []string `json:"members"`
	Concerts     []Concert
}

type Relation struct {
	ID             int                 `json:"id"`
	DatesLocations map[string][]string `json:"datesLocations"`
}

type WrapIndx struct{
   Index []Relation `json:"index"`
}
package main

import (
	"encoding/json"
	"log"
	"fmt"
)

type Movie struct {
	Title  string
	Year   int `json:"released"`
	Color  bool `json:"color,omitempty"`
	Actors []string
}

func main() {
	var movies = []Movie{
		{Title : "Casablanca", Year: 1942, Color: false,
			Actors: []string{"Humphrey Bogart", "Ingrid Bergman"}},
		{Title: "Cool Hand Luke", Year: 1967, Color: true,
			Actors: []string{"Paul Newman"}},
		{Title: "Bullitt", Year: 1968, Color: true,
			Actors: []string{"Steve McQueen", "Jacqueline Bisset"}},
	}

	data, err := json.MarshalIndent(movies, "", "	")
	if err != nil {
		log.Fatalf("json marshaling failed:%s", err)
	}
	fmt.Printf("%s\n", data)
	var actors []struct{Actors []string}
	if err = json.Unmarshal(data,&actors); err != nil{
		log.Fatalf("json unmarshaling failed:%s", err)
	}
	fmt.Println(actors)
}

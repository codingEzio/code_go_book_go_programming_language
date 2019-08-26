package main

import (
	"encoding/json"
	"fmt"
	"log"
)

// Only capitalized names will be marshaled :)
type Movie struct {
	Title  string `json:"title"`
	Year   int    `json:"released"`
	Rating int    `json:"rating,omitempty"` // if no vals given, omit it (0=>null)
	Sieres []int
}

func main() {
	var movies = []Movie{
		{
			Title:  "Wonder Woman",
			Year:   2017,
			Rating: 7,
			Sieres: []int{1, 2},
		},
		{
			Title:  "Interstellar",
			Year:   2014,
			Sieres: []int{1},
		},
	}

	// => JSON
	data, err := json.MarshalIndent(movies, "", "    ")
	if err != nil {
		log.Fatalf("JSON marshaling failed\n%s", err)
	}
	_ = fmt.Sprintf("%s\n", data)

	// => Object
	var movieUnmarshalled []struct {
		Title  string `json:"title"`
		Rating int    `json:"rating,omitempty"`
	}
	if err := json.Unmarshal(data, &movieUnmarshalled); err != nil {
		log.Fatalf("JSON unmarshaling failed\n%s", err)
	}
	fmt.Printf("%v\n", movieUnmarshalled)
}

package server

import (
	"encoding/json"
	"fmt"
	"net/http"

	"betty/fb88/api/types"
)

func GetMatches(date string) (matches types.Events, error error) {
	url := fmt.Sprintf("https://api.sofascore.com/api/v1/sport/football/scheduled-events/%s", date)
	response, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
	}
	err = json.NewDecoder(response.Body).Decode(&matches)
	if err != nil {
		fmt.Printf("error decoding sakura response: %v", err)
		if e, ok := err.(*json.SyntaxError); ok {
			fmt.Printf("syntax error at byte offset %d", e.Offset)
		}
		fmt.Printf("sakura response: %q", response.Status)
	}
	return matches, err
}

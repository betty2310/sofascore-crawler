package server

import (
	"betty/fb88/api/types"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/fatih/color"
)

func GetTable() (table types.TableData, error error) {
	url := "https://api.sofascore.com/api/v1/tournament/1/season/52186/standings/total"
	response, err := http.Get(url)
	if err != nil {
		color.Red("Error fetching: %v\n\nMaybe you get banned :v", err)
	}

	err = json.NewDecoder(response.Body).Decode(&table)
	if err != nil {
		fmt.Println(err)
	}
	return table, err
}

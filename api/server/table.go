package server

import (
	"betty/fb88/api/types"
	"encoding/json"
	"fmt"
	"net/http"
)

func GetTable() (table types.TableData, error error) {
	url := "https://api.sofascore.com/api/v1/tournament/1/season/52186/standings/total"
	response, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
	}

	err = json.NewDecoder(response.Body).Decode(&table)
	if err != nil {
		fmt.Println(err)
	}
	return table, err
}

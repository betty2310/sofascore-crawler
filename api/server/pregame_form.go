package server

import (
	"betty/fb88/api/types"
	"encoding/json"
	"fmt"
	"net/http"
)

func GetPregameForm(id int) (statistic types.PregameFormData, error error) {
	url := fmt.Sprintf("https://api.sofascore.com/api/v1/event/%d/pregame-form", id)
	response, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
	}

	err = json.NewDecoder(response.Body).Decode(&statistic)
	if err != nil {
		fmt.Println(err)
	}
	return statistic, err
}

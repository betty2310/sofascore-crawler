package server

import (
	"betty/fb88/api/types"
	"encoding/json"
	"fmt"
	"net/http"
)

func GetH2H(id int) (h2h types.H2H, error error) {
	url := fmt.Sprintf("https://api.sofascore.com/api/v1/event/%d/h2h", id)
	response, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
	}

	err = json.NewDecoder(response.Body).Decode(&h2h)
	if err != nil {
		fmt.Println(err)
	}
	return h2h, err
}

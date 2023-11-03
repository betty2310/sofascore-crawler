package server

import (
	"betty/fb88/api/types"
	"encoding/json"
	"fmt"
	"net/http"
)

func GetStatistic(id int) (statistic types.StatisticsData, error error) {
	url := fmt.Sprintf("https://api.sofascore.com/api/v1/event/%d/statistics", id)
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

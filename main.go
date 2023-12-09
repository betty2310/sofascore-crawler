package main

import (
	"betty/fb88/api/server"
	"betty/fb88/api/types"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"time"

	"github.com/fatih/color"
)

const DEFAULT_PATH = "data"

func writeJson(data any, filename string) {
	// Create any missing directories in the path
	err := os.MkdirAll(filepath.Dir(filename), 0755)
	if err != nil {
		panic(err)
	}

	// Try to create a new file
	file, err := os.Create(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	err = encoder.Encode(data)
	if err != nil {
		panic(err)
	}
}

func fetchMatches(date string) types.Events {
	color.Blue("Crawl all matches on %s", date)
	matches, err := server.GetMatches(date)
	if err != nil {
		log.Println(err)
	}
	matchesFilePath := filepath.Join(DEFAULT_PATH, strings.Replace(date, "-", "_", -1), "matches.json")
	writeJson(matches, matchesFilePath)
	color.Green("✅ Successfully crawl matches on %s", date)
	return matches
}

func main() {
	start := time.Now()
	defer func() {
		color.Green("\nExecution Time: %f", time.Since(start).Seconds())
	}()

	startDate := time.Date(2021, 8, 14, 0, 0, 0, 0, time.UTC)
	endDate := time.Date(2022, 5, 22, 0, 0, 0, 0, time.UTC)

	outer_wg := sync.WaitGroup{}
	for currentDate := startDate; currentDate.Before(endDate) || currentDate.Equal(endDate); currentDate = currentDate.AddDate(0, 0, 1) {
		date := currentDate.Format("2006-01-02")
		outer_wg.Add(1)
		go func(date string) {
			defer outer_wg.Done()
			fmt.Println(date)
			mc := make(chan types.Events)

			go func() {
				mc <- fetchMatches(date)
			}()

			matches := <-mc

			wg := sync.WaitGroup{}
			color.Blue("Crawl match information")
			for _, event := range matches.Events {
				if event.Tournament.Name == "Premier League" && event.Tournament.ID == 1 {
					wg.Add(1)
					filePath := filepath.Join(DEFAULT_PATH, strings.Replace(date, "-", "_", -1), "matches")
					go func(event types.Event) {
						defer wg.Done()
						if event.Status.Type == "inprogress" {
							filePath = filepath.Join(filePath, "in_progress_matches")
						}
						filePath = filepath.Join(filePath, event.Slug)

						// statistic
						color.Yellow("\tCrawl match %s statistic", event.Slug)
						statistic, err := server.GetStatistic(event.ID)
						if err != nil {
							log.Println(err)
						}
						name := fmt.Sprint(event.ID) + "_statistic.json"
						st := filepath.Join(filePath, name)
						writeJson(statistic, st)

						// lineup
						color.Cyan("\tCrawl match %s lineup", event.Slug)
						lineup, err := server.GetLineups(event.ID)
						if err != nil {
							log.Println(err)
						}
						name = fmt.Sprint(event.ID) + "_lineup.json"
						lineupFilePath := filepath.Join(filePath, name)
						writeJson(lineup, lineupFilePath)

						// h2h
						fmt.Println("\tCrawl match", event.Slug, "h2h")
						h2h, err := server.GetH2H(event.ID)
						if err != nil {
							log.Println(err)
						}
						name = fmt.Sprint(event.ID) + "_h2h.json"
						h2hFilePath := filepath.Join(filePath, name)
						writeJson(h2h, h2hFilePath)

						// pregame form
						color.Magenta("\tCrawl match %s pregame form", event.Slug)
						pre, err := server.GetPregameForm(event.ID)
						if err != nil {
							log.Println(err)
						}
						name = fmt.Sprint(event.ID) + "_pregame_form.json"
						preFilePath := filepath.Join(filePath, name)
						writeJson(pre, preFilePath)
					}(event)
				}
			}
			wg.Wait()
		}(date)
	}
	outer_wg.Wait()
}

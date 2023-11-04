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
	"time"

	"github.com/fatih/color"
	"github.com/joho/godotenv"
)

const DEFAULT_PATH = "output"

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
	return matches
}

func main() {
	start := time.Now()
	defer func() {
		color.Green("\nExecution Time: %f", time.Since(start).Seconds())
	}()
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	date := os.Getenv("DATE")
	matches := fetchMatches(date)
	color.Green("✅ Successfully crawl matches on %s", date)

	color.Blue("Crawl tournament table on %s", date)
	table, err := server.GetTable()
	if err != nil {
		log.Println(err)
	}
	name := strings.Replace(table.Standings[0].Name, "/", "-", -1) + "_table.json"
	tableFilePath := filepath.Join(DEFAULT_PATH, strings.Replace(date, "-", "_", -1), name)
	writeJson(table, tableFilePath)
	color.Green("✅ Successfully crawl %s table!", table.Standings[0].Name)

	color.Blue("Crawl tournament statistic on %s", date)
	statistic, err := server.GetTournamentStatistic()
	if err != nil {
		log.Println(err)
	}
	name = strings.Replace(table.Standings[0].Name, "/", "-", -1) + "_overall_statistic.json"
	sFilePath := filepath.Join(DEFAULT_PATH, strings.Replace(date, "-", "_", -1), name)
	writeJson(statistic, sFilePath)
	color.Green("✅ Successfully crawl %s overall statistic!", table.Standings[0].Name)

	for _, event := range matches.Events {
		if event.Tournament.Name == "Premier League" && event.Tournament.ID == 1 {
			filePath := filepath.Join(DEFAULT_PATH, strings.Replace(date, "-", "_", -1), "matches")
			if event.Status.Type == "inprogress" {
				filePath = filepath.Join(filePath, "inprogress")
				// inprogress match statistic
				fmt.Println("\tCrawl match", event.Slug, "statistic")
				statistic, err := server.GetStatistic(event.ID)
				if err != nil {
					log.Println(err)
				}
				name := fmt.Sprint(event.ID) + "_statistic.json"
				st := filepath.Join(filePath, name)
				writeJson(statistic, st)
			}
			filePath = filepath.Join(filePath, event.Slug)

			// lineup
			fmt.Println("\tCrawl match", event.Slug, "lineup")
			lineup, err := server.GetLineups(event.ID)
			if err != nil {
				log.Println(err)
			}
			name := fmt.Sprint(event.ID) + "_lineup.json"
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
			fmt.Println("\tCrawl match", event.Slug, "pregame form")
			pre, err := server.GetPregameForm(event.ID)
			if err != nil {
				log.Println(err)
			}
			name = fmt.Sprint(event.ID) + "_pregame_form.json"
			preFilePath := filepath.Join(filePath, name)
			writeJson(pre, preFilePath)
		}
	}

	color.Blue("Crawl team information")
	for _, team := range table.Standings[0].Rows {
		filePath := filepath.Join(DEFAULT_PATH, strings.Replace(date, "-", "_", -1), "Teams", team.Team.Name)
		color.Yellow("\tCrawl %s overall statistic", team.Team.Name)
		teamSt, err := server.GetTeamStatistic(team.Team.ID)
		if err != nil {
			log.Println(err)
		}
		preFilePath := filepath.Join(filePath, "overall_statistic.json")
		writeJson(teamSt, preFilePath)

		color.White("\tCrawl %s player overall", team.Team.Name)
		teamPl, err := server.GetTeamPlayerOverall(team.Team.ID)
		if err != nil {
			log.Println(err)
		}
		plFilePath := filepath.Join(filePath, "player.json")
		writeJson(teamPl, plFilePath)
	}
}

package main

import (
	"betty/fb88/api/server"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"

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

func main() {

	start := time.Now()
	defer func() {
		fmt.Println("Execution Time: ", time.Since(start).Seconds())
	}()
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	date := os.Getenv("DATE")
	fmt.Println("DATE", date)
	matches, err := server.GetMatches(date)
	if err != nil {
		log.Println(err)
	}
	fmt.Println("Crawl all matches on", date)
	matchesFilePath := filepath.Join(DEFAULT_PATH, strings.Replace(date, "-", "_", -1), "matches.json")
	fmt.Println(matchesFilePath)
	writeJson(matches, matchesFilePath)
	// for _, event := range matches.Events {
	// 	if event.Tournament.Name == "Premier League" {
	// 		go func(id int) {
	// 			fmt.Println("\tCrawl match ", event.Slug, " lineup")
	// 			lineup, err := server.GetLineups(id)
	// 			if err != nil {
	// 				log.Println(err)
	// 			}

	// 			writeJson()
	// 		}(event.ID)
	// 	}
	// }

}

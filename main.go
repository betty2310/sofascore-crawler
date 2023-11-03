package main

import (
	"betty/fb88/api/server"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
)

func writeJson(data any, filename string) {
	file, err := os.Create(filename)
	if err != nil {
		println(err)
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
	lineup, err := server.GetTable()
	if err != nil {
		log.Println(err)
	}

	writeJson(lineup, "lineup.json")
}

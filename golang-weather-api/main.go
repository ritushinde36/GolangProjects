package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	//load the environment variables.
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal(err)
	}

	BASE_URL := os.Getenv("BASE_URL")
	API_KEY := os.Getenv("API_KEY")
	CITY := "london"

	SEARCH_URL := BASE_URL + "?appid=" + API_KEY + "&q=" + CITY

	resp, err := http.Get(SEARCH_URL)
	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		weatherBytes, _ := io.ReadAll(resp.Body)
		fmt.Print(string(weatherBytes))

	}
}

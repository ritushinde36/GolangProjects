package data

import (
	"encoding/json"
	"log"
	"os"
	"quotes-scraper/crawler"
)

func CreateJSON(tag string) {
	filename := tag + "_tag.json"

	json_data, _ := json.Marshal(crawler.Quotes)
	err := os.WriteFile(filename, json_data, 044)
	if err != nil {
		log.Fatal("unable to write to file")
	}

}

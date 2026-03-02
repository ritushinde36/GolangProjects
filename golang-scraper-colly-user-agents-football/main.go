package main

import (
	"colly-user-agents-football-scraper/crawler"
	"os"
	"strings"
)

func main() {

	//set the URL of the website that you want to scrape
	websiteURL := "https://www.scrapethissite.com/pages/forms/"

	//get the name of the team from the user as a command line argument
	userInput := os.Args
	userInput = userInput[1:]
	teamName := strings.Join(userInput, "+")

	// fmt.Println(teamName)

	//pass the name to the crawl function
	crawler.Crawl(websiteURL, teamName)

}

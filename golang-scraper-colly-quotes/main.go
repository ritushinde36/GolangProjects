package main

import (
	"fmt"
	"log"
	"os"
	"quotes-scraper/crawler"
	"quotes-scraper/data"
)

func main() {
	websiteURL := "https://quotes.toscrape.com/"

	//this is the list of all the valid tags present of the website
	// validTags := []string{"love", "inspirational", "life", "humor", "books", "reading", "friendship", "friends", "truth", "simile"}

	//check if the len of the agruments passed to the command line are more than one
	if len(os.Args) > 2 {
		log.Fatal("Only one argument is allowed")
	}

	//get the name of the tag that the user will pass in the command line
	userTag := os.Args[1]

	// check if the userTag is present in the the validTags slice
	// isValid := checkIfUserTagValid(validTags, userTag)

	// if !isValid {
	// 	log.Fatal("Invalid tag\nThe accepted tags are :- ", validTags)
	// }
	pageCount := 1

	for {
		noQuotes := crawler.Crawl(userTag, websiteURL, pageCount)
		if noQuotes {
			if pageCount >= 2 {
				fmt.Printf("Found all the Quotes for tag %v", userTag)
				data.CreateJSON(userTag)

			} else {
				log.Fatal("No Quotes found")
			}
			return
		}
		pageCount++

	}

}

func checkIfUserTagValid(validTags []string, userTag string) bool {
	for _, tag := range validTags {
		if tag == userTag {
			return true
		}
	}

	return false
}

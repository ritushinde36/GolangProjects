package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"golang.org/x/net/html"
)

func getHref(t html.Token) (bool, string) {
	for _, attributes := range t.Attr {
		if attributes.Key == "href" {
			return true, attributes.Val
		}
	}
	return false, ""
}

func scrape(URL string, chanURLs chan string, chanFinished chan bool) {
	resp, err := http.Get(URL)

	defer func() {
		chanFinished <- true
	}()

	if err != nil {
		log.Fatal("error parsing the URL ", URL)
	}

	body := resp.Body

	defer body.Close()

	tokenizer := html.NewTokenizer(body)

	for {
		tt := tokenizer.Next()

		switch tt {
		case html.ErrorToken:
			return
		case html.StartTagToken:
			t := tokenizer.Token()
			isAnchorTag := t.Data == "a"
			if !isAnchorTag {
				continue
			}
			ok, anchorUrl := getHref(t)
			if !ok {
				continue
			}
			//do validation for URL
			if strings.HasPrefix(anchorUrl, "http") {
				chanURLs <- anchorUrl
			}
		}
	}
}

func main() {
	//get all the URLs that are passed to the go program
	seedURLs := os.Args[1:]

	//store all the found urls present on the seed url site in a map key=url_name and value=true, if we have found the URL
	foundURLs := make(map[string]bool)

	//channel to store all the found URLs
	chanURLs := make(chan string)

	//channel to say that we have finished scraping the body of the URL
	chanFinished := make(chan bool)

	for _, URL := range seedURLs {
		go scrape(URL, chanURLs, chanFinished)
	}

	for i := 0; i < len(seedURLs); {
		select {
		case url := <-chanURLs:
			foundURLs[url] = true
		case <-chanFinished:
			i++
		}

	}

	fmt.Printf("The number of found URLs is %v \n", len(foundURLs))
	for key, _ := range foundURLs {
		fmt.Println(key)

	}

	close(chanURLs)

}

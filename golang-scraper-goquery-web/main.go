package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/PuerkitoBio/goquery"
)

type blog struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}

var (
	filename string = "blogs.json"
)

func main() {
	URL := "https://www.golangcode.com/"
	err := getDataFromWebsite(URL)
	if err != nil {
		log.Fatal(err)
	}
}

func getDataFromWebsite(URL string) error {
	//get the response body from the URL
	resp, err := http.Get(URL)
	if err != nil {
		return err
	}

	body := resp.Body
	defer resp.Body.Close()
	doc, err := goquery.NewDocumentFromReader(body)
	if err != nil {
		return err
	}

	var titles []string
	var contents []string
	var blogs []blog

	doc.Find(".post-title").Each(func(i int, s *goquery.Selection) {
		titles = append(titles, s.Text())
	})

	doc.Find("p").Each(func(i int, s *goquery.Selection) {
		contents = append(contents, s.Text())
	})

	for i := 0; i < len(titles); i++ {
		newBlog := blog{
			Title:   titles[i],
			Content: contents[i],
		}
		blogs = append(blogs, newBlog)
	}

	json_data, _ := json.Marshal(blogs)
	err = os.WriteFile(filename, json_data, 0644)
	if err != nil {
		return err
	}

	return nil

}

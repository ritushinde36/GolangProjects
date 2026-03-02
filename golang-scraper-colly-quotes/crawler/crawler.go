package crawler

import (
	"fmt"
	"quotes-scraper/models"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/gocolly/colly"
)

var Quotes []models.Quote

func Crawl(userTag string, websiteURL string, pageCount int) bool {

	page := strconv.Itoa(pageCount)

	//create the url with the correct tag and pages
	websiteUrlWithTag := websiteURL + "tag/" + userTag + "/page/" + page + "/"
	// fmt.Println(websiteUrlWithTag)

	if checkForNoQuotes(websiteUrlWithTag) {
		// fmt.Print("No quotes Found")
		return true
	}

	c := colly.NewCollector()

	c.OnHTML("div.quote", func(h *colly.HTMLElement) {
		// var quote models.Quote
		var tags = []string{}

		text := h.ChildText("span.text")
		// author := h.ChildText("small.author")
		author_url := h.ChildAttr("a", "href")
		h.DOM.Find("a").Each(func(_ int, s *goquery.Selection) {
			tag := s.Text()
			tags = append(tags, tag)
		})
		tags = append(tags[1:])

		author_name, born_date, born_location, description := CrawlAuthor(websiteURL, author_url)

		quote := models.Quote{
			Content: text,
			Author_details: models.Author{
				Author_name:   author_name,
				Born_date:     born_date,
				Born_Location: born_location,
				Desciption:    description,
			},
			Tags: tags,
		}

		Quotes = append(Quotes, quote)

		// fmt.Println("text:", text)
		// fmt.Println("author:", author)
		// fmt.Println("url:", author_url)
		// fmt.Println("tags:", tags)

		// fmt.Printf("-----\n\n")

	})

	c.Visit(websiteUrlWithTag)

	// fmt.Print(Quotes)

	return false

	//loop through the pages until you get "no quotes found"

	//store the quote into a struct for quotes and author

	// for the quotes it should go to tag_quotes.json

	//for author it should go to author.json

	//define the colly collector

}

func checkForNoQuotes(websiteUrl string) bool {
	flag := false
	c := colly.NewCollector()
	count := 0

	c.OnHTML("div.col-md-8", func(h *colly.HTMLElement) {
		content := h.Text
		if count == 0 {
			count++
			return
		}
		trimmedContent := strings.TrimSpace(content)
		// fmt.Printf("\n%v", trimmedContent)

		if strings.Contains(trimmedContent, "No quotes found!") {

			flag = true
		}

	})

	c.Visit(websiteUrl)

	return flag

}

func CrawlAuthor(websiteURL string, authorUrl string) (string, string, string, string) {
	var author_name, born_date, born_location, description string

	trimmedAuthorUrl := fmt.Sprint(authorUrl[1:])
	websiteAuthorUrl := websiteURL + trimmedAuthorUrl
	// fmt.Println(websiteAuthorUrl)

	c := colly.NewCollector()

	c.OnHTML("div.author-details", func(h *colly.HTMLElement) {
		author_name = h.ChildText("h3.author-title")
		born_date = h.ChildText("span.author-born-date")
		born_location = h.ChildText("span.author-born-location")
		description = h.ChildText("div.author-description")
		// fmt.Println("author_name:", author_name)
		// fmt.Println("born_date:", born_date)
		// fmt.Println("born_location:", born_location)
		// fmt.Println("description:", description)
	})

	c.Visit(websiteAuthorUrl)

	return author_name, born_date, born_location, description
}

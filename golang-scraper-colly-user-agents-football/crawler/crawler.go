package crawler

import (
	"colly-user-agents-football-scraper/csv"
	"colly-user-agents-football-scraper/models"
	"fmt"
	"math/rand"
	"time"

	"github.com/gocolly/colly"
)

var TeamInfo []models.Team_data

var userAgents = []string{
	"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/61.0.3163.100 Safari/537.36",
	"Mozilla/5.0 (Windows NT 6.1; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/61.0.3163.100 Safari/537.36",
	"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_12_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/61.0.3163.100 Safari/537.36",
	"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_12_6) AppleWebKit/604.1.38 (KHTML, like Gecko) Version/11.0 Safari/604.1.38",
	"Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:56.0) Gecko/20100101 Firefox/56.0",
	"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_13) AppleWebKit/604.1.38 (KHTML, like Gecko) Version/11.0 Safari/604.1.38",
}

func Crawl(websiteURL, teamName string) {

	// rand.Seed(time.Now().UnixNano())
	seed := time.Now().UnixNano()
	rng := rand.New(rand.NewSource(seed))

	toScrapeWebsiteUrl := websiteURL + "?q=" + teamName
	fmt.Print(toScrapeWebsiteUrl)

	c := colly.NewCollector()
	// count := 0

	c.OnRequest(func(r *colly.Request) {
		ua := userAgents[rng.Intn(len(userAgents))]
		r.Headers.Set("User-Agent", ua)
		fmt.Println("Using User-Agent:", ua)
	})

	c.OnHTML("tr.team", func(h *colly.HTMLElement) {

		name := h.ChildText("td.name")
		year := h.ChildText("td.year")
		wins := h.ChildText("td.wins")
		losses := h.ChildText("td.losses")
		otlosses := h.ChildText("td.ot-losses")
		win_percent := h.ChildText("td.pct[class*='text-']")
		goalsFor := h.ChildText("td.gf")
		goalsAgainst := h.ChildText("td.ga")
		plusMinus := h.ChildText("td.diff[class*='text-']")

		teamData := models.Team_data{
			TeamName:     name,
			Year:         year,
			Wins:         wins,
			Losses:       losses,
			OTLosses:     otlosses,
			WinPrecent:   win_percent,
			GoalsFor:     goalsFor,
			GoalsAgainst: goalsAgainst,
			PlusMinus:    plusMinus,
		}

		TeamInfo = append(TeamInfo, teamData)

		// fmt.Println(name)
		// fmt.Println(year)
		// fmt.Println(wins)
		// fmt.Println(losses)
		// fmt.Println(otlosses)
		// fmt.Println(win_percent)
		// fmt.Println(goalsFor)
		// fmt.Println(goalsAgainst)
		// fmt.Println(plusMinus)
		// fmt.Println("--------------------")

	})

	c.Visit(toScrapeWebsiteUrl)

	// fmt.Print(TeamInfo)
	csv.WriteToCSV(TeamInfo, teamName)
}

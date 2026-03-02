package csv

import (
	"colly-user-agents-football-scraper/models"
	"encoding/csv"
	"log"
	"os"
)

func WriteToCSV(TeamInfo []models.Team_data, teamName string) {
	filename := teamName + ".csv"
	file, err := os.Create(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	writer.Write([]string{"TeamName", "Year", "Wins", "Losses", "OTLosses", "WinPrecent", "GoalsFor", "GoalsAgainst", "PlusMinus"})

	for _, info := range TeamInfo {
		writer.Write([]string{
			info.TeamName,
			info.Year,
			info.Wins,
			info.Losses,
			info.OTLosses,
			info.WinPrecent,
			info.GoalsFor,
			info.GoalsAgainst,
			info.PlusMinus,
		})
	}

}

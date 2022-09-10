package ttScraper

import (
	"log"
	"time"

	tt "github.com/Elfshot/tt-api-wrapper"
	mongo "github.com/Elfshot/tt-scraper/mongo"
)

func sotd(redo int) {
	defer func() {
		if r := recover(); r != nil {
			log.Println("Recovered:", r)
		}
	}()

	sotdData, err := tt.Get_Sotd()

	if err != nil {
		if redo >= 6 {
			log.Printf("Error getting SOTD (Attempts Stopped): %+v\n", redo-1)
			return
		} else {
			log.Printf("Error getting SOTD (Attempt %d): %+v\n", redo, err)
			time.Sleep(5 * time.Second)
			sotd(redo + 1)
		}
		return
	}

	t := time.Now()
	date := time.Date(t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), 0, 0, time.UTC)
	mongo.UpdateSotd(date, sotdData)
	log.Println("Sotd Scraper Finished")
}

func Sotd() {
	sotd(1)
}

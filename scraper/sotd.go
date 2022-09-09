package ttScraper

import (
	"log"
	"time"

	tt "github.com/Elfshot/tt-api-wrapper"
	mongo "github.com/Elfshot/tt-scraper/mongo"
)

func Sotd(redo ...int) {
	defer func() {
		if r := recover(); r != nil {
			log.Println("Recovered:", r)
		}
	}()

	sotd, err := tt.Get_Sotd()

	if err != nil {
		var iter int
		if len(redo) > 0 {
			iter = redo[0]
		} else {
			iter = 1
		}
		if iter >= 6 {
			log.Printf("Error getting SOTD (Attempts Stopped): %+v\n", iter-1)
			return
		} else {
			log.Printf("Error getting SOTD (Attempt %d): \n", iter)
			time.Sleep(5 * time.Second)
			Sotd(iter + 1)
		}
		return
	}

	t := time.Now()
	date := time.Date(t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), 0, 0, time.UTC)
	mongo.UpdateSotd(date, sotd)
	log.Println("Sotd Scraper Finished")
}

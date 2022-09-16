package main

import (
	"os"
	"time"

	ttapi "github.com/Elfshot/tt-api-wrapper"
	mongo "github.com/Elfshot/tt-scraper/mongo"
	scraper "github.com/Elfshot/tt-scraper/scraper"
	gocron "github.com/go-co-op/gocron"
)

func main() {
	mongo.Init()
	ttapi.Init()

	cron := gocron.NewScheduler(time.UTC)

	if os.Getenv("DB_TEST") != "" {
		scraper.Players()
		scraper.DataAdv()
		scraper.Sotd()
	}
	cron.WaitForScheduleAll()

	cron.Every(2).Minutes().Do(scraper.Players)
	cron.Every(10).Minutes().Do(scraper.DataAdv)
	cron.Every(1).Day().At("00:11").Do(scraper.Sotd)

	cron.StartBlocking()
}

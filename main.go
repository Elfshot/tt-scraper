package main

import (
	"time"

	ttapi "github.com/Elfshot/tt-api-wrapper"
	mongo "github.com/Elfshot/tt-scraper/mongo"
	scraper "github.com/Elfshot/tt-scraper/scraper"
	gocron "github.com/go-co-op/gocron"
)

func main() {
	mongo.Init()
	ttapi.Init()

	//run players as blocking function to initialize the list of players for dataadv
	scraper.Players()

	cron := gocron.NewScheduler(time.Local)

	cron.Every(2).Minutes().Do(scraper.Players)
	cron.Every(10).Minutes().Do(scraper.DataAdv)

	cron.StartBlocking()
}

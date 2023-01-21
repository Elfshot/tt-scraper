package ttScraper

import mongo_m "github.com/Elfshot/tt-scraper/mongo/models"

func Init() {
	recentPlayers = make([]mongo_m.UsersCollModel, 0)
	Players(false)
}

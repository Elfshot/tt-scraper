package ttScraper

import (
	"log"
	"sync"
	"time"

	tt "github.com/Elfshot/tt-api-wrapper"
	mongo "github.com/Elfshot/tt-scraper/mongo"
	mongo_m "github.com/Elfshot/tt-scraper/mongo/models"
)

func DataAdv() {
	defer func() {
		if r := recover(); r != nil {
			log.Println("Recovered:", r)
		}
	}()

	if recentPlayers == nil {
		return
	}

	wg := sync.WaitGroup{}

	for _, player := range recentPlayers {
		wg.Add(1)
		time.Sleep(time.Millisecond * 90)
		go func(player mongo_m.UsersCollModel) {
			data, err := tt.Get_DataAdv(player.VrpId)
			if err != nil {
				log.Printf("Error getting data adv for %d: %+v\n", player.VrpId, err)
				return
			}
			mongo.UpdateDataAdv(player.VrpId, data, time.Now().UTC())
			wg.Done()
		}(player)
	}
	wg.Wait()
	log.Println("DataAdv scraper finished")
}

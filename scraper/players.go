package ttScraper

import (
	"log"
	"regexp"
	"strings"
	"sync"
	"time"

	tt "github.com/Elfshot/tt-api-wrapper"
	tt_m "github.com/Elfshot/tt-api-wrapper/models"
	mongo "github.com/Elfshot/tt-scraper/mongo"
	mongo_m "github.com/Elfshot/tt-scraper/mongo/models"
)

var recentPlayers []mongo_m.UsersCollModel

func Players() {
	defer func() {
		if r := recover(); r != nil {
			log.Println("Recovered:", r)
		}
	}()

	players, err := tt.GetTotalPlayers()
	if err != nil {
		log.Printf("Error getting total players: %+v\n", err)
		return
	}

	recentPlayers = make([]mongo_m.UsersCollModel, 0)
	timeDate := time.Now().UTC()
	wg := sync.WaitGroup{}
	re, regexErr := regexp.Compile(`[^\w]`)

	for _, playerR := range players {
		var discordId string
		searchName := playerR.Name

		if playerR.AvatarUrl != "" && strings.HasPrefix(playerR.AvatarUrl, "https://cdn.discordapp.com/avatars/") {
			discordId = strings.Split(playerR.AvatarUrl, "/")[4]
		}

		if regexErr != nil {
			searchName = playerR.Name
		} else {
			searchName = re.ReplaceAllString(playerR.Name, "")
		}

		player := mongo_m.UsersCollModel{
			UserName: playerR.Name, SearchName: searchName, LastFound: timeDate, DiscordId: discordId, VrpId: playerR.VrpId,
		}

		recentPlayers = append(recentPlayers, player)

		wg.Add(1)
		go func(playerR tt_m.BaseTotalPlayer) {
			mongo.UpdateUser(playerR.VrpId, player, timeDate)
			wg.Done()
		}(playerR)
	}
	wg.Wait()
	log.Println("Players scraper finished")
}

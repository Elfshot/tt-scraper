package mongo_models

import (
	"time"

	tt_m "github.com/Elfshot/tt-api-wrapper/models"
)

type UsersCollModel struct {
	VrpId      uint32    `bson:"vrpId,omitempty,minsize,truncate"`
	UserName   string    `bson:"userName,omitempty"`
	SearchName string    `bson:"searchName,omitempty"`
	DiscordId  string    `bson:"discordId,omitempty"`
	CountFound uint32    `bson:"countFound,omitempty,minsize,truncate"`
	FirstFound time.Time `bson:"firstFound,omitempty"`
	LastFound  time.Time `bson:"lastFound,omitempty"`
	// identifiers map[string]string `bson:"identifiers,omitempty,inline"`
}

// type usersCollModelIdentifiers {
// 	steam string
// 	license string
// 	license2 string
// 	live string
// 	fivem string
// 	...
// }

type DataAdvCollModel struct {
	VrpId uint32            `bson:"vrpId,omitempty,minsize,truncate"`
	Data  tt_m.UserDataData `bson:"data,omitempty"`
	Date  time.Time         `bson:"date,omitempty"`
}

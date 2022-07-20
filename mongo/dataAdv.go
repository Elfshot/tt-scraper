package mongo

import (
	"context"
	"log"
	"time"

	tt_m "github.com/Elfshot/tt-api-wrapper/models"
	models "github.com/Elfshot/tt-scraper/mongo/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var dataAdvCol *mongo.Collection

func UpdateDataAdv(vrpId uint32, userData *tt_m.UserData, date time.Time) {
	filter := bson.D{primitive.E{Key: "vrpId", Value: int32(vrpId)}}

	newData := models.DataAdvCollModel{
		VrpId: vrpId,
		Data:  userData.Data,
		Date:  date,
	}

	update := bson.D{primitive.E{Key: "$set", Value: newData}}

	_, err := dataAdvCol.UpdateOne(context.TODO(), filter, update, upOpts)
	if err != nil {
		log.Printf("DataAdv update error: %+v\n", err)
		return
	}
}

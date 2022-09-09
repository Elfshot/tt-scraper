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

var sotdCol *mongo.Collection

func UpdateSotd(date time.Time, sotdData *tt_m.Sotd) {
	filter := bson.D{primitive.E{Key: "timestamp", Value: date}}

	newData := models.SotdCollModel{
		Timestamp: date,
		Skill:     sotdData.Skill,
		Bonus:     sotdData.Bonus,
		Aptitude:  sotdData.Aptitude,
	}

	update := bson.D{primitive.E{Key: "$set", Value: newData}}

	_, err := sotdCol.UpdateOne(context.TODO(), filter, update, upOpts)
	if err != nil {
		log.Printf("Sotd update error: %+v\n", err)
		return
	}
}

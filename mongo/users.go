package mongo

import (
	"context"
	"log"
	"time"

	models "github.com/Elfshot/tt-scraper/mongo/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var usersCol *mongo.Collection

func UpdateUser(vrpId uint32, userData models.UsersCollModel, date time.Time) {
	filter := bson.D{primitive.E{Key: "vrpId", Value: int32(vrpId)}}

	// var result bson.M
	// usersCol.FindOne(context.TODO(), filter).Decode(&result)
	// if result == nil {
	// 	userData.FirstFound = date
	// 	log.Printf("Found a new user: %d!\n", vrpId)
	// }

	set := primitive.E{Key: "$set", Value: userData}
	inc := primitive.E{Key: "$inc", Value: bson.D{primitive.E{Key: "countFound", Value: 1}}}
	setOnInsert := primitive.E{Key: "$setOnInsert", Value: bson.D{primitive.E{Key: "firstFound", Value: date}}}

	update := bson.D{set, inc, setOnInsert}

	_, err := usersCol.UpdateOne(context.TODO(), filter, update, upOpts)
	if err != nil {
		log.Printf("User update error: %+v\n", err)
		return
	}
}

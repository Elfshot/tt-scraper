package mongo

import (
	"context"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var client *mongo.Client
var db *mongo.Database
var upOpts *options.UpdateOptions

const TimeFormat = "2006-01-02T15:04:05.000Z"

// coll := client.Database("Tycoon").Collection("usersTest")
// var result bson.M

// err = coll.FindOne(context.TODO(), bson.D{{"vrpId", 59504}}).Decode(&result)
// if err == mongo.ErrNoDocuments {
// 	fmt.Println("No document was found")
// 	return
// }
// if err != nil {
// 	log.Fatal(err)
// }
// jsonData, err := json.MarshalIndent(result, "", "    ")
// if err != nil {
// 	log.Fatal(err)
// }
// fmt.Printf("%s\n", jsonData)

func setClient() {
	uri := os.Getenv("MONGO_URI")
	if uri == "" {
		log.Fatal("No 'MONGO_URI' in env")
	}

	var err error
	client, err = mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		client.Disconnect(context.TODO())
		log.Fatal(err)
	}
}

func setDb() {
	tycoonDb := os.Getenv("DB_TITLE")
	db = client.Database(tycoonDb)
}

func GetCollection(wantedDb string) *mongo.Collection {
	res := db.Collection(wantedDb)
	return res
}

func setColls() {
	env := os.Getenv("DB_TEST")
	setter := func(db string) *mongo.Collection {
		if env != "" {
			db += "Test"
		}
		log.Printf("Setting collection %s\n", db)
		return GetCollection(db)
	}
	usersCol = setter("users")
	dataAdvCol = setter("dataadvs")
	sotdCol = setter("sotd")
}

func setOpts() {
	upOpts = options.Update()
	upOpts.SetUpsert(true)
}

func Init() {
	setClient()
	setDb()
	setOpts()
	setColls()
}

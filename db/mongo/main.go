package main

import (
	"context"
	"fmt"

	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

const version = "v3"

func main() {
	logrus.Infof("Starting program version #%v", version)

	ctx := context.Background()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		logrus.Fatal(err)
	}

	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		logrus.Fatal(err)
	}

	// db name users
	db := client.Database("users")

	// db collection users
	coll := db.Collection("users")

	cursor, err := coll.Find(ctx, bson.D{{"name", "Vasily"}, {"age", "30"}})
	if err != nil {
		logrus.Error(err)
		return
	}

	for cursor.Next(ctx) {
		fmt.Println(cursor.Current.String())
	}

	//res, err := coll.InsertOne(ctx, bson.D{
	//	{"name", "Anton"},
	//	{"age", "30"},
	//	{"email", "anton@gmail.com"},
	//	{"params", bson.D{
	//		{"weight", 70},
	//		{"height", 180},
	//	}},
	//})
	//if err != nil {
	//	logrus.Error(err)
	//	return
	//}

	//fmt.Println(res.InsertedID)
}

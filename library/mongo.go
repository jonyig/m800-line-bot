package library

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"m800-line-bot/config"
	"sync"
)

const (
	prodUri = "mongodb+srv://%s:%s@cluster0.qlumwac.mongodb.net/?retryWrites=true&w=majority"
	devUri  = "mongodb://%s:%s@localhost:27017/"
)

var (
	MongoInstance *mongo.Client
	once          sync.Once
)

func GetMongoDBInstance() *mongo.Client {
	once.Do(func() {
		MongoInstance = connectToMongo()
	})

	return MongoInstance
}

func connectToMongo() *mongo.Client {
	session, err := mongo.NewClient(
		options.Client().ApplyURI(
			getUrl(),
		),
	)

	if err != nil {
		log.Fatal(err)
	}

	err = session.Connect(context.TODO())
	if err != nil {
		log.Fatal(err)
	}

	return session
}

func getUrl() string {
	configuration := config.NewConfig()
	uri := prodUri
	if config.IsNotEnv() {
		uri = devUri
	}

	url := fmt.Sprintf(
		uri,
		configuration.GetMongoDBUsername(),
		configuration.GetMongoDBPassword(),
	)

	return url
}

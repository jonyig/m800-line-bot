package storage

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"m800-line-bot/library"
)

const (
	MongoDBDatabase   = "Cluster0"
	MongoDBCollection = "message"
	MongoDBUniqueCode = "user"
)

func SetMessage() {
	mongoClient := library.GetMongoDBInstance()
	collection := mongoClient.Database(MongoDBDatabase).Collection(MongoDBCollection)
	setUniqueIndex := mongo.IndexModel{
		Keys:    bson.D{{MongoDBUniqueCode, 1}},
		Options: options.Index().SetUnique(false),
	}

	_, err := collection.Indexes().CreateOne(context.TODO(), setUniqueIndex)
	if err != nil {
		log.Print(err)
	}
}

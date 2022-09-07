package repository

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"m800-line-bot/storage"
)

type LineBotMongoRepository struct {
	Client *mongo.Client
}

func NewLineBotMongoRepository(client *mongo.Client) *LineBotMongoRepository {
	return &LineBotMongoRepository{
		Client: client,
	}
}

func (r *LineBotMongoRepository) SaveMessage(userId string, message string) {
	type MessageInfo struct {
		UserId  string `json:"user_id" bson:"user_id"`
		Message string `json:"message" bson:"message"`
	}

	coll := r.Client.Database(storage.MongoDBDatabase).Collection(storage.MongoDBCollection)

	_, err := coll.InsertOne(context.Background(), MessageInfo{UserId: userId, Message: message})
	if err != nil {
		log.Print(err)
	}
}

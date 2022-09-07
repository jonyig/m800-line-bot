package repository

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"m800-line-bot/models"
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

	coll := r.GetMessageCollection()

	_, err := coll.InsertOne(context.Background(), models.MessageInfo{UserId: userId, Message: message})
	if err != nil {
		log.Print(err)
	}
}
func (r *LineBotMongoRepository) GetMessages() (results []models.MessageInfo, err error) {
	coll := r.GetMessageCollection()
	cursor, err := coll.Find(context.TODO(), bson.D{})
	if err != nil {
		return
	}
	if err = cursor.All(context.Background(), &results); err != nil {
		return
	}

	return
}

func (r *LineBotMongoRepository) GetMessageCollection() *mongo.Collection {
	return r.Client.Database(storage.MongoDBDatabase).Collection(storage.MongoDBCollection)
}

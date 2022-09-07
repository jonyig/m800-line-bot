package repository

import "go.mongodb.org/mongo-driver/mongo"

type LineBotMongoRepository struct {
	Client *mongo.Client
}

func NewLineBotMongoRepository(client *mongo.Client) *LineBotMongoRepository {
	return &LineBotMongoRepository{
		Client: client,
	}
}

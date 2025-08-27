package mongodb

import (
	"context"

	"github.com/mo-lab/jobv/api/v2/internal/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func GetUserCollection() (*mongo.Collection, error) {
	client, err := ConnectToClient()
	if err != nil {
		return nil, err
	}
	return client.Database("testdb").Collection("users"), nil
}
func CheckUserPhone(phone string) bool {
	collection, err := GetUserCollection()
	if err != nil {
		return false
	}

	var user models.User
	err = collection.FindOne(context.TODO(), bson.M{"phone": phone}).Decode(&user)
	if err != nil {
		return false

	}
	return true
}

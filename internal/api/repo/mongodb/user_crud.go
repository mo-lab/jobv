package mongodb

import (
	"context"

	"github.com/mo-lab/jobv/api/v2/internal/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func CreateUser(user models.User) (string, error) {
	collection, err := GetUserCollection()
	if err != nil {
		return "", err
	}
	res, err := collection.InsertOne(context.TODO(), user)
	if err != nil {
		return "", err
	}
	id := res.InsertedID.(primitive.ObjectID).Hex()
	return id, nil
}

func CreateUsers(users []models.User) ([]string, error) {
	collection, err := GetUserCollection()
	if err != nil {
		return nil, err
	}
	var userIDs []string
	for _, user := range users {
		res, err := collection.InsertOne(context.TODO(), user)
		if err != nil {
			return nil, err
		}
		id := res.InsertedID.(primitive.ObjectID).Hex()
		userIDs = append(userIDs, id)
	}
	return userIDs, nil
}

func GetUserByID(id string) (models.User, error) {
	collection, err := GetUserCollection()
	if err != nil {
		return models.User{}, err
	}
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return models.User{}, err
	}
	var user models.User
	err = collection.FindOne(context.TODO(), bson.M{"_id": objID}).Decode(&user)
	if err != nil {
		return models.User{}, err

	}
	return user, nil
}

func GetUserByPhone(phone string) (models.User, error) {
	collection, err := GetUserCollection()
	if err != nil {
		return models.User{}, err
	}

	var user models.User
	err = collection.FindOne(context.TODO(), bson.M{"phone": phone}).Decode(&user)
	if err != nil {
		return models.User{}, err

	}
	return user, nil
}

func UpdateUser(id string, user models.User) error {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}
	collection, err := GetUserCollection()
	if err != nil {
		return err
	}
	_, err = collection.UpdateOne(context.TODO(),
		bson.M{"_id": objID},
		bson.M{"$set": user},
	)
	if err != nil {
		return err
	}
	return nil
}
func DeleteUser(id string) error {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}
	collection, err := GetUserCollection()
	if err != nil {
		return err
	}
	_, err = collection.DeleteOne(context.TODO(), bson.M{"_id": objID})
	if err != nil {
		return err
	}
	return nil
}
func GetUsers(ctx context.Context, filter bson.M, page int, limit int) ([]models.User, error) {
	collection, err := GetUserCollection()
	if err != nil {
		return nil, err
	}
	skip := (page - 1) * limit
	// Set options for pagination
	findOptions := options.Find()
	findOptions.SetSkip(int64(skip))
	findOptions.SetLimit(int64(limit))
	cursor, err := collection.Find(ctx, filter, findOptions)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var users []models.User
	for cursor.Next(ctx) {
		var user models.User
		if err := cursor.Decode(&user); err == nil {
			users = append(users, user)
		}
	}
	return users, nil
}

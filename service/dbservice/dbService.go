package dbservice

import (
	"context"
	"errors"
	"log"
	"os"

	"github.com/emre-guler/question-answer/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var applyUri string = os.Getenv("DB_APLLY_URI")

func InsertUser(userData *models.User) bool {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(applyUri))
	if err != nil {
		log.Println("Database connection failed: ", err)
	}
	userCollection := client.Database("question-answer").Collection("users")
	_, err = userCollection.InsertOne(context.TODO(), userData)

	if err != nil {
		log.Println("Insert failed: ", err)
		return false
	}
	return true
}

func IsUserExist(githubId int) (bool, error) {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(applyUri))
	if err != nil {
		log.Println("Database connection failed: ", err)
	}
	userCollection := client.Database("question-answer").Collection("users")
	filter := bson.D{{Key: "githubid", Value: githubId}}
	var result bson.D
	err = userCollection.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		log.Println("Select failed: ", err)
		return true, errors.New("select process faield")
	}
	if result != nil {
		return true, nil
	}
	return false, nil
}

func UpdateUserAccessToken(newUser *models.User) bool {
	// TODO Update token
	return true
}

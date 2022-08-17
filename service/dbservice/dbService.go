package dbservice

import (
	"context"
	"log"
	"os"

	"github.com/emre-guler/question-answer/models"
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

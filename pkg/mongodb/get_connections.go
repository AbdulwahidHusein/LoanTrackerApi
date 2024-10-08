package mongodb

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Client *mongo.Client

var UserCollection *mongo.Collection
var TokenCollection *mongo.Collection
var LoanCollection *mongo.Collection

// ConnectDB connects to the MongoDB database
func ConnectDB(uri string) (*mongo.Client, error) {
	clientOptions := options.Client().ApplyURI(uri)
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Connected to MongoDB!")

	Client = client
	return client, err
}

// GetCollection returns a collection from the connected MongoDB database
func GetCollection(collectionName string) *mongo.Collection {

	return Client.Database("loanDb").Collection(collectionName)
}

// DisconnectDB disconnects from the MongoDB database
func DisconnectDB() {
	err := Client.Disconnect(context.TODO())
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Disconnected from MongoDB")
}

func InitializeCollections() {
	UserCollection = GetCollection("users")
	TokenCollection = GetCollection("tokens")
	LoanCollection = GetCollection("loans")
}

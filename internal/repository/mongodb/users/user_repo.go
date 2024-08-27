package users

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type MongoUserRepository struct {
	userCollection *mongo.Collection
}

func NewMongoUserRepository(userCollection *mongo.Collection) *MongoUserRepository {
	return &MongoUserRepository{
		userCollection: userCollection,
	}
}

func (userRepo *MongoUserRepository) Create(user *entity.User) error {
	_, err := userRepo.userCollection.InsertOne(nil, user)

	return err
}

func (userRepo *MongoUserRepository) FindByEmail(email string) (*entity.User, error) {
	var user entity.User
	err := userRepo.userCollection.FindOne(nil, bson.M{"email": email}).Decode(&user)
	return &user, err
}

func (userRepo *MongoUserRepository) FindByID(id string) (*entity.User, error) {
	var user entity.User
	err := userRepo.userCollection.FindOne(nil, bson.M{"id": id}).Decode(&user)
	return &user, err
}

func (userRepo *MongoUserRepository) Update(user *entity.User) error {
	_, err := userRepo.userCollection.UpdateOne(nil, bson.M{"id": user.ID}, bson.M{"$set": user})

	return err
}

func (userRepo *MongoUserRepository) GetAllUsers() ([]*entity.User, error) {
	var users []*entity.User
	cursor, err := userRepo.userCollection.Find(nil, bson.M{})
	if err != nil {
		return nil, err
	}
	for cursor.Next(context.Background()) {
		var user entity.User
		err := cursor.Decode(&user)
		if err != nil {
			return nil, err
		}
		users = append(users, &user)
	}
	return users, nil
}

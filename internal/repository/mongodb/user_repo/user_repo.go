package user_repo

import (
	"LoanTrackerApi/internal/entity"
	"LoanTrackerApi/internal/repository"
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoUserRepository struct {
	userCollection *mongo.Collection
}

// NewMongoUserRepository creates a new instance of MongoUserRepository.
func NewMongoUserRepository(userCollection *mongo.Collection) repository.UserRepository {
	return &MongoUserRepository{
		userCollection: userCollection,
	}
}

// Create inserts a new user into the user collection.
func (userRepo *MongoUserRepository) Create(ctx context.Context, user *entity.User) error {
	_, err := userRepo.userCollection.InsertOne(ctx, user)
	if err != nil {
		return err
	}
	return nil
}

// FindByEmail retrieves a user by their email address and returns a GetUserDTO.
func (userRepo *MongoUserRepository) FindByEmail(ctx context.Context, email string) (*entity.User, error) {
	var user entity.User
	err := userRepo.userCollection.FindOne(ctx, bson.M{"email": email}).Decode(&user)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil // Return nil if no user is found
		}
		return nil, err
	}

	return &user, nil
}

// FindByID retrieves a user by their ID and returns a GetUserDTO.
func (userRepo *MongoUserRepository) FindByID(ctx context.Context, id string) (*entity.GetUserDTO, error) {
	var user entity.User
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	err = userRepo.userCollection.FindOne(ctx, bson.M{"_id": objectId}).Decode(&user)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil // Return nil if no user is found
		}
		return nil, err
	}

	userDTO := mapUserToGetUserDTO(&user)
	return userDTO, nil
}

// Update modifies an existing user's information.
func (userRepo *MongoUserRepository) Update(ctx context.Context, userDTO *entity.UpdateUserDTO) error {
	filter := bson.M{"email": userDTO.Email} // Assuming email is used as an identifier
	update := bson.M{"$set": userDTO}
	opts := options.Update().SetUpsert(false) // Upsert=false means update only, no insert
	_, err := userRepo.userCollection.UpdateOne(ctx, filter, update, opts)
	if err != nil {
		return err
	}
	return nil
}

// GetAllUsers retrieves all users from the user collection and returns a slice of GetUserDTOs.
func (userRepo *MongoUserRepository) GetAllUsers(ctx context.Context, page, pageSize int) ([]*entity.GetUserDTO, error) {
	var users []*entity.GetUserDTO

	// Calculate the number of documents to skip
	skip := (page - 1) * pageSize

	// Find options with pagination
	findOptions := options.Find()
	findOptions.SetSkip(int64(skip))
	findOptions.SetLimit(int64(pageSize))

	// Execute the query with pagination
	cursor, err := userRepo.userCollection.Find(ctx, bson.M{}, findOptions)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var user entity.GetUserDTO
		if err := cursor.Decode(&user); err != nil {
			return nil, err
		}
		users = append(users, &user)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return users, nil
}

func (userRepo *MongoUserRepository) Delete(ctx context.Context, id string) error {
	objId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}
	_, err = userRepo.userCollection.DeleteOne(ctx, bson.M{"_id": objId})
	if err != nil {
		return err
	}
	return nil
}

func (userRepo *MongoUserRepository) Verify(ctx context.Context, id string) error {
	_, err := userRepo.userCollection.UpdateOne(ctx, bson.M{"_id": id}, bson.M{"$set": bson.M{"verified": true}})
	if err != nil {
		return err
	}
	return nil
}

func (userRepo *MongoUserRepository) ChangePassword(ctx context.Context, user *entity.User, password string) error {
	_, err := userRepo.userCollection.UpdateOne(ctx, bson.M{"email": user.Email}, bson.M{"$set": bson.M{"password": password}})
	if err != nil {
		return err
	}
	return nil
}

// Helper function to map a User entity to a GetUserDTO
func mapUserToGetUserDTO(user *entity.User) *entity.GetUserDTO {
	return &entity.GetUserDTO{
		ID:       user.ID,
		UserName: user.Profile.FirstName + " " + user.Profile.LastName,
		Email:    user.Email,
		Profile:  user.Profile,
		Role:     user.Role,
		Created:  user.Created,
		Updated:  user.Updated,
		Verified: user.Verified,
	}
}

// Helper function to map a list of User entities to a list of GetUserDTOs
// func mapUsersToGetUserDTOs(users []*entity.User) []*entity.GetUserDTO {
// 	var userDTOs []*entity.GetUserDTO
// 	for _, user := range users {
// 		userDTO := mapUserToGetUserDTO(user)
// 		userDTOs = append(userDTOs, userDTO)
// 	}
// 	return userDTOs
// }

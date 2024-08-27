package loan_repo

import (
	"LoanTrackerApi/internal/entity"
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoLoanRepository struct {
	loanCollection *mongo.Collection
}

func NewMongoLoanRepository(loanCollection *mongo.Collection) *MongoLoanRepository {
	return &MongoLoanRepository{
		loanCollection: loanCollection,
	}
}

func (r *MongoLoanRepository) Create(ctx context.Context, loan *entity.Loan) error {
	loan.ID = primitive.NewObjectID()
	_, err := r.loanCollection.InsertOne(ctx, loan)
	return err
}

// Find a loan by its ID
func (r *MongoLoanRepository) FindByID(ctx context.Context, id string) (*entity.Loan, error) {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, fmt.Errorf("invalid id: %w", err)
	}

	var loan entity.Loan
	err = r.loanCollection.FindOne(ctx, bson.M{"_id": objectID}).Decode(&loan)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}
		return nil, err
	}

	return &loan, nil
}

// Find loans by the user's ID
func (r *MongoLoanRepository) FindByUserID(ctx context.Context, id string) ([]*entity.Loan, error) {
	userID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, fmt.Errorf("invalid user id: %w", err)
	}

	cursor, err := r.loanCollection.Find(ctx, bson.M{"issuer_id": userID})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var loans []*entity.Loan
	for cursor.Next(ctx) {
		var loan entity.Loan
		if err := cursor.Decode(&loan); err != nil {
			return nil, err
		}
		loans = append(loans, &loan)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return loans, nil
}

// Update an existing loan
func (r *MongoLoanRepository) Update(ctx context.Context, loan *entity.Loan) error {
	filter := bson.M{"_id": loan.ID}
	update := bson.M{"$set": loan}

	_, err := r.loanCollection.UpdateOne(ctx, filter, update)
	return err
}

// Delete a loan by its ID
func (r *MongoLoanRepository) Delete(ctx context.Context, id string) error {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return fmt.Errorf("invalid id: %w", err)
	}

	_, err = r.loanCollection.DeleteOne(ctx, bson.M{"_id": objectID})
	return err
}

// Get paginated, filtered, and sorted loans
func (r *MongoLoanRepository) GetLoans(ctx context.Context, page, pageSize int, filter entity.LoanFilter) ([]*entity.Loan, error) {
	skip := (page - 1) * pageSize
	findOptions := options.Find()
	findOptions.SetSkip(int64(skip))
	findOptions.SetLimit(int64(pageSize))
	findOptions.SetSort(parseSortOptions(filter.OrderBy))

	filterBson := bson.M{}
	if filter.Status != "" {
		filterBson["status"] = filter.Status
	}

	cursor, err := r.loanCollection.Find(ctx, filterBson, findOptions)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var loans []*entity.Loan
	for cursor.Next(ctx) {
		var loan entity.Loan
		if err := cursor.Decode(&loan); err != nil {
			return nil, err
		}
		loans = append(loans, &loan)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return loans, nil
}

// Helper function to parse sort options from a slice of strings
func parseSortOptions(orderBy []string) bson.D {
	var sortOptions bson.D
	for _, order := range orderBy {
		switch order {
		case "date_asc":
			sortOptions = append(sortOptions, bson.E{Key: "date", Value: 1})
		case "date_desc":
			sortOptions = append(sortOptions, bson.E{Key: "date", Value: -1})
		case "amt_asc":
			sortOptions = append(sortOptions, bson.E{Key: "amount", Value: 1})
		case "amt_desc":
			sortOptions = append(sortOptions, bson.E{Key: "amount", Value: -1})
		default:
			// Default sorting
			sortOptions = append(sortOptions, bson.E{Key: "date", Value: 1})
		}
	}
	return sortOptions
}

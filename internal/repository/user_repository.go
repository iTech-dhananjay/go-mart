package repository

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/yourusername/yourproject/internal/models"
)

type UserRepository struct {
	client *mongo.Client
}

func NewUserRepository(ctx context.Context, dbURI, dbName string) (*UserRepository, error) {
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(dbURI))
	if err != nil {
		return nil, err
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		return nil, err
	}

	fmt.Println("Connected to MongoDB!")

	return &UserRepository{
		client: client,
	}, nil
}

func (ur *UserRepository) GetAllUsers(ctx context.Context) ([]models.User, error) {
	var users []models.User
	collection := ur.client.Database("yourdatabase").Collection("users")
	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)
	for cursor.Next(ctx) {
		var user models.User
		if err := cursor.Decode(&user); err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}

package repositories

import (
	"context"
	. "go-mongo-redis/domain/datasources"
	"go-mongo-redis/domain/entities"
	"os"

	fiberlog "github.com/gofiber/fiber/v2/log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type usersRepository struct {
	Context    context.Context
	Collection *mongo.Collection
}

type IUsersRepository interface {
	FindAll() ([]entities.UserDataFormat, error)
}

func NewUsersRepository(db *MongoDB) IUsersRepository {
	return &usersRepository{
		Context:    db.Context,
		Collection: db.MongoDB.Database(os.Getenv("DATABASE_NAME")).Collection("users_mock"),
	}
}

func (repo usersRepository) FindAll() ([]entities.UserDataFormat, error) {
	cursor, err := repo.Collection.Find(repo.Context, bson.M{}, options.Find())
	if err != nil {
		fiberlog.Errorf("Users -> FindAll: %s \n", err)
		return nil, err
	}
	defer cursor.Close(repo.Context)

	var users []entities.UserDataFormat
	if err := cursor.All(repo.Context, &users); err != nil {
		fiberlog.Errorf("Users -> FindAll: %s \n", err)
		return nil, err
	}
	return users, nil
}

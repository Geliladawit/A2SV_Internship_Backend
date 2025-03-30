package Repositories

import (
	"context"
	"errors"
	"fmt"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"task_manager/Domain"
)

const (
	userCollectionName = "users"
)

var (
	userCollection *mongo.Collection
)

func init() {
	mongoURI := os.Getenv("MONGODB_URI")
	if mongoURI == "" {
		mongoURI = "mongodb://localhost:27017"
	}

	clientOptions := options.Client().ApplyURI(mongoURI)
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB")

	userCollection = client.Database(dbName).Collection(userCollectionName)
}

type UserRepository interface {
	Create(user Domain.User) (*Domain.User, error)
	GetByUsername(username string) (*Domain.User, error)
	GetByID(id primitive.ObjectID) (*Domain.User, error)
	Promote(id primitive.ObjectID) error
}

type UserRepositoryImpl struct{}

func NewUserRepository() UserRepository {
	return &UserRepositoryImpl{}
}

func (r *UserRepositoryImpl) Create(user Domain.User) (*Domain.User, error) {
	count, err := userCollection.CountDocuments(context.TODO(), bson.D{})
	if err != nil {
		return nil, err
	}

	if count == 0 {
		user.IsAdmin = true
	}

	insertResult, err := userCollection.InsertOne(context.TODO(), user)
	if err != nil {
		return nil, err
	}

	objID := insertResult.InsertedID.(primitive.ObjectID)
	user.ID = objID
	return &user, nil
}

func (r *UserRepositoryImpl) GetByUsername(username string) (*Domain.User, error) {
	filter := bson.M{"username": username}
	var user Domain.User

	err := userCollection.FindOne(context.TODO(), filter).Decode(&user)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, errors.New("user not found")
		}
		return nil, err
	}
	return &user, nil
}

func (r *UserRepositoryImpl) Promote(id primitive.ObjectID) error {
	filter := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"is_admin": true}}

	updateResult, err := userCollection.UpdateOne(context.TODO(), filter, update)

	if err != nil {
		return err
	}

	if updateResult.MatchedCount == 0 {
		return errors.New("user not found")
	}

	return nil
}

func (r *UserRepositoryImpl) GetByID(id primitive.ObjectID) (*Domain.User, error) {
	filter := bson.M{"_id": id}
	var user Domain.User

	err := userCollection.FindOne(context.TODO(), filter).Decode(&user)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, errors.New("user not found")
		}
		return nil, err
	}
	return &user, nil
}

package Repositories

import (
	"errors"
	"context"
	"log"
	"os"

	"task_manager/Domain"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang.org/x/crypto/bcrypt"
)

type UserRepository interface {
	Create(ctx context.Context, user *Domain.User) (*Domain.User, error)
	GetByUsername(ctx context.Context, username string) (*Domain.User, error)
	GetByID(ctx context.Context, id primitive.ObjectID) (*Domain.User, error)
	PromoteUser(ctx context.Context, id primitive.ObjectID) error
}

const (
	userCollectionName = "users"
)

type UserRepositoryImpl struct {
	client         *mongo.Client
	userCollection *mongo.Collection
}

func NewUserRepository() (*UserRepositoryImpl, error) {
	mongoURI := os.Getenv("MONGODB_URI")
	if mongoURI == "" {
		mongoURI = "mongodb://localhost:27017"
	}

	clientOptions := options.Client().ApplyURI(mongoURI)
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	userCollection := client.Database(dbName).Collection(userCollectionName)

	return &UserRepositoryImpl{client: client, userCollection: userCollection}, nil
}

func (r *UserRepositoryImpl) Create(ctx context.Context, user *Domain.User) (*Domain.User, error) {
	hashedPassword, err := HashPassword(user.Password)
	if err != nil {
		return nil, err
	}
	user.Password = hashedPassword

	count, err := r.userCollection.CountDocuments(ctx, bson.D{})
	if err != nil {
		return nil, err
	}

	if count == 0 {
		user.IsAdmin = true
	}

	insertResult, err := r.userCollection.InsertOne(ctx, user)
	if err != nil {
		return nil, err
	}

	objID := insertResult.InsertedID.(primitive.ObjectID)
	user.ID = objID
	return user, nil
}

func (r *UserRepositoryImpl) GetByUsername(ctx context.Context, username string) (*Domain.User, error) {
	filter := bson.M{"username": username}
	var user Domain.User

	err := r.userCollection.FindOne(ctx, filter).Decode(&user)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, errors.New("user not found")
		}
		return nil, err
	}
	return &user, nil
}

func (r *UserRepositoryImpl) GetByID(ctx context.Context, id primitive.ObjectID) (*Domain.User, error) {
	filter := bson.M{"_id": id}
	var user Domain.User

	err := r.userCollection.FindOne(ctx, filter).Decode(&user)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, errors.New("user not found")
		}
		return nil, err
	}
	return &user, nil
}

func (r *UserRepositoryImpl) PromoteUser(ctx context.Context, id primitive.ObjectID) error {
	filter := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"is_admin": true}}

	updateResult, err := r.userCollection.UpdateOne(ctx, filter, update)

	if err != nil {
		return err
	}

	if updateResult.MatchedCount == 0 {
		return errors.New("user not found")
	}

	return nil
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}
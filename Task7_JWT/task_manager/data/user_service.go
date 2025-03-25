package data

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
	"golang.org/x/crypto/bcrypt"
	"task_manager/models"
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

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func CreateUser(user models.User) (*models.User, error) {
	hashedPassword, err := HashPassword(user.Password)
	if err != nil {
		return nil, err
	}
	user.Password = hashedPassword

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

func GetUserByUsername(username string) (*models.User, error) {
	filter := bson.M{"username": username}
	var user models.User

	err := userCollection.FindOne(context.TODO(), filter).Decode(&user)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, errors.New("user not found")
		}
		return nil, err
	}
	return &user, nil
}

func PromoteUser(id string) error {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return errors.New("invalid user ID format")
	}
	filter := bson.M{"_id": objID}
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

func GetUserByID(id primitive.ObjectID) (*models.User, error) {
	filter := bson.M{"_id": id}
	var user models.User

	err := userCollection.FindOne(context.TODO(), filter).Decode(&user)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, errors.New("user not found")
		}
		return nil, err
	}
	return &user, nil
}
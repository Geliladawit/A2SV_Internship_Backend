package data

import (
	"context"
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"task_manager/models"
)

const (
	dbName = "task_manager_db"
	collectionName = "tasks"
)

var (
	mongoClient *mongo.Client
	taskCollection *mongo.Collection
)

func init () {
	mongoURI := os.Getenv("MONGODB_URI")
	if mongoURI == ""{
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
	mongoClient = client

	taskCollection = client.Database(dbName).Collection(collectionName)
}

func GetAllTasks()([]models.Task, error){
	filter := bson.D{{}}
	cursor, err := taskCollection.Find(context.TODO(), filter)
	if err != nil {
		return nil, err
}

var tasks []models.Task
if err = cursor.All(context.TODO(), &tasks); err != nil{
	return nil, err
}
return tasks, nil
}

func GetTaskByID(id string) (*models.Task, error){
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, errors.New("invalid task ID format")
	}
	filter := bson.M{"_id": objID}
	var task models.Task

	err = taskCollection.FindOne(context.TODO(), filter).Decode(&task)
	if err != nil{
		if errors.Is(err, mongo.ErrNoDocuments){
			return nil, errors.New("task not found")
		}
		return nil, err
	}
	return &task, nil
	}

	func CreateTask(task models.Task) (*models.Task, error){
		insertResult, err := taskCollection.InsertOne(context.TODO(), task)
		if err != nil {
			return nil, err
	}

	objID := insertResult.InsertedID.(primitive.ObjectID)
	task.ID = objID

	return &task, nil
}

func UpdateTask(id string, updatedTask models.Task) (*models.Task, error){
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, errors.New("invalid task ID format")
	}
	filter := bson.M{"_id":objID}
	update := bson.M{"$set": updatedTask}

	updateResult, err := taskCollection.UpdateOne(context.TODO(), filter, update)

	if err != nil{
		return nil, err
	}
	if updateResult.MatchedCount == 0{
		return nil, errors.New("task not found")
	}
	updateTaskResult, err := GetTaskByID(id)
	if err != nil{
		return nil, err
	}
	return updateTaskResult, nil
}

func DeleteTask(id string) error{
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return errors.New("invalid task ID format")
	}
	filter := bson.M{"_id":objID}
	deleteResult, err := taskCollection.DeleteOne(context.TODO(), filter)
	if err != nil {
		return err
	}
	if deleteResult.DeletedCount==0{
		return errors.New("task not found")
	}
	return nil
}

func ParseTime(timeStr string) (time.Time, error){
	parsedTime, err := time.Parse(time.RFC3339, timeStr)

	if err != nil{
		return time.Time{}, errors.New("error parsing time"+ err.Error())
	}
	return parsedTime, nil
	}
	
func ParseInt(intStr string) (int, error){
	parsedInt, err := strconv.Atoi(intStr)
	if err != nil{
		return -1, errors.New("error parsing time" + err.Error())
	}
	return parsedInt, nil
}
	func InitMemoryData(){
}
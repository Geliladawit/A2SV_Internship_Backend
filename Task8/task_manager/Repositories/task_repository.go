package Repositories

import (
	"context"
	"errors"
	"log"
	"os"

	"task_manager/Domain"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type TaskRepository interface {
	GetAll(ctx context.Context) ([]Domain.Task, error)
	GetByID(ctx context.Context, id primitive.ObjectID) (*Domain.Task, error)
	Create(ctx context.Context, task *Domain.Task) (*Domain.Task, error)
	Update(ctx context.Context, id primitive.ObjectID, task *Domain.Task) (*Domain.Task, error)
	Delete(ctx context.Context, id primitive.ObjectID) error
}

const (
	dbName         = "task_manager_db"
	taskCollectionName = "tasks"
)

type TaskRepositoryImpl struct {
	client         *mongo.Client
	taskCollection *mongo.Collection
}

func NewTaskRepository() (*TaskRepositoryImpl, error) {
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

	taskCollection := client.Database(dbName).Collection(taskCollectionName)

	return &TaskRepositoryImpl{client: client, taskCollection: taskCollection}, nil
}

func (r *TaskRepositoryImpl) GetAll(ctx context.Context) ([]Domain.Task, error) {
	filter := bson.D{{}}
	cursor, err := r.taskCollection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}

	var tasks []Domain.Task
	if err = cursor.All(ctx, &tasks); err != nil {
		return nil, err
	}
	return tasks, nil
}

func (r *TaskRepositoryImpl) GetByID(ctx context.Context, id primitive.ObjectID) (*Domain.Task, error) {
	filter := bson.M{"_id": id}
	var task Domain.Task

	err := r.taskCollection.FindOne(ctx, filter).Decode(&task)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, errors.New("task not found")
		}
		return nil, err
	}
	return &task, nil
}

func (r *TaskRepositoryImpl) Create(ctx context.Context, task *Domain.Task) (*Domain.Task, error) {
	insertResult, err := r.taskCollection.InsertOne(ctx, task)
	if err != nil {
		return nil, err
	}

	objID := insertResult.InsertedID.(primitive.ObjectID)
	task.ID = objID

	return task, nil
}

func (r *TaskRepositoryImpl) Update(ctx context.Context, id primitive.ObjectID, updatedTask *Domain.Task) (*Domain.Task, error) {
	filter := bson.M{"_id": id}
	update := bson.M{"$set": updatedTask}

	updateResult, err := r.taskCollection.UpdateOne(ctx, filter, update)

	if err != nil {
		return nil, err
	}
	if updateResult.MatchedCount == 0 {
		return nil, errors.New("task not found")
	}

	updatedTask.ID = id 
	return updatedTask, nil
}

func (r *TaskRepositoryImpl) Delete(ctx context.Context, id primitive.ObjectID) error {
	filter := bson.M{"_id": id}
	deleteResult, err := r.taskCollection.DeleteOne(ctx, filter)
	if err != nil {
		return err
	}
	if deleteResult.DeletedCount == 0 {
		return errors.New("task not found")
	}
	return nil
}
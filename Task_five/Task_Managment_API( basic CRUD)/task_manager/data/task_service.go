package data

import (
	"errors"
	"strconv"
	"task_manager/models"
	"time"
)

var (
	tasks  = []*models.Task{}
	nextID = 1
)

func GetAllTasks() []*models.Task {
	return tasks
}

func GetTask(id int) (*models.Task, error) {
	for i, task := range tasks {
		if task.ID == id {
			return tasks[i], nil
		}
	}
	return nil, errors.New("task not found")
}

func CreateTask(task models.Task) (*models.Task, error) {
	task.ID = nextID
	nextID++
	tasks = append(tasks, &task)
	return tasks[len(tasks)-1], nil
}

func UpdateTask(id int, updatedTask models.Task) (*models.Task, error) {
	for i, task := range tasks {
		if task.ID == id {
			updatedTask.ID = id
			tasks[i] = &updatedTask
			return &updatedTask, nil
		}
	}
	return nil, errors.New("task not found")
}

func DeleteTask(id int) error {
	for i, task := range tasks {
		if task.ID == id {
			tasks = append(tasks[:i], tasks[i+1:]...)
			return nil
		}
	}
	return errors.New("task not found")
}

func ParseTime(timeStr string) (time.Time, error) {
	parsedTime, err := time.Parse(time.RFC3339, timeStr)
	if err != nil {
		return time.Time{}, errors.New("error parsing time" + err.Error())
	}

	return parsedTime, nil
}

func ParseInt(intStr string) (int, error) {
	parsedInt, err := strconv.Atoi(intStr)
	if err != nil {
		return -1, errors.New("error parsing time" + err.Error())
	}

	return parsedInt, nil
}

func InitMemoryData() {
	time1, _ := ParseTime("2024-12-20T00:00:00Z")
	time2, _ := ParseTime("2024-12-22T00:00:00Z")
	task1 := models.Task{ID: 1, Title: "Do laundry", Description: "Clean all clothes", DueDate: time1, Status: "In Progress"}
	task2 := models.Task{ID: 2, Title: "Buy grocery", Description: "Buy daily products", DueDate: time2, Status: "Completed"}
	tasks = append(tasks, &task1, &task2)
	nextID = 3
}

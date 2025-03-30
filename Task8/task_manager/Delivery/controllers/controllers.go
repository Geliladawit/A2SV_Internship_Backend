package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"task_manager/Domain"
	"task_manager/Infrastructure"
	"task_manager/Usecases"
)

type UserController struct {
	userUsecase Usecases.UserUsecase
	jwtService  Infrastructure.JWTService
}

func NewUserController(userUsecase Usecases.UserUsecase, jwtService Infrastructure.JWTService) *UserController {
	return &UserController{userUsecase: userUsecase, jwtService: jwtService}
}

func (uc *UserController) RegisterUser(c *gin.Context) {
	var user Domain.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	createdUser, err := uc.userUsecase.Register(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create user", "details": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, createdUser)
}

func (uc *UserController) LoginUser(c *gin.Context) {
	var loginRequest Domain.LoginRequest
	if err := c.ShouldBindJSON(&loginRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := uc.userUsecase.Login(loginRequest.Username, loginRequest.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	token, err := uc.jwtService.GenerateToken(*user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}

	c.JSON(http.StatusOK, Domain.LoginResponse{Token: token})
}

func (uc *UserController) PromoteUser(c *gin.Context) {
	id := c.Param("id")
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid user ID format"})
		return
	}

	err = uc.userUsecase.Promote(objID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to promote user", "details": err.Error()})
		return
	}
	c.Status(http.StatusOK)
}

type TaskController struct {
	taskUsecase Usecases.TaskUsecase
}

func NewTaskController(taskUsecase Usecases.TaskUsecase) *TaskController {
	return &TaskController{taskUsecase: taskUsecase}
}

func (tc *TaskController) GetTasks(c *gin.Context) {
	tasks, err := tc.taskUsecase.GetAllTasks()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to retrieve tasks", "details": err.Error()})
		return
	}
	c.JSON(http.StatusOK, tasks)
}

func (tc *TaskController) GetTask(c *gin.Context) {
	id := c.Param("id")
	task, err := tc.taskUsecase.GetTaskByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "task not found"})
		return
	}
	c.JSON(http.StatusOK, task)
}

func (tc *TaskController) CreateTask(c *gin.Context) {
	var task Domain.Task

	if err := c.ShouldBindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	createdTask, err := tc.taskUsecase.CreateTask(task)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create task", "details": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, createdTask)
}

func (tc *TaskController) UpdateTask(c *gin.Context) {
	id := c.Param("id")
	var updatedtask Domain.Task
	if err := c.ShouldBindJSON(&updatedtask); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	updatedTaskResult, err := tc.taskUsecase.UpdateTask(id, updatedtask)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "task not found"})
		return
	}
	c.JSON(http.StatusOK, updatedTaskResult)
}

func (tc *TaskController) DeleteTask(c *gin.Context) {
	id := c.Param("id")
	err := tc.taskUsecase.DeleteTask(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "task not found"})
		return
	}
	c.Status(http.StatusNoContent)
}

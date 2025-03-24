package controllers

import (
	"net/http"
	"strconv"
	"github.com/gin-gonic/gin"
	"task_manager/data"
	"task_manager/models"

)
func GetTasks(c *gin.Context){
	tasks := data.GetAllTasks()
	c.JSON(http.StatusOK, tasks)
}

func GetTask(c *gin.Context){
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid task ID"+err.Error()})
		return
	}
	task, err := data.GetTask(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, task)
}

func CreateTask(c *gin.Context){
	var task models.Task

	if err := c.ShouldBindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid task data"+err.Error()})
		return
	}
	createdTask, err := data.CreateTask(task)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error":"failed to create task"})
		return
	}
	c.JSON(http.StatusCreated, createdTask)
}

func UpdateTask(c *gin.Context){
	idStr := c.Param("id")
	id, err := data.ParseInt(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid task ID"+err.Error()})
		return
	}
	var updateTask models.Task
	if err := c.ShouldBindJSON(&updateTask); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid task data"+err.Error()})
		return
	}
	updatedTaskResult, err := data.UpdateTask(id, updateTask)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error":"failed to update task"})
		return
	}
	c.JSON(http.StatusOK, 	updatedTaskResult)
}

func DeleteTask(c *gin.Context){
	idStr := c.Param("id")
	id, err := data.ParseInt(idStr)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid tasl ID"+ err.Error()})
		return 
	}
	err = data.DeleteTask(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
}
c.Status(http.StatusNoContent)
}

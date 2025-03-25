package router
import (
	"github.com/gin-gonic/gin"
	"task_manager/controllers"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	
	r.GET("/tasks", controllers.GetTasks)
	r.POST("/tasks", controllers.CreateTask)
	r.GET("/tasks/:id", controllers.GetTask)
	r.PUT("/tasks/id",controllers.UpdateTask)
	r.DELETE("/tasks/:id", controllers.DeleteTask)
	return r

}

func InitData(){

}
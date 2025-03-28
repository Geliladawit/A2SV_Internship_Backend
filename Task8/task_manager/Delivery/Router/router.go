package router

import (
	"github.com/gin-gonic/gin"
	"task_manager/Delivery/Controllers"
	"task_manager/Infrastructure"
)

func SetupRouter(taskController *controllers.TaskController, userController *controllers.UserController, jwtService Infrastructure.JWTService) *gin.Engine {
	r := gin.Default()

	r.POST("/register", userController.RegisterUser)
	r.POST("/login", userController.LoginUser)

	tasks := r.Group("/tasks")
	tasks.Use(Infrastructure.Authenticate(jwtService))
	{
		tasks.GET("", taskController.GetTasks)
		tasks.GET("/:id", taskController.GetTask)
		tasks.POST("", Infrastructure.AuthorizeAdmin, taskController.CreateTask)
		tasks.PUT("/:id", Infrastructure.AuthorizeAdmin, taskController.UpdateTask)
		tasks.DELETE("/:id", Infrastructure.AuthorizeAdmin, taskController.DeleteTask)
	}
	admin := r.Group("/admin")
	admin.Use(Infrastructure.Authenticate(jwtService), Infrastructure.AuthorizeAdmin)
	{
		admin.PUT("/promote/:id", userController.PromoteUser)
	}
	return r

}
package routers

import (
	"github.com/gin-gonic/gin"
	"task_manager/Delivery/controllers"
	"task_manager/Infrastructure"
	"task_manager/Repositories"
	"task_manager/Usecases"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	jwtService := Infrastructure.NewJWTService()
	passwordService := Infrastructure.NewPasswordService()

	userRepository := Repositories.NewUserRepository()
	taskRepository := Repositories.NewTaskRepository()

	userUsecase := Usecases.NewUserUsecase(userRepository, passwordService, jwtService)
	taskUsecase := Usecases.NewTaskUsecase(taskRepository)

	userController := controllers.NewUserController(userUsecase, jwtService)
	taskController := controllers.NewTaskController(taskUsecase)

	r.POST("/register", userController.RegisterUser)
	r.POST("/login", userController.LoginUser)

	tasks := r.Group("/tasks")
	tasks.Use(Infrastructure.Authenticate(jwtService))
	{
		tasks.GET("", taskController.GetTasks)
		tasks.GET("/:id", taskController.GetTask)
		tasks.POST("", Infrastructure.AuthorizeAdmin(), taskController.CreateTask)
		tasks.PUT("/:id", Infrastructure.AuthorizeAdmin(), taskController.UpdateTask)
		tasks.DELETE("/:id", Infrastructure.AuthorizeAdmin(), taskController.DeleteTask)
	}

	admin := r.Group("/admin")
	admin.Use(Infrastructure.Authenticate(jwtService), Infrastructure.AuthorizeAdmin())
	{
		admin.PUT("/promote/:id", userController.PromoteUser)
	}

	return r
}


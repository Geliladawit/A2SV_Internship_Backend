package router

import (
	"github.com/gin-gonic/gin"
	"task_manager/controllers"
	"task_manager/middleware"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	r.POST("/register", controllers.RegisterUser)
	r.POST("/login", controllers.LoginUser)

	tasks := r.Group("/tasks")
	tasks.Use(middleware.Authenticate) 
	{
		tasks.GET("", controllers.GetTasks)        
		tasks.GET("/:id", controllers.GetTask)       
		tasks.POST("", middleware.AuthorizeAdmin, controllers.CreateTask)   
		tasks.PUT("/:id", middleware.AuthorizeAdmin, controllers.UpdateTask)  
		tasks.DELETE("/:id", middleware.AuthorizeAdmin, controllers.DeleteTask) 
	}
	admin := r.Group("/admin")
	admin.Use(middleware.Authenticate, middleware.AuthorizeAdmin)
	{
		admin.PUT("/promote/:id", controllers.PromoteUser) 
	}
	return r
}

func InitData() {

}
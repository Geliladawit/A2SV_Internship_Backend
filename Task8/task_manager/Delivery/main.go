package main

import (
	"log"
	"os"
	"task_manager/Delivery/Controllers"
	"task_manager/Repositories"
	"task_manager/Delivery/router"
	"task_manager/Usecases"
	"task_manager/Infrastructure"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	taskRepo, err := Repositories.NewTaskRepository()
	if err != nil {
		log.Fatalf("Failed to create task repository: %v", err)
	}
	userRepo, err := Repositories.NewUserRepository()
	if err != nil {
		log.Fatalf("Failed to create user repository: %v", err)
	}

	jwtService := Infrastructure.NewJWTService()
	passwordService := Infrastructure.NewPasswordService()

	taskUsecase := Usecases.NewTaskUsecase(taskRepo)
	userUsecase := Usecases.NewUserUsecase(userRepo, jwtService, passwordService)

	taskController := controllers.NewTaskController(taskUsecase)
	userController := controllers.NewUserController(userUsecase)

	r := router.SetupRouter(taskController, userController, jwtService)

	log.Printf("Server is running on port %s", port)
	err = r.Run(":" + port)
	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
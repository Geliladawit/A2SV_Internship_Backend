package main
import (
	"task_manager/router"
)

func main() {
	router.InitData()
	r := router.SetupRouter()

	r.Run(":8080")

}

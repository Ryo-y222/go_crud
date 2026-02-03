package main

import (
	"fmt"
	"go_crud/internal/controller"
	"go_crud/internal/repository"
	"go_crud/internal/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("Hello, World!")

	r := gin.Default()
	r.GET("/health", func(c *gin.Context) {
		c.String(http.StatusOK, "Health")
	})

	repo := repository.NewInMemoryTodoRepository()
	svc := service.NewTodoService(repo)
	todoCtl := controller.NewTodoController(svc)
	todoCtl.RegisterRoutes(r)

	r.Run(":3003")

}

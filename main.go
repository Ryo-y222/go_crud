package main

import (
	"database/sql"
	"fmt"
	"go_crud/internal/controller"
	"go_crud/internal/repository"
	"go_crud/internal/service"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	fmt.Println("Hello, World!")

	r := gin.Default()
	r.GET("/health", func(c *gin.Context) {
		c.String(http.StatusOK, "Health")
	})

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true&charset=utf8mb4&loc=Local",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASS"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
	)

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}

	repo := repository.NewMySQLTodoRepository(db)
	svc := service.NewTodoService(repo)
	todoCtl := controller.NewTodoController(svc)
	todoCtl.RegisterRoutes(r)

	r.Run(":3003")

}

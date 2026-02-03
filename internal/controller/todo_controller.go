package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"go_crud/internal/service"
)

type TodoController struct {
	svc *service.TodoService
}

func NewTodoController(svc *service.TodoService) *TodoController {
	return &TodoController{svc: svc}
}

func (ctl *TodoController) RegisterRoutes(r *gin.Engine) {
	r.GET("/todos", ctl.list)
	r.POST("/todos", ctl.create)
}

func (ctl *TodoController) list(c *gin.Context) {
	todos, err := ctl.svc.ListTodos()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "internal error"})
		return
	}
	c.JSON(http.StatusOK, todos)
}

type createTodoRequest struct {
	Title string `json:"title" binding:"required"`
}

func (ctl *TodoController) create(c *gin.Context) {
	var req createTodoRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "title is required"})
		return
	}

	t, err := ctl.svc.CreateTodo(req.Title)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "internal error"})
		return
	}

	c.JSON(http.StatusCreated, t)
}

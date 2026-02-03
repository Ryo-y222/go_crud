package controller

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"go_crud/internal/repository"
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
	r.PUT("/todos/:id", ctl.updateDone)
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

type updateDoneRequest struct {
	Done bool `json:"done"`
}

func (ctl *TodoController) updateDone(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	var req updateDoneRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid body"})
		return
	}

	err = ctl.svc.UpdateTodoDone(id, req.Done)
	if err != nil {
		if errors.Is(err, repository.ErrTodoNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "todo not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "internal error"})
		return
	}

	c.Status(http.StatusNoContent)
}

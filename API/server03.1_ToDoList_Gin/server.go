package main

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

// Model
type Todo struct {
	ID   int    `json:"id"`
	Task string `json:"task"`
}

// Logger
var logger = logrus.New() 

// Error Response

type ErrorResponse struct{
	Error ErrorDetail `json:"error"`
}
type ErrorDetail struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

func NewErrorResponse(code,message string) ErrorResponse{
	return ErrorResponse{
		Error:ErrorDetail{
			Code: code,
			Message:message,
		},
	}
}

//Controller 
type TodoController struct{
	todos []Todo
	nextID int
}

// Constructor FN
func NewTodoController() *TodoController{
	return &TodoController{
		todos: []Todo{},
		nextID: 1,
	}
}

func (c *TodoController) AddTodo(ctx *gin.Context){
	var t []Todo
	if err := ctx.ShouldBindJSON(&t); err != nil{
		logger.WithError(err).Error("Invalid JSON")
		ctx.JSON(http.StatusBadRequest,NewErrorResponse("ERR_INVALID_ID","Invalid ID"))
		return
	}
	
	for i := range t{
		if strings.TrimSpace(t[i].Task) == ""{
			logger.Warn("Task is required")
			ctx.JSON(http.StatusBadRequest,NewErrorResponse("ERR_TASK_REQUIRED","Task is required"))
			return
		}
		t[i].ID = c.nextID
		c.nextID++
		c.todos = append(c.todos, t[i])

		logger.WithFields(logrus.Fields{
		"id":   t[i].ID,
		"task": t[i].Task,
		}).Info("Added new Todo")

	}
	ctx.JSON(http.StatusOK,t)
}

func (c *TodoController) ListTodo(ctx *gin.Context){
	logger.Info("Listing all todos")
	ctx.JSON(http.StatusOK,c.todos)
}

func (c *TodoController) UpdateTodo(ctx *gin.Context) {
	// Get id of http and convert to int
    idStr := ctx.Param("id")
    id, err := strconv.Atoi(idStr)
    if err != nil || id <= 0 {
		logger.WithField("idParam",idStr).Error("Invalid ID")
        ctx.JSON(http.StatusBadRequest,NewErrorResponse("ERR_INVALID_ID", "Invalid ID"))
        return
    }
	
	var updateData Todo
    if err := ctx.ShouldBindJSON(&updateData); err != nil {
		logger.WithError(err).Error("Invalid JSON")
		ctx.JSON(http.StatusBadRequest,NewErrorResponse("ERR_INVALID_JSON", "Invalid JSON"))
        return
    }
	
	
    if strings.TrimSpace(updateData.Task) == "" {
		logger.Warn("Task is required")
		ctx.JSON(http.StatusBadRequest,NewErrorResponse("ERR_INVALID_REQUIRED", "Task is required"))
        return
    }

    updated := false
    for i, todo := range c.todos {
        if todo.ID == id {
            c.todos[i].Task = updateData.Task
            updated = true
			logger.WithFields(logrus.Fields{
				"id":id,
				"task":updateData.Task,
			}).Info("Todo updated")
            ctx.JSON(http.StatusOK, c.todos[i])
            break
        }
    }
	
    if !updated {
		logger.WithField("id", id).Warn("Todo not found")
		ctx.JSON(http.StatusNotFound, NewErrorResponse("ERR_TODO_NOT_FOUND", "Todo not found"))
		return
	}
}


func (c *TodoController) DeleteTodo(ctx *gin.Context){
	idStr := ctx.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil || id <= 0 {
		logger.WithField("idParam",idStr).Error("Invalid ID")
        ctx.JSON(http.StatusBadRequest,NewErrorResponse("ERR_INVALID_ID", "Invalid ID"))        
		return
	}

	found := false
	for i, data := range c.todos {
		if data.ID == id {
			c.todos = append(c.todos[:i], c.todos[i+1:]...)
			found = true

			logger.WithField("id",id).Info("Todo Deleted")
			break
		}
	}

	if !found {
		logger.WithField("id", id).Warn("Todo not found")
		ctx.JSON(http.StatusNotFound, NewErrorResponse("ERR_TODO_NOT_FOUND", "Todo not found"))
		return
	}

	ctx.Status(http.StatusNoContent)
}


func main(){
	logger.SetFormatter(&logrus.JSONFormatter{})
	logger.SetLevel(logrus.InfoLevel)

	r := gin.Default()
	ctrl := NewTodoController()

	r.GET("/todo/",ctrl.ListTodo)
	r.POST("/todo/",ctrl.AddTodo)
	r.PUT("/todo/edit/:id",ctrl.UpdateTodo)
	r.DELETE("/todo/delete/:id",ctrl.DeleteTodo)
	r.Run(":8080")
}
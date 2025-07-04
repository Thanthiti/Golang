package main

import (
	"context"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// ===== Model =====
type Todo struct {
	gorm.Model
	Task string `json:"task"`
}

// ===== Logger =====
var logger = logrus.New()

// ===== DB =====
var db *gorm.DB

func initDB() {
	var err error
	db, err = gorm.Open(sqlite.Open("todos.db"), &gorm.Config{})
	if err != nil {
		logger.WithError(err).Fatal("Failed to connect to database")
	}
	db.AutoMigrate(&Todo{})
	logger.Info("Database migrated")
}

// ===== Error Response =====
type ErrorResponse struct {
	Error ErrorDetail `json:"error"`
}
type ErrorDetail struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

func NewErrorResponse(code, message string) ErrorResponse {
	return ErrorResponse{
		Error: ErrorDetail{
			Code:    code,
			Message: message,
		},
	}
}

// ===== Controller =====
type TodoController struct{}

func NewTodoController() *TodoController {
	return &TodoController{}
}

func (c *TodoController) AddTodo(ctx *gin.Context) {
	var todo Todo
	if err := ctx.ShouldBindJSON(&todo); err != nil {
		logger.WithError(err).Warn("Invalid JSON for AddTodo")
		ctx.JSON(http.StatusBadRequest, NewErrorResponse("ERR_INVALID_JSON", "Invalid JSON"))
		return
	}

	err := db.Transaction(func(tx *gorm.DB) error {
		todo.Task = strings.TrimSpace(todo.Task)
		if todo.Task == "" {
			logger.Warn("Task is required but empty in AddTodo")
			ctx.JSON(http.StatusBadRequest, NewErrorResponse("ERR_TASK_REQUIRED", "Task is required"))
			return gorm.ErrInvalidData
		}
		if err := tx.Create(&todo).Error; err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		logger.WithError(err).Error("Failed to save todo in AddTodo")
		ctx.JSON(http.StatusInternalServerError, NewErrorResponse("ERR_DB", "Database error"))
		return
	}

	logger.WithField("id", todo.ID).Info("Added new todo")
	ctx.JSON(http.StatusCreated, todo)
}

func (c *TodoController) AddTodoBulk(ctx *gin.Context) {
	var todos []Todo
	if err := ctx.ShouldBindJSON(&todos); err != nil {
		logger.WithError(err).Warn("Invalid JSON for AddTodoBulk")
		ctx.JSON(http.StatusBadRequest, NewErrorResponse("ERR_INVALID_JSON", "Invalid JSON"))
		return
	}

	err := db.Transaction(func(tx *gorm.DB) error {
		for i := range todos {
			todos[i].Task = strings.TrimSpace(todos[i].Task)
			if todos[i].Task == "" {
				logger.Warnf("Task is required but empty in AddTodoBulk at index %d", i)
				ctx.JSON(http.StatusBadRequest, NewErrorResponse("ERR_TASK_REQUIRED", "Task is required"))
				return gorm.ErrInvalidData
			}
			if err := tx.Create(&todos[i]).Error; err != nil {
				return err
			}
		}
		return nil
	})

	if err != nil {
		logger.WithError(err).Error("Failed to save todos in AddTodoBulk")
		ctx.JSON(http.StatusInternalServerError, NewErrorResponse("ERR_DB", "Database error"))
		return
	}

	logger.WithField("count", len(todos)).Info("Added todos in bulk")
	ctx.JSON(http.StatusCreated, todos)
}

func (c *TodoController) ListTodo(ctx *gin.Context) {
	var todos []Todo

	page, _ := strconv.Atoi(ctx.DefaultQuery("page", "1"))
	if page < 1 {
		page = 1
	}
	size, _ := strconv.Atoi(ctx.DefaultQuery("size", "10"))
	if size < 1 {
		size = 10
	}

	ctxWithTimeout, cancel := context.WithTimeout(ctx.Request.Context(), 5*time.Second)
	defer cancel()

	result := db.WithContext(ctxWithTimeout).
		Where("deleted_at IS NULL").
		Offset((page - 1) * size).
		Limit(size).
		Order("created_at DESC").
		Find(&todos)

	if result.Error != nil {
		logger.WithError(result.Error).Error("Database error in ListTodo")
		ctx.JSON(http.StatusInternalServerError, NewErrorResponse("ERR_DB", "Database error"))
		return
	}

	logger.WithFields(logrus.Fields{
		"page": page,
		"size": size,
		"count": len(todos),
	}).Info("Listed todos")
	ctx.JSON(http.StatusOK, todos)
}

func (c *TodoController) UpdateTodo(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil || id <= 0 {
		logger.Warnf("Invalid ID param in UpdateTodo: %s", idStr)
		ctx.JSON(http.StatusBadRequest, NewErrorResponse("ERR_INVALID_ID", "Invalid ID"))
		return
	}

	var updateData Todo
	if err := ctx.ShouldBindJSON(&updateData); err != nil {
		logger.WithError(err).Warn("Invalid JSON in UpdateTodo")
		ctx.JSON(http.StatusBadRequest, NewErrorResponse("ERR_INVALID_JSON", "Invalid JSON"))
		return
	}

	updateData.Task = strings.TrimSpace(updateData.Task)
	if updateData.Task == "" {
		logger.Warn("Task is required but empty in UpdateTodo")
		ctx.JSON(http.StatusBadRequest, NewErrorResponse("ERR_TASK_REQUIRED", "Task is required"))
		return
	}

	var todo Todo
	result := db.First(&todo, id)
	if result.Error != nil {
		logger.Warnf("Todo not found in UpdateTodo: id=%d", id)
		ctx.JSON(http.StatusNotFound, NewErrorResponse("ERR_TODO_NOT_FOUND", "Todo not found"))
		return
	}

	todo.Task = updateData.Task
	todo.UpdatedAt = time.Now()

	if err := db.Save(&todo).Error; err != nil {
		logger.WithError(err).Error("Database error in UpdateTodo")
		ctx.JSON(http.StatusInternalServerError, NewErrorResponse("ERR_DB", "Database error"))
		return
	}

	logger.WithFields(logrus.Fields{
		"id":   todo.ID,
		"task": todo.Task,
	}).Info("Updated todo")
	ctx.JSON(http.StatusOK, todo)
}

func (c *TodoController) DeleteTodo(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil || id <= 0 {
		logger.Warnf("Invalid ID param in DeleteTodo: %s", idStr)
		ctx.JSON(http.StatusBadRequest, NewErrorResponse("ERR_INVALID_ID", "Invalid ID"))
		return
	}

	result := db.Delete(&Todo{}, id)
	if result.RowsAffected == 0 {
		logger.Warnf("Todo not found in DeleteTodo: id=%d", id)
		ctx.JSON(http.StatusNotFound, NewErrorResponse("ERR_TODO_NOT_FOUND", "Todo not found"))
		return
	}

	logger.WithField("id", id).Info("Deleted todo")
	ctx.Status(http.StatusNoContent)
}

// ===== Main =====
func main() {
	logger.SetFormatter(&logrus.JSONFormatter{})
	logger.SetLevel(logrus.InfoLevel)

	initDB()

	r := gin.Default()
	ctrl := NewTodoController()

	r.GET("/todo", ctrl.ListTodo)
	r.POST("/todo", ctrl.AddTodo)
	r.POST("/todo/bulk", ctrl.AddTodoBulk)
	r.PUT("/todo/:id", ctrl.UpdateTodo)
	r.DELETE("/todo/:id", ctrl.DeleteTodo)

	logger.Info("Server running on :8080")
	r.Run(":8080")
}

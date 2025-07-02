package controller

// import (
// 	"mymodule/API/server03.2_ToDoList_Gin/db"
// 	"mymodule/API/server03.2_ToDoList_Gin/model"
// 	"net/http"

// 	"github.com/gin-gonic/gin"
// )

// func GetTodos(c *gin.Context) {
// 	var todos []model.Todo
// 	result := db.DB.Find(&todos)
// 	if result.Error != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
// 		return
// 	}
// 	c.JSON(http.StatusOK, todos)
// }
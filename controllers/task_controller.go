// controllers/task_controller.go
package controllers

import (
	"net/http"
	"strconv"
	"taskmanager-api/models"
	"taskmanager-api/services"

	"fmt"

	"github.com/gin-gonic/gin"
)

func CreateTask(c *gin.Context) {
	var input models.Task
	fmt.Println("===============================================")
	fmt.Println(c)
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	fmt.Println(c.Request)

	task := services.CreateTask(input.Title, input.Description)
	c.JSON(http.StatusCreated, task)
}

func GetTasks(c *gin.Context) {
	fmt.Println("===============================================")
	fmt.Println(c.Request.Body)
	tasks := services.GetTasks()
	c.JSON(http.StatusOK, tasks)
}

func GetTaskByID(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	task := services.GetTaskByID(id)
	if task == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
		return
	}
	c.JSON(http.StatusOK, task)
}

func UpdateTask(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var input models.Task
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	updated := services.UpdateTask(id, input)
	if !updated {
		c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Task updated successfully"})
}

func DeleteTask(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	deleted := services.DeleteTask(id)
	if !deleted {
		c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Task deleted successfully"})
}

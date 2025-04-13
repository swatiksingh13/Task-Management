package handlers

import (
	"Task_Manager/models"
	"Task_Manager/services"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

func GetTasks(c *gin.Context) {
	log.Println("[GET] /tasks - Fetching all tasks")

	tasks, err := services.GetAllTasks()
	if err != nil {
		log.Printf("Error fetching tasks: %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch tasks"})
		return
	}

	log.Printf("Fetched %d tasks successfully\n", len(tasks))
	c.JSON(http.StatusOK, tasks)
}

func CreateTask(c *gin.Context) {
	log.Println("[POST] /tasks - Creating new task")

	var task models.Task
	if err := c.ShouldBindJSON(&task); err != nil {
		log.Printf("Invalid task data: %v\n", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	log.Printf("Creating task: %+v\n", task)
	if err := services.CreateTask(&task); err != nil {
		log.Printf("Failed to create task: %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create task"})
		return
	}

	log.Println("Task created successfully")
	c.JSON(http.StatusCreated, task)
}

func UpdateTask(c *gin.Context) {
	id := c.Param("id")
	taskID, err := strconv.Atoi(id)
	if err != nil {
		log.Printf("Invalid task ID: %v\n", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid task ID"})
		return
	}

	var task models.Task
	if err := c.ShouldBindJSON(&task); err != nil {
		log.Printf("Invalid task data for update: %v\n", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid task data"})
		return
	}

	log.Printf("Updating task ID %d with data: %+v\n", taskID, task)
	if err := services.UpdateTask(taskID, &task); err != nil {
		log.Printf("Failed to update task ID %d: %v\n", taskID, err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update task"})
		return
	}

	log.Printf("Task ID %d updated successfully\n", taskID)
	c.JSON(http.StatusOK, task)
}

func DeleteTask(c *gin.Context) {
	id := c.Param("id")
	taskID, err := strconv.Atoi(id)
	if err != nil {
		log.Printf("Invalid task ID for delete: %v\n", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid task ID"})
		return
	}

	log.Printf("Deleting task ID %d\n", taskID)
	if err := services.DeleteTask(taskID); err != nil {
		log.Printf("Failed to delete task ID %d: %v\n", taskID, err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete task"})
		return
	}

	log.Printf("Task ID %d deleted successfully\n", taskID)
	c.JSON(http.StatusNoContent, gin.H{"message": "Task deleted successfully"})
}

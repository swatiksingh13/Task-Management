package services

import (
	"Task_Manager/database"
	"Task_Manager/models"
	"errors"
	"log"
)

func GetAllTasks() ([]models.Task, error) {
	log.Println("Fetching all tasks from DB")

	var tasks []models.Task
	if err := database.DB.Find(&tasks).Error; err != nil {
		log.Printf("Error: %v\n", err)
		return nil, err
	}

	log.Printf("Successfully fetched %d tasks\n", len(tasks))
	return tasks, nil
}

func CreateTask(task *models.Task) error {
	log.Printf("Creating new task: %+v\n", task)

	if err := database.DB.Create(&task).Error; err != nil {
		log.Printf("Failed to create task: %v\n", err)
		return err
	}

	log.Printf("Task created successfully with ID %d\n", task.ID)
	return nil
}

func UpdateTask(id int, task *models.Task) error {
	log.Printf("Updating task ID %d\n", id)

	var existingTask models.Task
	if err := database.DB.First(&existingTask, id).Error; err != nil {
		log.Printf("Task not found with ID %d: %v\n", id, err)
		return errors.New("task not found")
	}

	log.Printf("Existing Task: %+v\n", existingTask)

	existingTask.Title = task.Title
	existingTask.Description = task.Description
	existingTask.Status = task.Status
	existingTask.DueDate = task.DueDate

	if err := database.DB.Save(&existingTask).Error; err != nil {
		log.Printf("Failed to update task ID %d: %v\n", id, err)
		return err
	}

	log.Printf("Task ID %d updated successfully\n", id)
	task.ID = existingTask.ID
	return nil
}

func DeleteTask(id int) error {
	log.Printf("Deleting task ID %d\n", id)

	if err := database.DB.Delete(&models.Task{}, id).Error; err != nil {
		log.Printf("Failed to delete task ID %d: %v\n", id, err)
		return err
	}

	log.Printf("Task ID %d deleted successfully\n", id)
	return nil
}

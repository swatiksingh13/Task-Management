package handlers

import (
	"Task_Manager/database"
	"Task_Manager/models"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"testing"
)

func init() {
	gin.SetMode(gin.TestMode)
	database.Connect()
	database.DB.Exec("DELETE FROM tasks")
}

func TestCreateTask(t *testing.T) {
	router := gin.Default()
	router.POST("/tasks", CreateTask)

	payload := `{
		"title": "Test Task",
		"description": "This is a test task",
		"status": "Pending",
		"due_date": "2025-04-15T14:30:00Z"
	}`

	req, _ := http.NewRequest("POST", "/tasks", strings.NewReader(payload))
	req.Header.Set("Content-Type", "application/json")
	resp := httptest.NewRecorder()

	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusCreated, resp.Code)

	var task models.Task
	err := json.Unmarshal(resp.Body.Bytes(), &task)
	assert.Nil(t, err)
	assert.Equal(t, "Test Task", task.Title)
	assert.Equal(t, "Pending", task.Status)
}

func TestGetTasks(t *testing.T) {
	router := gin.Default()
	router.GET("/tasks", GetTasks)

	req, _ := http.NewRequest("GET", "/tasks", nil)
	resp := httptest.NewRecorder()

	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusOK, resp.Code)

	var tasks []models.Task
	err := json.Unmarshal(resp.Body.Bytes(), &tasks)
	assert.Nil(t, err)
	assert.GreaterOrEqual(t, len(tasks), 1)
}

func TestUpdateTask(t *testing.T) {
	dummy := models.Task{
		Title:       "Update Me",
		Description: "Initial Desc",
		Status:      "Pending",
	}
	database.DB.Create(&dummy)

	router := gin.Default()
	router.PUT("/tasks/:id", UpdateTask)

	payload := `{
		"title": "Updated Task",
		"description": "Updated Desc",
		"status": "Completed"
	}`

	url := "/tasks/" + strconv.Itoa(int(dummy.ID))
	req, _ := http.NewRequest("PUT", url, strings.NewReader(payload))
	req.Header.Set("Content-Type", "application/json")
	resp := httptest.NewRecorder()

	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusOK, resp.Code)

	var updated models.Task
	json.Unmarshal(resp.Body.Bytes(), &updated)
	assert.Equal(t, "Completed", updated.Status)
}

func TestDeleteTask(t *testing.T) {
	dummy := models.Task{
		Title:       "Delete Me",
		Description: "Delete Desc",
		Status:      "Pending",
	}
	database.DB.Create(&dummy)

	router := gin.Default()
	router.DELETE("/tasks/:id", DeleteTask)

	url := "/tasks/" + strconv.Itoa(int(dummy.ID))
	req, _ := http.NewRequest("DELETE", url, nil)
	resp := httptest.NewRecorder()

	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusNoContent, resp.Code)

	var task models.Task
	err := database.DB.First(&task, dummy.ID).Error
	assert.NotNil(t, err)
}

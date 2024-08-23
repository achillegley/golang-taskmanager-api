package services

import (
	"taskmanager-api/database"
	"taskmanager-api/models"
)

func CreateTask(title, description string) models.Task {
	task := models.Task{
		Title:       title,
		Description: description,
		Completed:   false,
	}
	database.DB.Create(&task)
	return task
}

func GetTasks() []models.Task {
	var tasks []models.Task
	database.DB.Find(&tasks)
	return tasks
}

func GetTaskByID(id int) *models.Task {

	var task models.Task
	result := database.DB.First(&task, id)
	if result.Error != nil {
		return nil
	}
	return &task
}

func UpdateTask(id int, updatedTask models.Task) bool {
	var task models.Task
	result := database.DB.First(&task, id)
	if result.Error != nil {
		return false
	}
	task.Title = updatedTask.Title
	task.Description = updatedTask.Description
	task.Completed = updatedTask.Completed

	database.DB.Save(&task)
	return true

}

func DeleteTask(id int) bool {
	var task models.Task
	result := database.DB.First(&task, id)
	if result.Error != nil {
		return false
	}

	database.DB.Delete(&task)

	return true
}

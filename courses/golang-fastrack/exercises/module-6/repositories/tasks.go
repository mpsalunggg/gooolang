package repositories

import "module-6/models"

var tasks = []models.Task{}

func GetAllTasks() []models.Task {
	return tasks
}

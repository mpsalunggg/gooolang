package repositories

import "module-6/models"

type TaskRepository struct {
	tasks  []models.Task
	nextID int
}

func NewTaskRepository() *TaskRepository {
	return &TaskRepository{
		tasks:  []models.Task{},
		nextID: 1,
	}
}

func (r *TaskRepository) GetAllTasks() []models.Task {
	return r.tasks
}

func (r *TaskRepository) CreateTask(task models.TaskRequest) models.Task {
	newTask := models.Task{
		ID:          r.nextID,
		Title:       task.Title,
		Description: task.Description,
		Status:      task.Status,
	}

	r.tasks = append(r.tasks, newTask)
	r.nextID++
	return newTask
}

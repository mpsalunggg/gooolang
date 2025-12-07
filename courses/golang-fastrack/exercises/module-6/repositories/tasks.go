package repositories

import (
	"fmt"
	"module-6/models"
)

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

func (r *TaskRepository) GetTaskByID(id int) (models.Task, error) {
	for _, task := range r.tasks {
		if task.ID == id {
			return task, nil
		}
	}
	return models.Task{}, fmt.Errorf("task with ID %d not found", id)
}

func (r *TaskRepository) UpdateTask(id int, task models.TaskRequest) (models.Task, error) {
	for index, t := range r.tasks {
		if t.ID == id {
			r.tasks[index] = models.Task{
				ID:          id,
				Title:       task.Title,
				Description: task.Description,
				Status:      task.Status,
			}

			return r.tasks[index], nil
		}
	}
	return models.Task{}, fmt.Errorf("task with ID %d not found", id)
}

func (r *TaskRepository) DeleteTask(id int) error {
	for index, task := range r.tasks {
		if task.ID == id {
			r.tasks = append(r.tasks[:index], r.tasks[index+1:]...)
			return nil
		}
	}

	return fmt.Errorf("task with ID %d not found", id)
}

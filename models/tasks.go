package models

import (
	"slices"
	"time"
)

type Task struct {
	ID   int       `json:"id"`
	Text string    `json:"text"`
	Tags []string  `json:"tags"`
	Due  time.Time `json:"due"`
}

var TasksDB []Task

// Создание задачи
func CreateTask(task Task) int {
	task.ID = len(TasksDB) + 1
	TasksDB = append(TasksDB, task)
	return task.ID
}

// Получение задачи по ID.
func GetTaskByID(id int) (Task, bool) {
	for _, task := range TasksDB {
		if task.ID == id {
			return task, true
		}
	}
	return Task{}, false
}

// Удаление задачи по ID.
func DeleteTaskByID(id int) bool {
	for i, task := range TasksDB {
		if task.ID == id {
			TasksDB = slices.Delete(TasksDB, i, i+1)
			return true
		}
	}
	return false
}

// Удаляем все задачи
func DeleteAllTasks() {
	TasksDB = []Task{}
}

// Получение задач по дате
func GetTasksByDueDate(year, month, day int) []Task {
	var tasksByDate []Task
	for _, task := range TasksDB {
		if task.Due.Year() == year && int(task.Due.Month()) == month && task.Due.Day() == day {
			tasksByDate = append(tasksByDate, task)
		}
	}
	return tasksByDate
}

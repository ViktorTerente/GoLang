package handlers

import (
	"TasksAPI/models"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func initHeaders(writer http.ResponseWriter) {
	writer.Header().Set("Content-Type", "application/json")
}

// Создание новой задачи
func CreateTaskHandler(writer http.ResponseWriter, request *http.Request) {
	initHeaders(writer)
	var task models.Task
	err := json.NewDecoder(request.Body).Decode(&task)
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(writer).Encode(map[string]string{"message": "Invalid JSON format"})
		return
	}
	taskID := models.CreateTask(task)
	writer.WriteHeader(http.StatusCreated)
	json.NewEncoder(writer).Encode(map[string]int{"id": taskID})
}

// Получить задачу по  ID.
func GetTaskByIDHandler(writer http.ResponseWriter, request *http.Request) {
	initHeaders(writer)
	id, err := strconv.Atoi(mux.Vars(request)["taskid"])
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(writer).Encode(map[string]string{"message": "Invalid task ID"})
		return
	}
	task, found := models.GetTaskByID(id)
	if !found {
		writer.WriteHeader(http.StatusNotFound)
		json.NewEncoder(writer).Encode(map[string]string{"message": "Task not found"})
		return
	}
	writer.WriteHeader(http.StatusOK)
	json.NewEncoder(writer).Encode(task)
}

// Получить список всех задач
func GetAllTasksHandler(writer http.ResponseWriter, request *http.Request) {
	initHeaders(writer)
	writer.WriteHeader(http.StatusOK)
	json.NewEncoder(writer).Encode(models.TasksDB)
}

// Удалить задачу по ID.
func DeleteTaskByIDHandler(writer http.ResponseWriter, request *http.Request) {
	initHeaders(writer)
	id, err := strconv.Atoi(mux.Vars(request)["taskid"])
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(writer).Encode(map[string]string{"message": "Invalid task ID"})
		return
	}
	if models.DeleteTaskByID(id) {
		writer.WriteHeader(http.StatusOK)
		json.NewEncoder(writer).Encode(map[string]string{"message": "Task deleted"})
	} else {
		writer.WriteHeader(http.StatusNotFound)
		json.NewEncoder(writer).Encode(map[string]string{"message": "Task not found"})
	}
}

// Удалить все задачи
func DeleteAllTasksHandler(writer http.ResponseWriter, request *http.Request) {
	initHeaders(writer)
	models.DeleteAllTasks()
	writer.WriteHeader(http.StatusOK)
	json.NewEncoder(writer).Encode(map[string]string{"message": "All tasks deleted"})
}

// Получение задач на дату
func GetTasksByDueDateHandler(writer http.ResponseWriter, request *http.Request) {
	initHeaders(writer)
	year, _ := strconv.Atoi(mux.Vars(request)["yy"])
	month, _ := strconv.Atoi(mux.Vars(request)["mm"])
	day, _ := strconv.Atoi(mux.Vars(request)["dd"])
	tasks := models.GetTasksByDueDate(year, month, day)
	writer.WriteHeader(http.StatusOK)
	json.NewEncoder(writer).Encode(tasks)
}

package utils

import (
	"TasksAPI/handlers"

	"github.com/gorilla/mux"
)

func BuildTaskRoutes(router *mux.Router) {
	router.HandleFunc("/tasks", handlers.CreateTaskHandler).Methods("POST")
	router.HandleFunc("/tasks/{taskid}", handlers.GetTaskByIDHandler).Methods("GET")
	router.HandleFunc("/tasks", handlers.GetAllTasksHandler).Methods("GET")
	router.HandleFunc("/tasks/{taskid}", handlers.DeleteTaskByIDHandler).Methods("DELETE")
	router.HandleFunc("/tasks", handlers.DeleteAllTasksHandler).Methods("DELETE")
	router.HandleFunc("/due/{yy}/{mm}/{dd}", handlers.GetTasksByDueDateHandler).Methods("GET")
}

package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Task struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

var tasks []Task

func addTask(w http.ResponseWriter, r *http.Request) {
	var task Task
	// Decode the request body into the task struct
	err := json.NewDecoder(r.Body).Decode(&task)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Generate an ID for the task
	task.ID = len(tasks) + 1
	// Add the task to the tasks slice
	tasks = append(tasks, task)

	w.Header().Set("Content-Type", "application/json")
	// Return the added task as JSON
	json.NewEncoder(w).Encode(task)
}

func getTasks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	// Return all tasks as JSON
	json.NewEncoder(w).Encode(tasks)
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/tasks", getTasks).Methods("GET")
	router.HandleFunc("/tasks", addTask).Methods("POST")

	log.Fatal(http.ListenAndServe(":8080", router))
}

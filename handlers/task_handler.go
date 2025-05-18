package handlers

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/dvdmarveira/simple-crud-go/models"
)

type TaskHandler struct {
	DB *sql.DB
}

// Construtor para o TaskHandler
func NewTaskHandler(db *sql.DB) *TaskHandler {
	return &TaskHandler{DB: db}
}

func (taskHandler *TaskHandler) ReadTasks(writer http.ResponseWriter, request *http.Request){

	rows, err := taskHandler.DB.Query("SELECT * FROM tasks")
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}

		var tasks []models.Task

		for rows.Next() {
			var task models.Task
			err := rows.Scan(&task.ID, &task.Title, &task.Description, &task.Status)

			if err != nil {
				http.Error(writer, err.Error(), http.StatusInternalServerError)
				return
			}
			
			tasks = append(tasks, task)
		}
		
		writer.Header().Set("Content-Type", "application/json")
		json.NewEncoder(writer).Encode(tasks)
}

func (taskHandler *TaskHandler) CreateTask(writer http.ResponseWriter, request *http.Request) {
	var task models.Task
	err := json.NewDecoder(request.Body).Decode(&task)

	if err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	}

	query := "INSERT INTO tasks (title, description, status) VALUES ($1, $2, $3) RETURNING id"
	err = taskHandler.DB.QueryRow(query, task.Title, task.Description, task.Status).Scan(&task.ID)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}

	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusCreated)
	json.NewEncoder(writer).Encode(task)
}
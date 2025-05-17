package main

import (
	"log"
	"net/http"

	"github.com/dvdmarveira/simple-crud-go/config"
	"github.com/dvdmarveira/simple-crud-go/handlers"
	"github.com/dvdmarveira/simple-crud-go/models"
	"github.com/gorilla/mux"
)

// Função chamada quando a aplicação é iniciada
// Ponto de entrada da aplicação
func main() {
		// Chama a função SetupDB
		db := config.SetupDB()

		// Chama a função Close para fechar a conexão com o banco de dados
		defer db.Close()

		_, err := db.Exec(models.CreateTableSQL)
		if err != nil {
				log.Fatal(err)
		}

		router := mux.NewRouter()

		taskHandler := handlers.NewTaskHandler(db)

		router.HandleFunc("/tasks", taskHandler.ReadTasks).Methods("GET")
		router.HandleFunc("/tasks", taskHandler.CreateTask).Methods("POST")
		// router.HandleFunc("/tasks/{id}", taskHandler.UpdateTask).Methods("PUT")
		// router.HandleFunc("/tasks/{id}", taskHandler.DeleteTask).Methods("DELETE")

		// Define o manipulador de rotas
		log.Fatal(http.ListenAndServe(":8080", router))
}
package main

import (
	"log"
	"net/http"

	"github.com/dvdmarveira/simple-crud-go/config"
	"github.com/dvdmarveira/simple-crud-go/models"
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

		// Define o manipulador de rotas
		log.Fatal(http.ListenAndServe(":8080", nil))
}
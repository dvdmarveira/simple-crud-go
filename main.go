package main

import (
	"log"
	"net/http"

	"github.com/dvdmarveira/simple-crud-go/config"
)

// Função chamada quando a aplicação é iniciada
// Ponto de entrada da aplicação
func main() {
		// Chama a função SetupDB
		dbConnection := config.SetupDB()

		// Chama a função Close para fechar a conexão com o banco de dados
		defer dbConnection.Close()

		log.Fatal(http.ListenAndServe(":8080", nil))
}
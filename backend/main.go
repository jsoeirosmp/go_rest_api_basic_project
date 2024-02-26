package main

import (
	"fmt"

	"github.com/jsoeirosmp/go_rest_api_basic_project/backend/database"
	"github.com/jsoeirosmp/go_rest_api_basic_project/backend/routes"
)

func main() {
	// models.Personalidades = []models.Personalidade{
	// 	{Id: 1, Nome: "Nome 1", Historia: "Historia 1"},
	// 	{Id: 2, Nome: "Nome 2", Historia: "Historia 2"},
	// }
	database.ConectaComBancoDeDados()
	fmt.Println("Iniciando o servidor Go/ REST")
	routes.HandleResquest()
}

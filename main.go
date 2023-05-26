package main

import (
	"github.com/Mariano-JR/sistema-de-login/database"
	"github.com/Mariano-JR/sistema-de-login/routes"
)

func main() {
	database.ConectaDB()
	routes.HandleRequests()
}

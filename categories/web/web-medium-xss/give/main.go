package main

import (
	"example/v3/handlers"

	_ "github.com/go-sql-driver/mysql"
)

func main() {

	router := handlers.SetupRouter()
	router.LoadHTMLGlob("templates/*")
	router.Run(":60888")
}

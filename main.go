package main

import (
	"git/config"
	"git/router"
)

func main() {

	config.InitDB()

	router.Inicializar() // Initialize the router
}

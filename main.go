package main

import (
	"Api-Aula1/config"
	"Api-Aula1/router"
	"fmt"
	"log"
	"net/http"
)

func main() {
	config.LoadEnv()

	r := router.New()

	fmt.Printf("Servidor rodando em %s ...\n", config.Port)
	log.Fatal(http.ListenAndServe(config.Port, r))
}

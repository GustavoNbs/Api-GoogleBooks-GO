package config

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var (
	StringConexaoBanco = ""
	Porta              = 0
)

func LoadEnv() {
	var erro error

	if erro = godotenv.Load(); erro != nil {
		log.Fatal("Erro ao carregar o arquivo .env")
	}

	Porta, erro = strconv.Atoi(os.Getenv("API_PORT"))
	if erro != nil {
		Porta = 8080 // porta padrão
	}

	StringConexaoBanco = fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_ADDR"),
		os.Getenv("DB_DATABASE"),
	)
}

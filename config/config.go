package config

import (
	"log"
	"os"
	"path/filepath"
	"runtime"

	"github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

var (
	Port      string
	Cfg       mysql.Config
	SecretKey []byte
)

func LoadEnv() {
	var erro error
	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		log.Fatal("Não foi possível obter o caminho do arquivo de configuração")
	}

	configDir := filepath.Dir(filename)
	envPath := filepath.Join(configDir, ".env")

	if erro = godotenv.Load(envPath); erro != nil {
		baseDir := filepath.Dir(configDir)
		envPathRoot := filepath.Join(baseDir, ".env")
		if erro = godotenv.Load(envPathRoot); erro != nil {
			log.Fatal("Erro ao carregar arquivos .env")
		}
	}

	Port = os.Getenv("API_PORT")
	if Port == "" {
		Port = ":8080"
	}

	Cfg = mysql.Config{
		User:                 os.Getenv("DB_USER"),
		Passwd:               os.Getenv("DB_PASSWORD"),
		Net:                  "tcp",
		Addr:                 os.Getenv("DB_ADDR"),
		DBName:               os.Getenv("DB_DATABASE"),
		AllowNativePasswords: true,
	}

	SecretKey = []byte(os.Getenv("SECRET_KEY"))
}

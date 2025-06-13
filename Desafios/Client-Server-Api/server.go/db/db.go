package db

import (
	"fmt"
	"log"
	"moduloInicial/Desafios/Client-Server-Api/server.go/model"

	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type DBConfig struct {
	Driver     string
	Host       string
	Port       int
	User       string
	Password   string
	DBName     string
	SSLMode    string
	SQLitePath string
}

func NewDB(config DBConfig) *gorm.DB {
	var dialector gorm.Dialector

	switch config.Driver {
	case "postgres":
		dsn := fmt.Sprintf(
			"host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
			config.Host, config.Port, config.User, config.Password, config.DBName, config.SSLMode,
		)
		dialector = postgres.Open(dsn)

	case "mysql":
		dsn := fmt.Sprintf(
			"%s:%s@tcp(%s:%d)/%s?parseTime=true",
			config.User, config.Password, config.Host, config.Port, config.DBName,
		)
		dialector = mysql.Open(dsn)

	case "sqlite":
		if config.SQLitePath == "" {
			log.Fatal("Caminho do SQLite não informado")
		}
		dialector = sqlite.Open(config.SQLitePath)

	default:
		log.Fatalf("Driver de banco não suportado: %s", config.Driver)
	}

	db, err := gorm.Open(dialector, &gorm.Config{})
	if err != nil {
		log.Fatalf("Erro ao abrir banco de dados: %v", err)
	}

	err = db.AutoMigrate(&model.CotacaoDB{})
	if err != nil {
		log.Fatalf("Erro ao migrar modelo Cotacao: %v", err)
	}

	return db
}

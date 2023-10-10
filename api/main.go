package main

import (
	"fmt"
	"log"

	"ddvinyaninov/assets-tracker-api/internal/handler"
	"ddvinyaninov/assets-tracker-api/internal/repository"
	"ddvinyaninov/assets-tracker-api/internal/service"

	"github.com/jmoiron/sqlx"
	"github.com/kelseyhightower/envconfig"
	_ "github.com/lib/pq"
)

type Config struct {
	Dsn     string `required:"true"`
	GinPort string `required:"true" split_words:"true"`
}

func LoadConfig() (*Config, error) {
	var c Config
	err := envconfig.Process("", &c)
	if err != nil {
		return nil, fmt.Errorf("config failed: %w", err)
	}
	return &c, nil
}

func NewPostgresDB(conf *Config) (*sqlx.DB, error) {
	db, err := sqlx.Connect("postgres", conf.Dsn)
	if err != nil {
		return nil, err
	}
	return db, nil
}

func main() {
	conf, err := LoadConfig()
	if err != nil {
		log.Fatalln(err)
	}

	db, err := NewPostgresDB(conf)
	if err != nil {
		log.Fatalln(err)
	}

	log.Println("db: ", db)

	apiRepository := repository.NewApiRepository(db)
	apiService := service.NewApiService(apiRepository)
	apiHandler := handler.NewApiHandler(apiService)

	ginServer := handler.SetupHandlers(apiHandler)
	if err := ginServer.Run(conf.GinPort); err != nil {
		log.Fatalln(err)
	}

}

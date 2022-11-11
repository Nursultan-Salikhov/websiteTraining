package main

import (
	_ "github.com/lib/pq"
	"github.com/spf13/viper"
	"log"
	todo "websiteTraining"
	"websiteTraining/pkg/handler"
	"websiteTraining/pkg/repository"
	service "websiteTraining/pkg/service"
)

func init() {
	err := initConfig()
	if err != nil {
		log.Fatalf("Error initConfig: %s", err.Error())
	}
}

func main() {
	db, err := repository.NewPostgresDB(repository.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		Password: viper.GetString("db.password"),
		DBName:   viper.GetString("db.dbname"),
		SSLMode:  viper.GetString("db.sslmode"),
	})

	if err != nil {
		log.Fatalf("Failed to init db: %s", err.Error())
	}
	rep := repository.NewRepository(db)
	services := service.NewService(rep)
	handlers := handler.NewHandler(services)

	srv := new(todo.Server)
	if err := srv.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil {
		log.Fatalf("Error occured while running http server: %s", err.Error())
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}

package main

import (
	// "fmt"
	// "fmt"
	// "norbekov/configs"
	// "norbekov/package/handler"
	// "norbekov/package/repository"
	// "norbekov/package/service"
	// "norbekov/server"
	// "norbekov/util/logrus"

	"os"

	_ "github.com/lib/pq"
	"github.com/mrboburs/Norbekov/configs"
	"github.com/mrboburs/Norbekov/package/handler"
	"github.com/mrboburs/Norbekov/package/repository"
	"github.com/mrboburs/Norbekov/package/service"
	"github.com/mrboburs/Norbekov/server"
	"github.com/mrboburs/Norbekov/util/logrus"
)

// @title Norbekov API
// @version 1.0
// @description API Server for Norbekov App
// @host localhost:8080
// @BasePath /
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
// @contact.name   Mr Bobur

func main() {

	logrus := logrus.GetLogger()

	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = "8080"
	}
	configs, err := configs.InitConfig()
	logrus.Infof("configs %v", configs)
	if err != nil {
		logrus.Fatalf("error initializing configs: %s", err.Error())
	}
	logrus.Info("successfull checked configs.")
	db, err := repository.NewPostgresDB(repository.Config{
		Host:     configs.DBHost,
		Port:     configs.DBPort,
		Username: configs.DBUsername,
		DBName:   configs.DBName,
		SSLMode:  configs.DBSSLMode,
		Password: configs.DBPassword,
	}, logrus)

	if err != nil {
		logrus.Fatalf("failed to initialize db: %s", err.Error())
	}

	logrus.Info("successfull connection DB")

	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services, logrus, configs)

	server := new(server.Server)
	err = server.Run(port, handlers.InitRoutes())

	if err != nil {
		logrus.Fatalf("error occurred while running http server: %s", err.Error())
	}

	defer logrus.Infof("run server port:%v", port)
}

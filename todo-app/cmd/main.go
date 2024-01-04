package main

import (
	"context"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"log"
	todo "myproject"
	"myproject/pkg/handler"
	"myproject/pkg/repository"
	"myproject/pkg/service"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	logrus.SetFormatter(new(logrus.JSONFormatter))

	if err := initCnfig(); err != nil {
		logrus.Fatal("error initializing configs: %s", err.Error())
	}

	if err := godotenv.Load(); err != nil {
		logrus.Fatalf("error loading env variables: %s", err.Error())
	}

	db, err := repository.NewPostgresDB(repository.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		Password: os.Getenv("DB_PASSWORD"),
		DBName:   viper.GetString("db.dbname"),
		SSLMode:  viper.GetString("db.sslmode"),
	})
	if err != nil {
		logrus.Fatal("failed to initialize db2: ", err.Error())
	}
	// ...
	if err := db.Ping(); err != nil {
		log.Fatalf("failed to connect to the database: %s", err.Error())
	}
	// ...

	//authRepo := repository.NewAuthPostgres(db)

	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	huj := handler.NewHandler(services)
	srv := new(todo.Server)

	go func() {
		if err := srv.Run(viper.GetString("port"), huj.InitRouters()); err != nil {
			logrus.Fatalf("error server %s", err.Error())
		}
	}()

	logrus.Print("TodoApp Started")

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	logrus.Print("TodoApp Shutting Down")

	if err := srv.Shutdown(context.Background()); err != nil {
		logrus.Errorf("error occured on server shutting down: %s", err.Error())
	}

	if err := srv.Shutdown(context.Background()); err != nil {
		logrus.Errorf("error occured on db connection close: %s", err.Error())
	}
}

func initCnfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}

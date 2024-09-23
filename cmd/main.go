package main

import (
	"context"
	"fmt"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"log"
	"net/http"
	"os"
	"todo"
	"todo/pkg/handler"
	"todo/pkg/logger"
	"todo/pkg/repository"
	"todo/pkg/service"
)

func main() {
	if err := initConfig(); err != nil {
		log.Fatalf("Failed to init configs: %v", err)
	}

	Log := logger.Setup(logger.Config{
		Evn: viper.GetString("evn"),
	})

	defer func(log *zap.Logger) {
		err := log.Sync()
		if err != nil {
			log.Warn(fmt.Sprintf("Failed to sync zap logger, errror: %s", err))
		}
	}(Log)

	if err := godotenv.Load(); err != nil {
		Log.Fatal(fmt.Sprintf("Failed to load .env, error: %s", err))
	}

	db, err := repository.NewPostgresBD(repository.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		Password: os.Getenv("BD_PASSWORD"),
		DBName:   viper.GetString("db.dbname"),
		SSLMode:  viper.GetString("db.ssl_mode"),
	})

	if err != nil {
		Log.Fatal(fmt.Sprintf("Failed to init db connection, error: %s", err))
	}

	repos := repository.NewRepository(db, *Log)
	services := service.NewService(repos, *Log)
	handlers := handler.NewHandler(services, *Log)

	srv := new(todo.Server)
	if err := srv.Run(&http.Server{
		Addr:           ":" + viper.GetString("port"),
		Handler:        handlers.InitRoutes(),
		MaxHeaderBytes: 1 << 20,
		ReadTimeout:    viper.GetDuration("timeout"),
		WriteTimeout:   viper.GetDuration("timeout"),
	}); err != nil {
		Log.Fatal(fmt.Sprintf("Failed to start server, error: %s", err))
	}

	defer func(srv *todo.Server, ctx context.Context) {
		if err := srv.Shutdown(ctx); err != nil {
			Log.Error(fmt.Sprintf("Failed to shutdown server, error: %s", err))
		}
	}(srv, context.Background())
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}

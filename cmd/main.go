package main

import (
	"errors"
	"fmt"

	"github.com/tasks-hub/users-service/internal/app"
	"github.com/tasks-hub/users-service/internal/config"

	"github.com/Netflix/go-env"
	"go.uber.org/zap"
)

func main() {
	var conf config.Config
	_, err := env.UnmarshalFromEnviron(&conf)
	if err != nil {
		panic(errors.New(fmt.Sprintf("cannot load env vars properly: %v", err)))
	}

	// logger
	logger := zap.Must(zap.NewProduction())
	if conf.Environment != "production" {
		logger = zap.Must(zap.NewDevelopment())
	}
	defer logger.Sync()

	srv, err := app.NewServer(conf)
	if err != nil {
		logger.Fatal(err.Error())

	}
	if err := srv.Run(); err != nil {
		logger.Fatal("server couldn't start")
	}
}

package main

import (
	"log"

	authapp "github.com/OlegBogdanov06102017/auth-app"
	"github.com/OlegBogdanov06102017/auth-app/pkg/handler"
	"github.com/OlegBogdanov06102017/auth-app/pkg/repository"
	"github.com/OlegBogdanov06102017/auth-app/pkg/service"
)

func main() {
	repo := repository.NewRepository()
	services := service.NewService(repo)
	handlers := handler.NewHandler(services)

	srv := new(authapp.Server)
	if err := srv.Run("8000", handlers.InitRoutes()); err != nil {
		log.Fatalf("error occured while running http server: %s", err.Error())
	}

}

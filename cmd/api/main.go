package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/Safwanseban/voixme-project/configs"
	"github.com/Safwanseban/voixme-project/internal/db"
	"github.com/Safwanseban/voixme-project/internal/handlers"
	"github.com/Safwanseban/voixme-project/internal/repo"
	"github.com/Safwanseban/voixme-project/internal/usecases"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	config := configs.NewConfig()
	db := db.ConnectToDB(config)

	repo := repo.NewRepo(db)
	usecase := usecases.NewCompanyUsecase(repo)

	handlers.NewServer(app, usecase)
	usecase.FetchAndAppend()
	go func() {
		if err := app.Listen(config.String("port")); err != nil {
			panic("error initializing server")
		}
	}()

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)

	<-sigChan
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	defer func() {
		if err := app.ShutdownWithContext(ctx); err != nil {
			panic("graceful shutdown failed")
		} else {
			log.Println("gracefully closed the server")
		}
	}()

}

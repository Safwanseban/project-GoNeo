package main

import (
	"github.com/Safwanseban/voixme-project/configs"
	"github.com/Safwanseban/voixme-project/internal/db"
	"github.com/Safwanseban/voixme-project/internal/handlers"
	"github.com/Safwanseban/voixme-project/internal/repo"
	"github.com/Safwanseban/voixme-project/internal/usecases"
	"github.com/gofiber/fiber/v2"
)

func main() {
	config := configs.NewConfig()

	db := db.ConnectToDB(config)
	repo := repo.NewRepo(db)
	usecase := usecases.NewProductUsecase(repo)

	app := fiber.New()
	handlers.NewServer(app, usecase)
	usecase.FetchAndAppend()
	app.Listen(config.String("port"))

}

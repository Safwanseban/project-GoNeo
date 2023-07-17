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
	configs := configs.NewConfig()

	db := db.ConnectToDB(configs)
	repo := repo.NewRepo(db)
	usecase := usecases.NewProductUsecase(repo)

	app := fiber.New()
	handlers.NewServer(app, usecase)

	app.Listen(configs.String("port"))

}

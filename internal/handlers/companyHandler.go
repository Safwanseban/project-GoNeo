package handlers

import (
	"net/http"

	"github.com/Safwanseban/voixme-project/internal/types"
	"github.com/Safwanseban/voixme-project/internal/usecases"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

type Server struct {
	App            *fiber.App
	ProductUseCase usecases.UsecasesCompany
}

func NewServer(app *fiber.App, companyUsecase usecases.UsecasesCompany) {

	server := &Server{
		App:            app,
		ProductUseCase: companyUsecase,
	}
	server.App.Use(logger.New(logger.Config{
		Format: "[${ip}]:${port} ${status} - ${method} ${path}\n",
	}))
	server.App.Use(compress.New(compress.Config{
		Level: compress.LevelBestSpeed,
	}))
	server.App.Use(recover.New())
	server.App.Post("/", server.Create)
	server.App.Get("/", server.Fetch)
}

func (s *Server) Create(ctx *fiber.Ctx) error {
	company := new(types.OfferCompany)
	if err := ctx.BodyParser(company); err != nil {

		return ctx.Status(http.StatusBadRequest).JSON(
			fiber.Map{
				"message": "bad request",
			},
		)
	}
	id, err := s.ProductUseCase.CreateProduct(company)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"message": "server error",
		})

	}
	return ctx.Status(http.StatusOK).JSON(fiber.Map{"message": "success", "id": id})
}

func (s *Server) Fetch(ctx *fiber.Ctx) error {

	country := ctx.Query("country")
	if country == "" {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{
			"message": "provide valid country parameter",
		})
	}
	var company types.OfferCompany
	company.Country = country
	companies, err := s.ProductUseCase.ShowOfferCompany(&company)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"message": "error fetching the records",
		})
	}

	return ctx.Status(http.StatusOK).JSON(companies)

}

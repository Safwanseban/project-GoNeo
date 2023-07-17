package handlers

import (
	"net/http"

	"github.com/Safwanseban/voixme-project/internal/types"
	"github.com/Safwanseban/voixme-project/internal/usecases"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

type Server struct {
	App            *fiber.App
	ProductUseCase usecases.UsecasesProduct
}

func NewServer(app *fiber.App, productUsecase usecases.UsecasesProduct) {

	server := &Server{
		App:            app,
		ProductUseCase: productUsecase,
	}
	server.App.Use(logger.New(logger.Config{
		Format: "[${ip}]:${port} ${status} - ${method} ${path}\n",
	}))
	server.App.Post("", server.Create)
	server.App.Get("", server.Fetch)
}

func (s *Server) Create(ctx *fiber.Ctx) error {
	product := new(types.Product)
	if err := ctx.BodyParser(product); err != nil {

		return ctx.Status(http.StatusBadRequest).JSON(
			fiber.Map{
				"message": "bad request",
			},
		)
	}
	id, err := s.ProductUseCase.CreateProduct(product)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"message": "server error",
		})

	}
	return ctx.Status(http.StatusOK).JSON(fiber.Map{"message": "success", "id": id, "body": product})
}

func (s *Server) Fetch(ctx *fiber.Ctx) error {
	country := ctx.Query("country")
	if country == "" {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{
			"message": "provide valid country parameter",
		})
	}
	var product types.Product
	product.SpecificCountry = types.Country(country)
	products, err := s.ProductUseCase.ShowProducts(&product)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"message": "error fetching the records",
		})
	}
	return ctx.Status(http.StatusOK).JSON(products)

}

package handlers

import (
	"net/http"

	"github.com/Safwanseban/voixme-project/internal/types"
	"github.com/Safwanseban/voixme-project/internal/usecases"
	"github.com/gofiber/fiber/v2"
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
	app.Post("", server.Create)
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

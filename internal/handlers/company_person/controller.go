package customer

import (
	"github.com/gofiber/fiber/v2"
	"github.com/internet-banking-ul/internal/handlers"
	"github.com/internet-banking-ul/internal/modules/entities"
)

func (h *CompanyPersonHandlerImpl) CompanyPersonList(ctx *fiber.Ctx) error {
	baseFilter, err := entities.NewBaseFilterFromQuery(ctx)
	if err != nil {
		return ctx.Status(fiber.ErrBadRequest.Code).JSON(fiber.Map{"error": err.Error()})
	}

	customers, count, err := h.CompanyPersonService.List(ctx.Context(), *baseFilter)
	if err != nil {
		return ctx.Status(fiber.ErrInternalServerError.Code).JSON(fiber.Map{"error": err.Error()})
	}

	return ctx.Status(fiber.StatusOK).JSON(handlers.NewResponse(customers, count))
}

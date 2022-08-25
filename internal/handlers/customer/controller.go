package customer

import (
	"github.com/internet-banking-ul/internal/handlers"
	"github.com/internet-banking-ul/internal/modules/entities"
)

func (h *CustomerHandlerImpl) CustomerList(ctx *fiber.Ctx) error {
	baseFilter, err := entities.NewBaseFilterFromQuery(ctx)
	if err != nil {
		return ctx.Status(fiber.ErrBadRequest.Code).JSON(fiber.Map{"error": err.Error()})
	}

	customers, count, err := h.CustomerService.List(ctx.Context(), *baseFilter)
	if err != nil {
		return ctx.Status(fiber.ErrInternalServerError.Code).JSON(fiber.Map{"error": err.Error()})
	}

	return ctx.Status(fiber.StatusOK).JSON(handlers.NewResponse(customers, count))
}

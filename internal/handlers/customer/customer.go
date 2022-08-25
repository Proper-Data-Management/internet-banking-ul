package customer

import (
	"github.com/gofiber/fiber/v2"
	"github.com/mak-alex/al_hilal_core/internal/middles"
	customersvc "github.com/mak-alex/al_hilal_core/internal/modules/customer/services"
)

type CustomerHandlerImpl struct {
	customersvc.CustomerService
}

func NewCustomerHandler(
	customerService customersvc.CustomerService,
) *CustomerHandlerImpl {
	return &CustomerHandlerImpl{
		CustomerService: customerService,
	}
}

func (h *CustomerHandlerImpl) RegisterCustomer(r fiber.Router) {
	customerGroup := r.Group("customer")
	r.Use(
		middles.SetupContextHolder(),
		middles.SetupLanguage(),
		middles.SetupRequestInfo(),
		middles.NewFiberRecovery(middles.FiberRecoveryConfig{}),
	)
	{
		customerGroup.Get("", h.CustomerList)
	}
}

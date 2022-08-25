package customer

import (
	"github.com/gofiber/fiber/v2"
	"github.com/internet-banking-ul/internal/middles"
	companyPersonService "github.com/internet-banking-ul/internal/modules/company_person/services"
)

type CompanyPersonHandlerImpl struct {
	companyPersonService.CompanyPersonService
}

func NewCompanyPersonHandler(
	customerService companyPersonService.CompanyPersonService,
) *CompanyPersonHandlerImpl {
	return &CompanyPersonHandlerImpl{
		CompanyPersonService: customerService,
	}
}

func (h *CompanyPersonHandlerImpl) RegisterCompanyPerson(r fiber.Router) {
	customerGroup := r.Group("company_person")
	r.Use(
		middles.SetupContextHolder(),
		middles.SetupLanguage(),
		middles.SetupRequestInfo(),
		middles.NewFiberRecovery(middles.FiberRecoveryConfig{}),
	)
	{
		customerGroup.Get("", h.CompanyPersonList)
	}
}

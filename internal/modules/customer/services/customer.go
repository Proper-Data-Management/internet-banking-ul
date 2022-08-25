package services

import (
	"context"
	"database/sql"

	companyPersonRepo "github.com/mak-alex/al_hilal_core/internal/modules/company_person/repositories"
	"github.com/mak-alex/al_hilal_core/internal/modules/customer/dto"
	customerRepo "github.com/mak-alex/al_hilal_core/internal/modules/customer/repositories"
	"github.com/mak-alex/al_hilal_core/internal/modules/entities"
	"github.com/mak-alex/al_hilal_core/modules/logger"
)

type CustomerService interface {
	List(context.Context, entities.BasePaginationFilters) (dto.CustomerListResponse, int64, error)
}

type CustomerServiceImpl struct {
	CustomerRepository      customerRepo.Repositories
	CompanyPersonRepository companyPersonRepo.Repositories
}

func NewCustomerService(
	db *sql.DB,
) *CustomerServiceImpl {
	return &CustomerServiceImpl{
		CustomerRepository:      customerRepo.NewCustomerRepository(db),
		CompanyPersonRepository: companyPersonRepo.NewCompanyPersonRepository(db),
	}
}

func (s CustomerServiceImpl) List(ctx context.Context, baseFilter entities.BasePaginationFilters) (dto.CustomerListResponse, int64, error) {
	customerList, count, err := s.CustomerRepository.List(ctx, baseFilter)
	if err != nil {
		logger.WorkLoggerWithContext(ctx).Error("Error fetch CustomerList from DB")
		return nil, 0, err
	}

	if customerList == nil {
		return nil, 0, err
	}

	for i := 0; i < len(customerList); i++ {
		companyPerson, err := s.CompanyPersonRepository.ByCustomerID(ctx, customerList[i].ID)
		if err != nil {
			logger.WorkLoggerWithContext(ctx).Error("Error fetch CompanyPerson from DB")
			continue
		}

		customerList[i].CompanyPerson = companyPerson
	}

	return dto.CreateCustomerListResponse(customerList), count, nil
}

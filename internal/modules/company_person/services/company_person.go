package services

import (
	"context"
	"database/sql"

	"github.com/mak-alex/al_hilal_core/internal/modules/company_person/dto"
	companyPersonRepo "github.com/mak-alex/al_hilal_core/internal/modules/company_person/repositories"
	customerRepo "github.com/mak-alex/al_hilal_core/internal/modules/customer/repositories"
	"github.com/mak-alex/al_hilal_core/internal/modules/entities"
	"github.com/mak-alex/al_hilal_core/modules/logger"
)

type CompanyPersonService interface {
	List(context.Context, entities.BasePaginationFilters) (dto.CompanyPersonListResponse, int64, error)
}

type CompanyPersonServiceImpl struct {
	CompanyPersonRepository companyPersonRepo.Repositories
	CustomerRepository      customerRepo.Repositories
}

func NewCompanyPersonService(
	db *sql.DB,
) *CompanyPersonServiceImpl {
	return &CompanyPersonServiceImpl{
		CustomerRepository:      customerRepo.NewCustomerRepository(db),
		CompanyPersonRepository: companyPersonRepo.NewCompanyPersonRepository(db),
	}
}

func (s CompanyPersonServiceImpl) List(ctx context.Context, baseFilter entities.BasePaginationFilters) (dto.CompanyPersonListResponse, int64, error) {
	companyPersonList, count, err := s.CompanyPersonRepository.List(ctx, baseFilter)
	if err != nil {
		logger.WorkLoggerWithContext(ctx).Error("Error fetch CompanyPersonList from DB")
		return nil, 0, err
	}

	if companyPersonList == nil {
		return nil, 0, err
	}

	return dto.CreateCompanyPersonListResponse(companyPersonList), count, nil
}

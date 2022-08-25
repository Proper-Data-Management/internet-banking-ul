package dto

import (
	companyPersonDTO "github.com/mak-alex/al_hilal_core/internal/modules/company_person/dto"
	customerModel "github.com/mak-alex/al_hilal_core/internal/modules/customer/entities"
)

type CustomerResponse struct {
	ID                       int64  `json:"id"`
	PersonType               string `json:"personType"`
	ExternalID               string `json:"externalID"`
	Name                     string `json:"name"`
	FullName                 string `json:"fullName"`
	IntlName                 string `json:"intlName"`
	Ownership                string `json:"ownership"`
	ResidencyAndEconomicCode string `json:"residencyAndEconomicCode"`
	TaxCode                  string `json:"taxCode"`

	CompanyPerson companyPersonDTO.CompanyPersonResponse `json:"companyPerson,omitempty"`
}

func CreateCustomerResponse(
	customer customerModel.Customer,
	companyPersonDTO companyPersonDTO.CompanyPersonResponse,
) CustomerResponse {
	return CustomerResponse{
		ID:                       customer.ID,
		PersonType:               customer.PersonType,
		ExternalID:               customer.ExternalID,
		Name:                     customer.Name,
		FullName:                 customer.FullName,
		IntlName:                 customer.IntlName.ValueOrZero(),
		Ownership:                customer.Ownership,
		ResidencyAndEconomicCode: customer.ResidencyAndEconomicCode,
		TaxCode:                  customer.TaxCode,
		CompanyPerson:            companyPersonDTO,
	}
}

type CustomerListResponse []*CustomerResponse

func CreateCustomerListResponse(customers customerModel.CustomerList) CustomerListResponse {
	customersResp := CustomerListResponse{}
	for _, p := range customers {
		customer := CreateCustomerResponse(*p, companyPersonDTO.CreateCompanyPersonResponse(p.CompanyPerson))
		customersResp = append(customersResp, &customer)
	}
	return customersResp
}

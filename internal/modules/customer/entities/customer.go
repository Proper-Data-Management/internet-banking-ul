package entities

import (
	"github.com/guregu/null"
	companyPersonModel "github.com/mak-alex/al_hilal_core/internal/modules/company_person/entities"
)

type Customer struct {
	ID                       int64       `db:"ID" json:"id"`
	PersonType               string      `db:"PERSON_TYPE" json:"person_type"`
	ExternalID               string      `db:"EXTERNAL_ID" json:"external_id"`
	Name                     string      `db:"NAME" json:"name"`
	FullName                 string      `db:"FULL_NAME" json:"full_name"`
	IntlName                 null.String `db:"INTL_NAME" json:"intl_name"`
	Ownership                string      `db:"OWNERSHIP" json:"ownership"`
	ResidencyAndEconomicCode string      `db:"RESIDENCY_AND_ECONOMIC_CODE" json:"residency_and_economic_code"`
	TaxCode                  string      `db:"TAX_CODE" json:"tax_code"`

	CompanyPerson companyPersonModel.CompanyPerson `db:"-" json:"company_person"`
}

type CustomerList []*Customer

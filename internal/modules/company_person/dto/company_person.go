package dto

import (
	"database/sql"
	"time"

	"github.com/internet-banking-ul/internal/modules/company_person/entities"
)

type CompanyPersonResponse struct {
	ID               int64         `json:"id"`
	IsDeleted        int           `json:"isDeleted"`
	ExternalID       string        `json:"externalID"`
	CompanyID        int64         `json:"companyID"`
	UserAccountID    int64         `json:"userAccountID"`
	ManagerID        sql.NullInt64 `json:"managerID"`
	ValidFrom        time.Time     `json:"validFrom"`
	ValidTo          time.Time     `json:"validTo"`
	SignLevel        string        `json:"signLevel"`
	OrganizationRole string        `json:"organizationRole"`
}

func CreateCompanyPersonResponse(companyPerson entities.CompanyPerson) CompanyPersonResponse {
	return CompanyPersonResponse{
		ID:               companyPerson.ID,
		IsDeleted:        companyPerson.IsDeleted,
		ExternalID:       companyPerson.ExternalID,
		CompanyID:        companyPerson.CompanyID,
		UserAccountID:    companyPerson.UserAccountID,
		ManagerID:        companyPerson.ManagerID,
		ValidFrom:        companyPerson.ValidFrom.Time,
		ValidTo:          companyPerson.ValidTo.Time,
		SignLevel:        companyPerson.SignLevel,
		OrganizationRole: companyPerson.OrganizationRole,
	}
}

type CompanyPersonListResponse []*CompanyPersonResponse

func CreateCompanyPersonListResponse(companyPersonList entities.CompanyPersonList) CompanyPersonListResponse {
	companyPersonListResp := CompanyPersonListResponse{}
	for _, p := range companyPersonList {
		companyPerson := CreateCompanyPersonResponse(*p)
		companyPersonListResp = append(companyPersonListResp, &companyPerson)
	}
	return companyPersonListResp
}

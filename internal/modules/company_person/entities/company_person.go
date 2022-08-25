package entities

import (
	"database/sql"
)

type CompanyPerson struct {
	ID               int64         `db:"ID" json:"id"`
	IsDeleted        int           `db:"IS_DELETED" json:"is_deleted"`
	ExternalID       string        `db:"EXTERNAL_ID" json:"external_id"`
	CompanyID        int64         `db:"COMPANY_ID" json:"company_id"`
	UserAccountID    int64         `db:"USER_ACCOUNT_ID" json:"user_account_id"`
	ManagerID        sql.NullInt64 `db:"MANAGER_ID" json:"manager_id"`
	ValidFrom        sql.NullTime  `db:"VALID_FROM" json:"valid_from"`
	ValidTo          sql.NullTime  `db:"VALID_TO" json:"valid_to"`
	SignLevel        string        `db:"SIGN_LEVEL" json:"sign_level"`
	OrganizationRole string        `db:"ORGANIZATION_ROLE" json:"organization_role"`
}

type CompanyPersonList []*CompanyPerson

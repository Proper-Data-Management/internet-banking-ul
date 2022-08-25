package repositories

import (
	"context"
	"database/sql"
	"fmt"

	companyPersonModel "github.com/mak-alex/al_hilal_core/internal/modules/company_person/entities"
	"github.com/mak-alex/al_hilal_core/internal/modules/entities"
	"github.com/mak-alex/al_hilal_core/modules/logger"
	sq "github.com/mak-alex/al_hilal_core/modules/squirrel"
	"go.uber.org/zap"
)

type RepositoryCompanyPersonQuery interface {
	Count(ctx context.Context) (count int64, err error)
	List(context.Context, entities.BasePaginationFilters) (companyPersonModel.CompanyPersonList, int64, error)
	ByCustomerID(ctx context.Context, customerID int64) (result companyPersonModel.CompanyPerson, err error)
}

type RepositoryCompanyPersonQueryImpl struct {
	DB *sql.DB
}

func (repo *RepositoryCompanyPersonQueryImpl) ByCustomerID(ctx context.Context, customerID int64) (result companyPersonModel.CompanyPerson, err error) {
	if repo.DB == nil {
		err = fmt.Errorf("db is nil")
		return result, err
	}

	l := logger.WorkLoggerWithContext(ctx).Named("List")

	q := sq.
		Select([]string{
			"ID",
			"IS_DELETED",
			"EXTERNAL_ID",
			"COMPANY_ID",
			"USER_ACCOUNT_ID",
			"MANAGER_ID",
			"VALID_FROM",
			"VALID_TO",
			"SIGN_LEVEL",
			"ORGANIZATION_ROLE",
		}...).
		From("COMPANY_PERSON").
		Where(sq.Eq{"COMPANY_ID": customerID}).
		PlaceholderFormat(sq.Colon)

	sql, args, e := q.ToSql()
	if e != nil {
		l.Error("ToSql", zap.Error(e))
		return result, err
	}

	l.Debug("Info", zap.String("sql", sql), zap.Any("args", args))

	err = q.RunWith(repo.DB).QueryRowContext(ctx).Scan(
		&result.ID,
		&result.IsDeleted,
		&result.ExternalID,
		&result.CompanyID,
		&result.UserAccountID,
		&result.ManagerID,
		&result.ValidFrom,
		&result.ValidTo,
		&result.SignLevel,
		&result.OrganizationRole,
	)
	if err != nil {
		l.Error("QueryRowContext", zap.Error(err))
		return
	}

	return
}

func (repo *RepositoryCompanyPersonQueryImpl) List(ctx context.Context, baseFilter entities.BasePaginationFilters) (results companyPersonModel.CompanyPersonList, count int64, err error) {
	if repo.DB == nil {
		err = fmt.Errorf("db is nil")
		return results, count, err
	}

	l := logger.WorkLoggerWithContext(ctx).Named("List")

	count, err = repo.Count(ctx)
	if err != nil {
		l.Error("Count", zap.Error(err))
		return results, count, err
	}

	q := sq.
		Select([]string{
			"ID",
			"IS_DELETED",
			"EXTERNAL_ID",
			"COMPANY_ID",
			"USER_ACCOUNT_ID",
			"MANAGER_ID",
			"VALID_FROM",
			"VALID_TO",
			"SIGN_LEVEL",
			"ORGANIZATION_ROLE",
		}...).
		From("COMPANY_PERSON").
		Offset(baseFilter.GetOffset()).
		Limit(baseFilter.GetSize()).
		PlaceholderFormat(sq.Colon)

	sql, args, e := q.ToSql()
	if e != nil {
		l.Error("ToSql", zap.Error(e))
		return results, count, err
	}

	l.Debug("Info", zap.String("sql", sql), zap.Any("args", args))

	rows, err := q.RunWith(repo.DB).QueryContext(ctx)
	if err != nil {
		l.Error("QueryContext", zap.Error(err))
		return results, count, err
	}
	defer rows.Close()

	results = companyPersonModel.CompanyPersonList{}
	for rows.Next() {
		row := new(companyPersonModel.CompanyPerson)
		if err := rows.Scan(
			&row.ID,
			&row.IsDeleted,
			&row.ExternalID,
			&row.CompanyID,
			&row.UserAccountID,
			&row.ManagerID,
			&row.ValidFrom,
			&row.ValidTo,
			&row.SignLevel,
			&row.OrganizationRole,
		); err != nil {
			l.Error("Scan", zap.Error(err))
			return results, count, err
		}

		results = append(results, row)
	}

	if err := rows.Close(); err != nil {
		return results, count, err
	}

	if err := rows.Err(); err != nil {
		return results, count, err
	}

	return results, count, err
}

func (repo *RepositoryCompanyPersonQueryImpl) Count(ctx context.Context) (count int64, err error) {
	if repo.DB == nil {
		err = fmt.Errorf("db is nil")
		return
	}

	l := logger.WorkLoggerWithContext(ctx).Named("Count")

	q := sq.Select("COUNT(1)").From("COMPANY_PERSON").PlaceholderFormat(sq.Colon)

	sql, args, e := q.ToSql()
	if e != nil {
		l.Error("ToSql", zap.Error(e))
		return
	}

	l.Debug("Info", zap.String("sql", sql), zap.Any("args", args))

	err = q.RunWith(repo.DB).QueryRowContext(ctx).Scan(&count)
	if err != nil {
		l.Error("QueryRowContext", zap.Error(err))
		return
	}

	return
}

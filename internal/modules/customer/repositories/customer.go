package repositories

import (
	"context"
	"database/sql"
	"fmt"

	customerModel "github.com/mak-alex/al_hilal_core/internal/modules/customer/entities"
	"github.com/mak-alex/al_hilal_core/internal/modules/entities"
	"github.com/mak-alex/al_hilal_core/modules/logger"
	sq "github.com/mak-alex/al_hilal_core/modules/squirrel"
	"go.uber.org/zap"
)

type RepositoryCustomerQuery interface {
	List(context.Context, entities.BasePaginationFilters) (customerModel.CustomerList, int64, error)
}

type RepositoryCustomerQueryImpl struct {
	DB *sql.DB
}

func (repo *RepositoryCustomerQueryImpl) List(ctx context.Context, baseFilter entities.BasePaginationFilters) (results customerModel.CustomerList, count int64, err error) {
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
			"PERSON_TYPE",
			"EXTERNAL_ID",
			"NAME",
			"FULL_NAME",
			"INTL_NAME",
			"OWNERSHIP",
			"RESIDENCY_AND_ECONOMIC_CODE",
			"TAX_CODE",
		}...).
		From("CUSTOMER").
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

	results = customerModel.CustomerList{}
	for rows.Next() {
		row := new(customerModel.Customer)
		if err := rows.Scan(
			&row.ID,
			&row.PersonType,
			&row.ExternalID,
			&row.Name,
			&row.FullName,
			&row.IntlName,
			&row.Ownership,
			&row.ResidencyAndEconomicCode,
			&row.TaxCode,
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

func (repo *RepositoryCustomerQueryImpl) Count(ctx context.Context) (count int64, err error) {
	if repo.DB == nil {
		err = fmt.Errorf("db is nil")
		return
	}

	l := logger.WorkLoggerWithContext(ctx).Named("Count")

	q := sq.Select("COUNT(1)").From("CUSTOMER").PlaceholderFormat(sq.Colon)

	sql, args, e := q.ToSql()
	if e != nil {
		l.Error("ToSql", zap.Error(e))
		return
	}

	l.Debug("Info", zap.String("sql", sql), zap.Any("args", args))

	err = q.RunWith(repo.DB).QueryRowContext(ctx).Scan(&count)
	if err != nil {
		l.Error("QueryContext", zap.Error(err))
		return
	}

	return
}

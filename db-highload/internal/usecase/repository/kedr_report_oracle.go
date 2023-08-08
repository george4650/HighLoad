package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"gitlabnew.nextcontact.ru/nextcontactcenter/services/highload-testers/db-highload-tester/internal/usecase"
	oracle "gitlabnew.nextcontact.ru/r.alfimov/go-oracle"
)

type kedrReportOracleRepository struct {
	*oracle.Oracle
}

func NewKedrReportOracle(ora *oracle.Oracle) usecase.KedrReportRepository {
	return &kedrReportOracleRepository{ora}
}

func (r *kedrReportOracleRepository) LoadEmployeeTimesheet(ctx context.Context, dateStart, dateEnd string) error {
	sqlText := `select 1 as done
                from table(kedr.fnc_report_employee_timesheet (
                              i_date_start => :dateStart
                            , i_date_end => :dateEnd))`
	type res struct {
		Done int `db:"done"`
	}
	if _, err := oracle.SelectMany[res](ctx, r.Oracle, sqlText, []interface{}{dateStart, dateEnd}); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil
		}
		return fmt.Errorf("oracle.SelectMany: %w", err)
	}

	return nil
}

func (r *kedrReportOracleRepository) LoadAPIStatuses(ctx context.Context, dateStart, dateEnd string) error {
	sqlText := `select 1 as done
				from table(kedr.fnc_api_statuses (
							  i_date_start => :dateStart
							, i_date_end => :dateEnd))`
	type res struct {
		Done int `db:"done"`
	}
	if _, err := oracle.SelectMany[res](ctx, r.Oracle, sqlText, []interface{}{dateStart, dateEnd}); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil
		}
		return fmt.Errorf("oracle.SelectMany: %w", err)
	}

	return nil
}

func (r *kedrReportOracleRepository) LoadDetailCalls(ctx context.Context, dateStart, dateEnd string) error {
	sqlText := `select 1 as done
				from table(kedr.FNC_REPORT_DETAIL_CALLS (
							  i_date_start => :dateStart
							, i_date_end => :dateEnd))`
	type res struct {
		Done int `db:"done"`
	}
	if _, err := oracle.SelectMany[res](ctx, r.Oracle, sqlText, []interface{}{dateStart, dateEnd}); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil
		}
		return fmt.Errorf("oracle.SelectMany: %w", err)
	}

	return nil
}

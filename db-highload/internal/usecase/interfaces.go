package usecase

import (
	"context"
)

type (
	Highload interface {
		GetEmployeeTimesheet(ctx context.Context, dateStart, dateEnd string) error
		GetAPIStatuses(ctx context.Context, dateStart, dateEnd string) error
		GetDetailCalls(ctx context.Context, dateStart, dateEnd string) error
	}

	KedrReportRepository interface {
		LoadEmployeeTimesheet(ctx context.Context, dateStart, dateEnd string) error
		LoadAPIStatuses(ctx context.Context, dateStart, dateEnd string) error
		LoadDetailCalls(ctx context.Context, dateStart, dateEnd string) error
	}
)

package usecase

import (
	"context"
	"fmt"
)

type highloadUseCase struct {
	repo KedrReportRepository
}

func NewHighload(repo KedrReportRepository) Highload {
	return &highloadUseCase{repo: repo}
}

func (uc *highloadUseCase) GetEmployeeTimesheet(ctx context.Context, dateStart, dateEnd string) error {
	if err := uc.repo.LoadEmployeeTimesheet(ctx, dateStart, dateEnd); err != nil {
		return fmt.Errorf("highloadUseCase - GetEmployeeTimesheet - uc.repo.LoadEmployeeTimesheet: %w", err)
	}

	return nil
}

func (uc *highloadUseCase) GetAPIStatuses(ctx context.Context, dateStart, dateEnd string) error {
	if err := uc.repo.LoadAPIStatuses(ctx, dateStart, dateEnd); err != nil {
		return fmt.Errorf("highloadUseCase - GetAPIStatuses - uc.repo.LoadAPIStatuses: %w", err)
	}

	return nil
}

func (uc *highloadUseCase) GetDetailCalls(ctx context.Context, dateStart, dateEnd string) error {
	if err := uc.repo.LoadDetailCalls(ctx, dateStart, dateEnd); err != nil {
		return fmt.Errorf("highloadUseCase - GetDetailCalls - uc.repo.LoadDetailCalls: %w", err)
	}

	return nil
}

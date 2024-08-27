package logging

import (
	"LoanTrackerApi/internal/entity"
	"context"
)

type Logger interface {
	AddLog(ctx context.Context, category, message string) error
	GetLogs(ctx context.Context) ([]entity.Log, error)
}

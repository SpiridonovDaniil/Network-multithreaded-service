package repository

import (
	"context"
	"diploma/internal/domain"
)

type Parser interface {
	ParseData(ctx context.Context) ([]domain.EmailData, error)
}

package tabledm

import (
	"context"
)

type TableRepository interface {
	CreateTable(ctx context.Context, name string) (*Table, error)
	FindTableByName(ctx context.Context, name string) (*Table, error)
	DeleteTable(ctx context.Context, name string) error
}

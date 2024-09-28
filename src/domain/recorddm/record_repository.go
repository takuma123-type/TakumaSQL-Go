package recorddm

import (
	"context"
)

// RecordRepository インターフェース
type RecordRepository interface {
	Save(ctx context.Context, record *Record) error
	FindByField(ctx context.Context, fieldName, value string) (*Record, error)
}

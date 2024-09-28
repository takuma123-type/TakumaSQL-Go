package repository

import (
	"context"

	"github.com/takuma123-type/TakumaSQL-Go/src/domain/recorddm"
)

type InMemoryRecordRepository struct {
	records []recorddm.Record
}

func NewInMemoryRecordRepository() *InMemoryRecordRepository {
	return &InMemoryRecordRepository{
		records: []recorddm.Record{},
	}
}

func (r *InMemoryRecordRepository) Save(ctx context.Context, record *recorddm.Record) error {
	r.records = append(r.records, *record)
	return nil
}

func (r *InMemoryRecordRepository) FindByField(ctx context.Context, fieldName, value string) (*recorddm.Record, error) {
	for _, rec := range r.records {
		if v, ok := rec.Fields[fieldName]; ok && v == value {
			return &rec, nil
		}
	}
	return nil, nil
}

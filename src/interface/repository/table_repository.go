package repository

import (
	"context"

	"github.com/takuma123-type/TakumaSQL-Go/src/domain/tabledm"
)

type InMemoryTableRepository struct {
	tables map[string]*tabledm.Table
}

func NewInMemoryTableRepository() *InMemoryTableRepository {
	return &InMemoryTableRepository{
		tables: make(map[string]*tabledm.Table),
	}
}

// CreateTable メソッド
func (r *InMemoryTableRepository) CreateTable(ctx context.Context, name string) (*tabledm.Table, error) {
	table := tabledm.NewTable(name)
	r.tables[name] = table
	return table, nil
}

// FindTableByName メソッドを追加
func (r *InMemoryTableRepository) FindTableByName(ctx context.Context, name string) (*tabledm.Table, error) {
	table, exists := r.tables[name]
	if !exists {
		return nil, nil
	}
	return table, nil
}

// DeleteTable メソッド
func (r *InMemoryTableRepository) DeleteTable(ctx context.Context, name string) error {
	delete(r.tables, name)
	return nil
}

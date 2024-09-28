package tableusecase

import (
	"context"
	"errors"

	"github.com/takuma123-type/TakumaSQL-Go/src/domain/tabledm"
)

type CreateTableUsecase struct {
	repository tabledm.TableRepository
}

func NewCreateTableUsecase(repo tabledm.TableRepository) *CreateTableUsecase {
	return &CreateTableUsecase{
		repository: repo,
	}
}

func (uc *CreateTableUsecase) Execute(ctx context.Context, name string) (*tabledm.Table, error) {
	// テーブルが既に存在しているか確認
	existingTable, err := uc.repository.FindTableByName(ctx, name)
	if err != nil {
		return nil, err
	}
	if existingTable != nil {
		return nil, errors.New("table already exists")
	}

	// テーブルを作成
	return uc.repository.CreateTable(ctx, name)
}

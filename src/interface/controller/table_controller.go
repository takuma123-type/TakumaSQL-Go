package controller

import (
	"context"

	"github.com/takuma123-type/TakumaSQL-Go/src/usecase/tableusecase"
	"github.com/takuma123-type/TakumaSQL-Go/src/usecase/tableusecase/tableinput"
	"github.com/takuma123-type/TakumaSQL-Go/src/usecase/tableusecase/tableoutput"
)

type TableController struct {
	createUsecase *tableusecase.CreateTableUsecase
}

func NewTableController(usecase *tableusecase.CreateTableUsecase) *TableController {
	return &TableController{
		createUsecase: usecase,
	}
}

func (c *TableController) CreateTable(ctx context.Context, in *tableinput.CreateTableInput, p tableoutput.TablePresenter) error {
	// in からテーブル名を取得して Execute メソッドに渡す
	table, err := c.createUsecase.Execute(ctx, in.Name)
	if err != nil {
		return err
	}

	// テーブル作成結果をプレゼンターに渡す
	output := &tableoutput.CreateTableOutput{
		Name: table.Name,
	}
	p.PresentCreateTable(output)
	return nil
}

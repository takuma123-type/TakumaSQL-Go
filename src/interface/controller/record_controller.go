package controller

import (
	"context"

	"github.com/takuma123-type/TakumaSQL-Go/src/usecase/recordusecase"
	"github.com/takuma123-type/TakumaSQL-Go/src/usecase/recordusecase/recordinput"
	"github.com/takuma123-type/TakumaSQL-Go/src/usecase/recordusecase/recordoutput"
)

type RecordController struct {
	createUsecase *recordusecase.CreateRecordUsecase
}

func NewRecordController(usecase *recordusecase.CreateRecordUsecase) *RecordController {
	return &RecordController{
		createUsecase: usecase,
	}
}

func (c *RecordController) CreateRecord(ctx context.Context, in *recordinput.CreateRecordInput) (*recordoutput.CreateRecordOutput, error) {
	return c.createUsecase.Execute(ctx, in)
}

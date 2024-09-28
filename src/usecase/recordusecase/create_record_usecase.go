package recordusecase

import (
	"context"

	"github.com/takuma123-type/TakumaSQL-Go/src/domain/recorddm"
	"github.com/takuma123-type/TakumaSQL-Go/src/usecase/recordusecase/recordinput"
	"github.com/takuma123-type/TakumaSQL-Go/src/usecase/recordusecase/recordoutput"
)

type CreateRecordUsecase struct {
	repository recorddm.RecordRepository
}

func NewCreateRecordUsecase(repository recorddm.RecordRepository) *CreateRecordUsecase {
	return &CreateRecordUsecase{repository: repository}
}

func (use *CreateRecordUsecase) Execute(ctx context.Context, in *recordinput.CreateRecordInput) (*recordoutput.CreateRecordOutput, error) {
	record := recorddm.NewRecord(in.Fields)

	if err := use.repository.Save(ctx, record); err != nil {
		return &recordoutput.CreateRecordOutput{
			Success: false,
			Message: "Failed to create record",
		}, err
	}

	return &recordoutput.CreateRecordOutput{
		Success: true,
		Message: "Record created successfully",
	}, nil
}

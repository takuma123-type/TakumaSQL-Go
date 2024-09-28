package presenter

import (
	"net/http"

	"github.com/takuma123-type/TakumaSQL-Go/src/usecase/recordusecase/recordoutput"

	"github.com/gin-gonic/gin"
)

type RecordPresenter struct {
	ctx *gin.Context
}

func NewRecordPresenter(ctx *gin.Context) *RecordPresenter {
	return &RecordPresenter{ctx: ctx}
}

func (p *RecordPresenter) PresentCreate(output *recordoutput.CreateRecordOutput) {
	if output.Success {
		p.ctx.JSON(http.StatusCreated, gin.H{
			"message": output.Message,
		})
	} else {
		p.ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": output.Message,
		})
	}
}

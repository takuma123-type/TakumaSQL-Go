package presenter

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/takuma123-type/TakumaSQL-Go/src/usecase/tableusecase/tableoutput"
)

type TablePresenter struct {
	ctx *gin.Context
}

func NewTablePresenter(ctx *gin.Context) *TablePresenter {
	return &TablePresenter{ctx}
}

func (p *TablePresenter) PresentCreateTable(output *tableoutput.CreateTableOutput) {
	p.ctx.JSON(http.StatusCreated, output)
}

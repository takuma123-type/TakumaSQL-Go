package router

import (
	"github.com/gin-gonic/gin"
	"github.com/takuma123-type/TakumaSQL-Go/src/interface/controller"
	"github.com/takuma123-type/TakumaSQL-Go/src/interface/presenter"
	"github.com/takuma123-type/TakumaSQL-Go/src/interface/repository"
	"github.com/takuma123-type/TakumaSQL-Go/src/usecase/tableusecase"
	"github.com/takuma123-type/TakumaSQL-Go/src/usecase/tableusecase/tableinput"
)

func NewTableRouter(r *gin.Engine) {
	api := r.Group("/api")
	{
		api.POST("/tables", func(ctx *gin.Context) {
			var in tableinput.CreateTableInput
			if err := ctx.ShouldBindJSON(&in); err != nil {
				ctx.JSON(400, gin.H{"error": "Invalid input"})
				return
			}

			// リポジトリ、ユースケース、コントローラーの初期化
			tableRepo := repository.NewInMemoryTableRepository()
			createUsecase := tableusecase.NewCreateTableUsecase(tableRepo)
			tableController := controller.NewTableController(createUsecase)

			// プレゼンターを使用して結果を返す
			p := presenter.NewTablePresenter(ctx)
			tableController.CreateTable(ctx.Request.Context(), &in, p)
		})
	}
}

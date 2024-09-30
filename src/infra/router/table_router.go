package router

import (
	"github.com/gin-gonic/gin"
	"github.com/takuma123-type/TakumaSQL-Go/src/interface/controller"
	"github.com/takuma123-type/TakumaSQL-Go/src/interface/presenter"
	"github.com/takuma123-type/TakumaSQL-Go/src/interface/repository"
	"github.com/takuma123-type/TakumaSQL-Go/src/usecase/tableusecase"
	"github.com/takuma123-type/TakumaSQL-Go/src/usecase/tableusecase/tableinput"
)

// src/infra/router/table_router.go
func NewTableRouter(r *gin.Engine) {
	api := r.Group("/api")
	{
		api.POST("/tables", func(ctx *gin.Context) {
			var in tableinput.CreateTableInput
			if err := ctx.ShouldBindJSON(&in); err != nil {
				ctx.JSON(400, gin.H{"error": "Invalid input"})
				return
			}

			tableRepo := repository.NewFileTableRepository("src/storage/tables.json")
			createUsecase := tableusecase.NewCreateTableUsecase(tableRepo)
			tableController := controller.NewTableController(createUsecase)

			p := presenter.NewTablePresenter(ctx)
			err := tableController.CreateTable(ctx.Request.Context(), &in, p)
			if err != nil {
				ctx.JSON(500, gin.H{"error": err.Error()})
				return
			}
			ctx.JSON(201, gin.H{"message": "Table created successfully"})
		})

		// 新しいエンドポイント：すべてのテーブルを取得する
		api.GET("/tables", func(ctx *gin.Context) {
			tableRepo := repository.NewFileTableRepository("src/storage/tables.json")
			tables := tableRepo.GetAllTables(ctx.Request.Context())
			ctx.JSON(200, gin.H{"tables": tables})
		})
	}
}

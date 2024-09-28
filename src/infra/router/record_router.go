package router

import (
	"github.com/gin-gonic/gin"
	"github.com/takuma123-type/TakumaSQL-Go/src/interface/controller"
	"github.com/takuma123-type/TakumaSQL-Go/src/interface/repository"
	"github.com/takuma123-type/TakumaSQL-Go/src/usecase/recordusecase"
	"github.com/takuma123-type/TakumaSQL-Go/src/usecase/recordusecase/recordinput"
)

func NewRecordRouter(r *gin.Engine) {
	api := r.Group("/api")
	{
		api.POST("/records", func(ctx *gin.Context) {
			var in recordinput.CreateRecordInput
			if err := ctx.ShouldBindJSON(&in); err != nil {
				ctx.JSON(400, gin.H{"error": "Invalid input"})
				return
			}

			// リポジトリ、ユースケース、コントローラーの初期化
			recordRepo := repository.NewInMemoryRecordRepository()
			createUsecase := recordusecase.NewCreateRecordUsecase(recordRepo)
			recordController := controller.NewRecordController(createUsecase)

			// プレゼンターを使わず、結果を直接返す
			output, err := recordController.CreateRecord(ctx.Request.Context(), &in)
			if err != nil {
				ctx.JSON(500, gin.H{"error": err.Error()})
				return
			}
			ctx.JSON(201, output)
		})
	}
}

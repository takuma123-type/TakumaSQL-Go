package main

import (
	"github.com/gin-gonic/gin"
	"github.com/takuma123-type/TakumaSQL-Go/src/infra/router"
)

func main() {
	r := gin.Default()
	router.NewRecordRouter(r)
	router.NewTableRouter(r)

	r.Run(":9090")
}

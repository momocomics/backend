package rest

import (
	"github.com/gin-gonic/gin"
	"github.com/momocomics/backend/pkg/api"
	"github.com/momocomics/backend/pkg/config"
)

func Routes(e *gin.Engine, sc *config.ServerConfig) {

	v1 := e.Group("/api/v1")

	v1.GET("/book/:id", api.GetBookFn(sc))

}

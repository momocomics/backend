package routes

import (
	"github.com/gin-gonic/gin"

	"github.com/momocomics/backend/http-server/pkg/api"
	"github.com/momocomics/backend/http-server/pkg/config"
)

func Routes(e *gin.Engine, sc *config.ServerConfig) {

	v1 := e.Group("/api/v1")
	v1.GET("/todo/:category", api.ListFn(sc))
	v1.POST("/todo/add", api.AddFn(sc))
	v1.POST("/signin", api.SigninFn(sc))
	v1.POST("/refresh_token", api.RefreshTokenFn(sc))
}

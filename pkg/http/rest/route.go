package rest

import (
	"github.com/gin-gonic/gin"
	"github.com/momocomics/backend/pkg/get"
	"github.com/momocomics/backend/pkg/storage"
)

func Routes(e *gin.Engine, db storage.DB) {

	v1 := e.Group("/api/v1")

	get.GetBook(v1, db)

}

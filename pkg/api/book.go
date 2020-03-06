package api

import (
	"context"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/momocomics/backend/pkg/config"
	"github.com/momocomics/backend/pkg/entity"
)

func GetBookFn(cfg *config.ServerConfig) func(*gin.Context) {
	return func(c *gin.Context) {

		ctx := context.Background()

		id := c.Param("id")
		rc, err := cfg.Db().Get(ctx, id)
		if err != nil {
			c.AbortWithError(http.StatusNotFound, err)
			return
		}
		defer rc.Close()

		data, err := ioutil.ReadAll(rc)
		if err != nil {
			c.AbortWithError(http.StatusInternalServerError, err)
			return
		}

		c.JSON(http.StatusOK, entity.Book{
			ID:      id,
			Title:   id,
			Content: string(data),
		})

	}
}

package get

import (
	"context"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/momocomics/backend/pkg/dto"
	"github.com/momocomics/backend/pkg/storage"
)

func GetBook(rg *gin.RouterGroup, db storage.DB) {
	rg.GET("/book/:id", func(c *gin.Context) {

		ctx := context.Background()

		id := c.Param("id")
		rc, err := db.Get(ctx, id)
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

		c.JSON(http.StatusOK, dto.Book{
			ID:      id,
			Title:   id,
			Content: string(data),
		})

	})
}

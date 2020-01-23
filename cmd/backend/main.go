package main

import (
	"context"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/momocomics/backend/pkg/http/rest"
	storage "github.com/momocomics/backend/pkg/storage/nosql"
)

func main() {
	ctx := context.Background()
	db, err := storage.NewGcs(ctx, "", "gcore")

	if err != nil {
		log.Fatal(err)
	}
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Backend running.")
	})
	rest.Routes(r, db)

	//http.ListenAndServe("8080", gin)
	r.Run(":8080")
}

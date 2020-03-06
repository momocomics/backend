package main

import (
	"context"
	"log"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/momocomics/backend/pkg/config"
	"github.com/momocomics/backend/pkg/http/rest"
	storage "github.com/momocomics/backend/pkg/storage/nosql"
)

//TODO: this is a test key. Create prod key and add it to k8s secret
// generate pk - openssl genrsa -out app.rsa 512(keysize)
// generate pub key - openssl rsa -in app.rsa -pubout > app.rsa.pub
var signBytes = []byte(`-----BEGIN RSA PRIVATE KEY-----
MIIBPAIBAAJBAMexg5p7etfjyJUO2Oz8vw5rxm0MC3wNsNEi1wInXGRwnUFWkQjN
6Jc0Y+t+eSYcvTpusprEecnJOKvrf+b7URcCAwEAAQJBAIFU0aQipuvdxdHsHMhX
5TFk0c1cSK/eeg7o3pGxhmAxfZLSQr73vP54Bk7xaaatOqqjCTUN0nixbEOeh18d
7DECIQD0fh2cf0xzQ07K3Fo6b1KAnjL4gR7XAJ/3rlXDAdUyCQIhANEXnSmU0tbf
CiYsqqloZTpZbpop6dEjt1vAWA++YLIfAiEAupvLtBgBXPRhnjpDb9hp6xtUIhJD
XKzwa9YXRUkP1SkCIQCnorYHQ2Eykklxx7ff8GnQOSlagiYK3gbAkdpIbQrbYwIg
STU47xTbB232EokQn4/ATVNXuYpRuFcLlIYeBNaB0aA=
-----END RSA PRIVATE KEY-----`)

func main() {
	ctx := context.Background()
	db, err := storage.NewGcs(ctx, "", "gcore")
	fatal(err)
	key, err := jwt.ParseRSAPrivateKeyFromPEM(signBytes)
	fatal(err)
	cfg := config.New(db, key)
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Backend running.")
	})
	rest.Routes(r, cfg)

	//http.ListenAndServe("8080", gin)
	r.Run(":8081")
}

func fatal(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

package main

import (
	"context"
	"log"
	"net/http"
	"regexp"

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
                        MIIBOgIBAAJBAJVKVZSF/8PIzJ1uoVrorfcjaULg4ZCtSHeTB6b9QnLJ/x6FXNBu
                        DxOR5dFhqiHpPtNhpS9mqQT9V3I6qXfGgzcCAwEAAQJAXsPl2TbKKPyQriqosC1d
                        KLDIw5Q+evkUNBsX02+WO4h3gCZeGouhGebaMgno1pUzsSal0W9U9XWnumWnkGwr
                        gQIhAMR9b00nFTjLMvRoi0R87nUNRgdhk431wfVmJopYU0VBAiEAwoFYoVz4Dckp
                        csRfg7bsb6HfcIlxncvNV2aleN/C0ncCICE/4J+7p1mu+PZm4no6cdeY4WrKVj/F
                        gIbYPFlYzO6BAiEAtb5Qx65sJc2CijeNnDBvarvRYYE8BZrqSzGhinlivG8CIG9r
                        CwQbSEUWg/snnrIOJH4O89fY9c/OZW48MLzsdyrh
                        -----END RSA PRIVATE KEY-----`)
var verifyBytes = []byte(`-----BEGIN PUBLIC KEY-----
                          MFwwDQYJKoZIhvcNAQEBBQADSwAwSAJBAJVKVZSF/8PIzJ1uoVrorfcjaULg4ZCt
                          SHeTB6b9QnLJ/x6FXNBuDxOR5dFhqiHpPtNhpS9mqQT9V3I6qXfGgzcCAwEAAQ==
                          -----END PUBLIC KEY-----`)

func main() {
	ctx := context.Background()
	db, err := storage.NewGcs(ctx, "", "gcore")
	fatal(err)
	//if s, ok := signBytes.([]byte);!ok{
	//
	//}
	signBytes = regexp.MustCompile(`\s*$`).ReplaceAll(signBytes, []byte{})
	privKey, err := jwt.ParseRSAPrivateKeyFromPEM(signBytes)
	fatal(err)
	verifyBytes = regexp.MustCompile(`\s*$`).ReplaceAll(verifyBytes, []byte{})
	pubkey, err := jwt.ParseRSAPublicKeyFromPEM(verifyBytes)
	fatal(err)
	cfg := config.New(db, privKey, pubkey, "momocomic.com")
	cfg.SetDebug(true)
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

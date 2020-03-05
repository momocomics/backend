package api

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/momocomics/backend/pkg/config"
)

//TODO: this is a test key. Create prod key and add it to k8s secret
var srt = []byte(`-----BEGIN RSA PRIVATE KEY-----
MIIBPAIBAAJBAMexg5p7etfjyJUO2Oz8vw5rxm0MC3wNsNEi1wInXGRwnUFWkQjN
6Jc0Y+t+eSYcvTpusprEecnJOKvrf+b7URcCAwEAAQJBAIFU0aQipuvdxdHsHMhX
5TFk0c1cSK/eeg7o3pGxhmAxfZLSQr73vP54Bk7xaaatOqqjCTUN0nixbEOeh18d
7DECIQD0fh2cf0xzQ07K3Fo6b1KAnjL4gR7XAJ/3rlXDAdUyCQIhANEXnSmU0tbf
CiYsqqloZTpZbpop6dEjt1vAWA++YLIfAiEAupvLtBgBXPRhnjpDb9hp6xtUIhJD
XKzwa9YXRUkP1SkCIQCnorYHQ2Eykklxx7ff8GnQOSlagiYK3gbAkdpIbQrbYwIg
STU47xTbB232EokQn4/ATVNXuYpRuFcLlIYeBNaB0aA=
-----END RSA PRIVATE KEY-----`)

func Signin(cfg *config.ServerConfig) func(*gin.Context) {
	return func(c *gin.Context) {
		var account Account
		ctx := context.Background()

		err := json.NewDecoder(c.Request.Body).Decode(&account)

		if err != nil {

			c.Writer.WriteHeader(http.StatusBadRequest)
			c.Writer.WriteString("User account info not in JSON format")
			return
		}

	}
}

func signToken(c Claims) (string, error) {

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
	key, err := jwt.ParseRSAPrivateKeyFromPEM(srt)
	if err != nil {
		return "", err
	}
	return token.SignedString(key)
}

type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

type Account struct {
	Password string `json:"password"`
	Username string `json:"username"`
}

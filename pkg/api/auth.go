package api

import (
	"crypto/rsa"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/momocomics/backend/pkg/config"
	"github.com/momocomics/backend/pkg/entity"
)

const (
	JwtTtl = 1 * time.Minute
	Issuer = "momocomic-backend"
)

//TODO: use real db
var users = map[string]string{
	"user1": "password1",
	"user2": "password2",
}

func SigninFn(cfg *config.ServerConfig) func(*gin.Context) {
	return func(c *gin.Context) {
		var account entity.Account
		err := json.NewDecoder(c.Request.Body).Decode(&account)

		if err != nil {

			c.Writer.WriteHeader(http.StatusBadRequest)
			c.Writer.WriteString("User account info not in JSON format")
			return
		}

		//TODO:read from db
		pwd, ok := users[account.Username]

		if !ok || pwd != account.Password {
			c.Writer.WriteHeader(http.StatusUnauthorized)
			return
		}

		tokenString, err := signToken(c.Request, account, cfg.PrivateKey())

		if err != nil {
			c.Writer.WriteHeader(http.StatusInternalServerError)
			c.Writer.WriteString("error while signing token")
			log.Printf("Token signing error: %v\n", err)
		}
		/*
			  // true means no scripts, http requests only. This has
			  // nothing to do with https vs http
			  HttpOnly: true,

			// Defaults to host-only, which means exact subdomain
			  // matching. Only change this to enable subdomains if you
			  // need to! The below code would work on any subdomain for
			  // yoursite.com
			  Domain: "yoursite.com",

			// Defaults to any path on your app, but you can use this
			  // to limit to a specific subdirectory. Eg:
			  Path: "/app/",
		*/
		c.SetCookie("token", tokenString, int(JwtTtl.Seconds()), "/", cfg.Domain(), true, true)
		//c.Header("Content-Type", "application/jwt")
		//c.Writer.WriteHeader(http.StatusOK)
		if cfg.IsDebug() {
			c.JSON(http.StatusOK, fmt.Sprintf("Token: %s", tokenString))
		}

	}
}

func signToken(r *http.Request, account entity.Account, key *rsa.PrivateKey) (string, error) {

	c := &entity.Claims{
		Username:  account.Username,
		UserAgent: r.UserAgent(),
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(JwtTtl).Unix(),
			Id:        uuid.New().String(),
			IssuedAt:  time.Now().Unix(),
			Issuer:    Issuer,
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
	return token.SignedString(key)
}

func VerifyToken(token string, key *rsa.PublicKey) (*jwt.Token, error) {

	t, err := jwt.Parse(token, func(token *jwt.Token) (i interface{}, e error) {
		return key, nil
	})

	if err != nil {
		return nil, fmt.Errorf("couldn't parse token: %v", err)
	}

	if !t.Valid {
		return nil, fmt.Errorf("token is invalid")
	}

	return t, nil

}

func RefreshTokenFn(cfg *config.ServerConfig) func(*gin.Context) {
	return func(c *gin.Context) {

	}
}

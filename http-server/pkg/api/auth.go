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

	"github.com/momocomics/backend/http-server/pkg/config"
	"github.com/momocomics/backend/http-server/pkg/entity"
)

const (
	JwtTtl           = 1 * time.Minute
	Issuer           = "momocomic-backend"
	CookieTokenEntry = "token"
	JwtExpiry        = 30 * time.Second
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
			return
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
		c.SetCookie(CookieTokenEntry, tokenString, int(JwtTtl.Seconds()), "/", cfg.Domain(), true, true)
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

func VerifyToken(token string, key *rsa.PublicKey) (*entity.Claims, error) {

	claims := &entity.Claims{}
	t, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (i interface{}, e error) {
		return key, nil
	})

	if err != nil {
		return nil, fmt.Errorf("couldn't parse token: %v", err)
	}

	if !t.Valid {
		return nil, fmt.Errorf("token is invalid")
	}

	return claims, nil

}

func RefreshTokenFn(cfg *config.ServerConfig) func(*gin.Context) {
	return func(c *gin.Context) {
		token, err := c.Cookie(CookieTokenEntry)
		if err != nil {
			c.Writer.WriteHeader(http.StatusBadRequest)
			c.Writer.WriteString("Token not in cookie")
			return
		}

		_, err = VerifyToken(token, cfg.PublicKey())

		if err != nil {
			c.Writer.WriteHeader(http.StatusUnauthorized)
			c.Writer.WriteString(err.Error())
			return
		}

		tokenString, err := refreshToken(token, cfg.PublicKey(), cfg.PrivateKey())
		if err != nil {
			c.Writer.WriteHeader(http.StatusInternalServerError)
			c.Writer.WriteString("error while signing token")
			log.Printf("Token refreshing error: %v\n", err)
			return
		}

		c.SetCookie(CookieTokenEntry, tokenString, int(JwtTtl.Seconds()), "/", cfg.Domain(), true, true)
		if cfg.IsDebug() {
			c.JSON(http.StatusOK, fmt.Sprintf("Token: %s", tokenString))
		}
	}
}

func refreshToken(token string, pubKey *rsa.PublicKey, privKey *rsa.PrivateKey) (string, error) {
	claims, err := VerifyToken(token, pubKey)
	if err != nil {
		return "", err
	}

	if time.Unix(claims.ExpiresAt, 0).Sub(time.Now()) > JwtExpiry {
		return "", fmt.Errorf("token expired more than %v seconds", JwtExpiry.Seconds())
	}

	claims.ExpiresAt = time.Now().Add(JwtTtl).Unix()
	t := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return t.SignedString(privKey)

}

func IsExpired(token string, key *rsa.PublicKey) (bool, error) {

	claims, err := VerifyToken(token, key)
	if err != nil {
		return false, err
	}
	return claims.ExpiresAt > time.Now().Unix(), nil

}

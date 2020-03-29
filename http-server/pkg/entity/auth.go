package entity

import "github.com/dgrijalva/jwt-go"

type Claims struct {
	//TODO: Add more info to track user see: https://stackoverflow.com/questions/216542/how-do-i-uniquely-identify-computers-visiting-my-web-site/3287761#
	Username  string `json:"username"`
	UserAgent string `json:"user_agent"`
	jwt.StandardClaims
}

type Account struct {
	Password string `json:"password"`
	Username string `json:"username"`
}

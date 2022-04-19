package models

import jwt "github.com/dgrijalva/jwt-go"

//Token struct declaration
type Token struct {
	UserID int
	Username   string
	Email  string
	*jwt.StandardClaims
}
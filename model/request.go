package model

import (
	"github.com/dgrijalva/jwt-go"
)

type CustomClaims struct {
	Id             uint
	UserName       string
	PasswordDigest string
	Nickname       string
	Status         string
	Avatar         string `gorm:"size:1000"`
	jwt.StandardClaims
}

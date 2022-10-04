package model

import (
	"github.com/golang-jwt/jwt"
	"gorm.io/gorm"
)

type Role string

const (
	Admin  Role = "Admin"
	Modo   Role = "Modo"
	Member Role = "Member"
)

type TypeToken string

const (
	AccessToken TypeToken = "AccessToken"
	RenewToken  TypeToken = "RenewToken"
)

type User struct {
	gorm.Model
	Email    string
	Pseudo   string
	Name     string
	Surname  string
	Password string
	Role     Role
}

type Token struct {
	gorm.Model
	Token string
}

type JwtCustomClaims struct {
	Pseudo    string `json:"pseudo"`
	Role      string `json:"role"`
	TokenType TypeToken
	jwt.StandardClaims
}

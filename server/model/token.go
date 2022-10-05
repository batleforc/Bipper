package model

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt"
	"gorm.io/gorm"
)

type TypeToken string

const (
	AccessToken TypeToken = "AccessToken"
	RenewToken  TypeToken = "RenewToken"
)

type JwtCustomClaims struct {
	Pseudo    string `json:"pseudo"`
	Role      string `json:"role"`
	TokenType TypeToken
	jwt.StandardClaims
}

// Create CustomClaims
func (c *JwtCustomClaims) CreateCustomClaims(pseudo string, role string, tokenType TypeToken) *JwtCustomClaims {
	c.Pseudo = pseudo
	c.Role = role
	c.TokenType = tokenType
	if tokenType == AccessToken {
		c.ExpiresAt = time.Now().Add(time.Hour * 8).Unix()
	} else {
		c.ExpiresAt = time.Now().Add(time.Hour * 24 * 7).Unix()
	}
	return c
}

// Create Token
func (c *JwtCustomClaims) CreateToken() (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
	if c.TokenType == AccessToken {
		return token.SignedString([]byte(os.Getenv("TOKEN_SIGN")))
	} else {
		return token.SignedString([]byte(os.Getenv("TOKEN_SIGN_RENEW")))
	}
}

type Token struct {
	gorm.Model
	UserID uint
	Token  string
}

// Get One Token by id
func (t *Token) GetOneToken(db *gorm.DB, id uint) (Token, error) {
	var token Token
	err := db.Where("id = ?", id).First(&token).Error
	return token, err
}

// Get All user's token
func (t *Token) GetAllToken(db *gorm.DB, id uint) ([]Token, error) {
	var tokens []Token
	err := db.Where("user_id = ?", id).Find(&tokens).Error
	return tokens, err
}

// Create Token
func (t *Token) CreateToken(db *gorm.DB, id uint, token string) error {
	t.UserID = id
	t.Token = token
	return db.Create(t).Error
}

// Delete Token
func (t *Token) DeleteToken(db *gorm.DB, id uint) error {
	return db.Where("user_id = ?", id).Delete(t).Error
}

package route

import (
	"batleforc/bipper/model"
	"net/http"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

type LoginBody struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginReturn struct {
	Pseudo      string     `json:"pseudo"`
	Role        model.Role `json:"role"`
	AccessToken string     `json:"access_token"`
	RenewToken  string     `json:"renew_token"`
}

func Login(c echo.Context) error {
	boudy := new(LoginBody)
	if err := c.Bind(boudy); err != nil {
		return echo.ErrUnauthorized
	}
	// TODO : Fetch user and validate password
	// https://echo.labstack.com/cookbook/jwt/

	AccessTokenClaim := &model.JwtCustomClaims{
		Pseudo:    "Joseph joestar",
		Role:      string(model.Member),
		TokenType: model.AccessToken,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 8).Unix(),
		},
	}
	RenewTokenClaim := &model.JwtCustomClaims{
		Pseudo:    "Joseph joestar",
		Role:      string(model.Member),
		TokenType: model.RenewToken,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24 * 7).Unix(),
		},
	}
	AccessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, AccessTokenClaim)
	RenewToken := jwt.NewWithClaims(jwt.SigningMethodHS256, RenewTokenClaim)
	signedAccess, err := AccessToken.SignedString([]byte(os.Getenv("TOKEN_SIGN")))
	if err != nil {
		return err
	}
	signedRenew, err := RenewToken.SignedString([]byte(os.Getenv("TOKEN_SIGN_RENEW")))
	if err != nil {
		return err
	}

	// todo : Insert renew in db and delete oldest if more than two
	return c.JSON(http.StatusOK, LoginReturn{
		Pseudo:      "JosephJoestar",
		Role:        model.Member,
		AccessToken: signedAccess,
		RenewToken:  signedRenew,
	})
}

package route

import (
	"batleforc/bipper/model"
	"net/http"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

type RenewTokenBody struct {
	RenewToken string `json:"renew_token"`
}

type RenewTokenReturn struct {
	Pseudo      string     `json:"pseudo"`
	Role        model.Role `json:"role"`
	AccessToken string     `json:"access_token"`
}

func RenewToken(c echo.Context) error {
	boudy := new(RenewTokenBody)
	if err := c.Bind(boudy); err != nil {
		return echo.ErrUnauthorized
	}
	// TODO : Check renew token in db and return new access token if renew not outdated and not deleted

	AccessTokenClaim := &model.JwtCustomClaims{
		Pseudo:    "Joseph joestar",
		Role:      string(model.Member),
		TokenType: model.AccessToken,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 8).Unix(),
		},
	}
	AccessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, AccessTokenClaim)
	signedAccess, err := AccessToken.SignedString([]byte(os.Getenv("TOKEN_SIGN")))
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, RenewTokenReturn{
		Pseudo:      "JosephJoestar",
		Role:        model.Member,
		AccessToken: signedAccess,
	})
}

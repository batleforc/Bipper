package route

import (
	"batleforc/bipper/model"
	"net/http"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type RenewTokenBody struct {
	RenewToken string `json:"renew_token"`
}

type RenewTokenReturn struct {
	Pseudo      string     `json:"pseudo"`
	Role        model.Role `json:"role"`
	AccessToken string     `json:"access_token"`
}

// Renew Token
// @Summary Renew Token
// @Description Renew Token via refresh token
// @Tags Auth
// @Accept  json
// @Param Request body route.RenewTokenBody true "Renew body"
// @Success 200 {object} route.RenewTokenReturn "Renew return"
// @Router /renewtoken [post]
func RenewToken(c echo.Context) error {
	boudy := new(RenewTokenBody)
	if err := c.Bind(boudy); err != nil {
		return echo.ErrUnauthorized
	}
	token := new(model.Token)
	err := token.GetOneTokenByToken(c.Get("db").(*gorm.DB), boudy.RenewToken)
	if err != nil {
		return c.String(http.StatusUnauthorized, "Token incorrect")
	}
	claims, err := token.ValidateRenewToken(boudy.RenewToken)
	if err != nil {
		return c.String(http.StatusUnauthorized, err.Error())
	}
	claims.TokenType = model.AccessToken
	signedToken, err := claims.CreateToken()
	if err != nil {
		return c.String(http.StatusUnauthorized, err.Error())
	}
	return c.JSON(http.StatusOK, RenewTokenReturn{
		Pseudo:      claims.Pseudo,
		Role:        model.Role(claims.Role),
		AccessToken: signedToken,
	})
}

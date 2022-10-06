package route

import (
	"batleforc/bipper/model"
	"net/http"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type LogoutBody struct {
	RenewToken string `json:"renew_token"`
}

// Logout user
// @Summary Logout user
// @Description Logout user
// @Tags Auth
// @Accept  json
// @Param Request body route.LogoutBody true "Logout body"
// @Router /logout [post]
func Logout(c echo.Context) error {
	boudy := new(LogoutBody)
	if err := c.Bind(boudy); err != nil {
		return echo.ErrUnauthorized
	}
	token := new(model.Token)
	err := token.GetOneTokenByToken(c.Get("db").(*gorm.DB), boudy.RenewToken)
	if err != nil {
		return c.String(http.StatusUnauthorized, "Token incorrect")
	}
	if err := token.DeleteToken(c.Get("db").(*gorm.DB), token.ID); err != nil {
		return c.String(http.StatusInternalServerError, "Error while deleting token")
	}
	return c.NoContent(http.StatusOK)
}

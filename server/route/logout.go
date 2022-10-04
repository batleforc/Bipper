package route

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type LogoutBody struct {
	RenewToken string `json:"renew_token"`
}

func LogOut(c echo.Context) error {
	boudy := new(LoginBody)
	if err := c.Bind(boudy); err != nil {
		return echo.ErrUnauthorized
	}
	// TODO : Delete renew token in db
	return c.NoContent(http.StatusOK)
}

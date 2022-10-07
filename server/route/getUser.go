package route

import (
	"batleforc/bipper/model"

	"github.com/labstack/echo/v4"
)

// Get user
// @Summary Get user
// @Description Get user
// @Tags User
// @Security BearerAuth
// @Success 200 {object} model.User "user return"
// @Router /user [get]
func GetUser(c echo.Context) error {
	user := c.Get("User").(*model.User)
	return c.JSON(200, user)
}

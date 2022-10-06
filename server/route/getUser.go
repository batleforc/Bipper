package route

import (
	"batleforc/bipper/model"
	"strings"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

// Get user
// @Summary Get user
// @Description Get user
// @Tags Auth
// @Security BearerAuth
// @Success 200 {object} model.User "user return"
// @Router /user [get]
func GetUser(c echo.Context) error {
	bearerToken := c.Request().Header.Get("Authorization")
	splitToken := strings.Split(bearerToken, " ")
	if len(splitToken) != 2 || splitToken[0] != "Bearer" {
		return echo.ErrUnauthorized
	}
	token := model.Token{}
	claim, err := token.ValidateAccessToken(splitToken[1])
	if err != nil {
		return echo.ErrUnauthorized
	}
	user := model.User{}
	err = user.GetUserByPseudo(c.Get("db").(*gorm.DB), claim.Pseudo)
	if err != nil {
		return echo.ErrUnauthorized
	}
	return c.JSON(200, user)
}

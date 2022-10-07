package middleware

import (
	"batleforc/bipper/model"
	"strings"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func Auth() func(next echo.HandlerFunc) echo.HandlerFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
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
			c.Set("User", &user)
			return next(c)
		}
	}
}

package route

import (
	"batleforc/bipper/model"
	"net/http"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
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

// Login user
// @Summary Login user
// @Description Login user
// @Tags Auth
// @Accept  json
// @Param Request body route.LoginBody true "Login body"
// @Success 200 {object} route.LoginReturn "Login return"
// @Router /auth/login [post]
func Login(c echo.Context) error {
	boudy := new(LoginBody)
	if err := c.Bind(boudy); err != nil {
		return echo.ErrUnauthorized
	}

	user := new(model.User)
	if err := user.GetUserByMail(c.Get("db").(*gorm.DB), boudy.Email); err != nil {
		return c.String(http.StatusUnauthorized, "Email or password incorrect")
	}
	if !user.CheckPassword(boudy.Password) {
		return c.String(http.StatusUnauthorized, "Password incorrect")
	}

	AccesTokenClaim := new(model.JwtCustomClaims)
	RefreshTokenClaim := new(model.JwtCustomClaims)

	AccesTokenClaim.CreateCustomClaims(user.Pseudo, string(user.Role), model.AccessToken)
	RefreshTokenClaim.CreateCustomClaims(user.Pseudo, string(user.Role), model.RenewToken)

	accessToken, err := AccesTokenClaim.CreateToken()
	if err != nil {
		return c.String(http.StatusInternalServerError, "Error while creating token")
	}
	renewToken, err := RefreshTokenClaim.CreateToken()
	if err != nil {
		return c.String(http.StatusInternalServerError, "Error while creating token")
	}
	token := new(model.Token)
	tokens, err := token.GetAllToken(c.Get("db").(*gorm.DB), user.ID)
	if err != nil {
		return c.String(http.StatusInternalServerError, "Error while creating token")
	}
	if len(tokens) <= 4 {
		token.DeleteToken(c.Get("db").(*gorm.DB), tokens[len(tokens)-1].ID)
	}
	token.CreateToken(c.Get("db").(*gorm.DB), user.ID, renewToken)
	return c.JSON(http.StatusOK, LoginReturn{
		Pseudo:      "JosephJoestar",
		Role:        model.Member,
		AccessToken: accessToken,
		RenewToken:  renewToken,
	})
}

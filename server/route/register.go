package route

import (
	"batleforc/bipper/model"
	"net/http"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type RegisterBody struct {
	Pseudo   string `json:"pseudo"`
	Email    string `json:"email"`
	Name     string `json:"name"`
	Surname  string `json:"surname"`
	Password string `json:"password"`
}

type RegisterReturn struct {
	Registered bool   `json:"registered"`
	Error      bool   `json:"error"`
	Message    string `json:"message"`
}

// Register User
// @Summary Register User
// @Description Register User
// @Tags Auth
// @Accept  json
// @Param Request body route.RegisterBody true "Register body"
// @Success 200 {object} route.RegisterReturn "Register return"
// @Router /register [post]
func Register(c echo.Context) error {
	body := new(RegisterBody)
	if err := c.Bind(body); err != nil {
		return c.JSON(http.StatusBadRequest, RegisterReturn{
			Registered: false,
			Error:      true,
			Message:    "Body is not valid",
		})
	}
	if body.Email != "" && body.Pseudo != "" && body.Name != "" && body.Surname != "" && body.Password != "" {
		return c.JSON(http.StatusBadRequest, RegisterReturn{
			Registered: false,
			Error:      true,
			Message:    "Body is not valid, missing or empty property",
		})
	}

	user := new(model.User)
	if err := user.GetUserByMail(c.Get("db").(*gorm.DB), body.Email); err != nil || user.ID != 0 {
		return c.JSON(http.StatusBadRequest, RegisterReturn{
			Registered: false,
			Error:      true,
			Message:    "Email already used or error while getting user",
		})
	}
	user = new(model.User)
	if err := user.GetUserByPseudo(c.Get("db").(*gorm.DB), body.Pseudo); err != nil || user.ID != 0 {
		return c.JSON(http.StatusBadRequest, RegisterReturn{
			Registered: false,
			Error:      true,
			Message:    "Pseudo already used or error while getting user",
		})
	}
	user.Email = body.Email
	user.Pseudo = body.Pseudo
	user.Name = body.Name
	user.Surname = body.Surname
	user.Role = model.Member
	user.Picture = ""
	user.HashPassword(body.Password)
	err := user.UpdateOrCreateUser(c.Get("db").(*gorm.DB))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, RegisterReturn{
			Registered: false,
			Error:      true,
			Message:    "Error while creating user : " + err.Error(),
		})
	}
	return c.JSON(http.StatusOK, RegisterReturn{
		Registered: true,
		Error:      false,
		Message:    "User created, Please login",
	})
}

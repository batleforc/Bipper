package route

import (
	"batleforc/bipper/model"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type SetUserBody struct {
	Surname *string `json:"surname,omitempty"`
	Name    *string `json:"name,omitempty"`
}

type SetUserReturn struct {
	Updated bool   `json:"updated"`
	Error   bool   `json:"error"`
	Message string `json:"message"`
}

// Set user
// @Summary Set user
// @Description Set user
// @Tags User
// @Accept  json
// @Security BearerAuth
// @Param Request body route.SetUserBody true "Set user body"
// @Success 200 {object} route.SetUserReturn "Set user return"
// @Router /user [post]
func SetUser(c echo.Context) error {
	body := new(SetUserBody)
	if err := c.Bind(body); err != nil {
		return c.JSON(400, SetUserReturn{
			Updated: false,
			Error:   true,
			Message: "Body is not valid",
		})
	}
	user := c.Get("User").(*model.User)
	if body.Name != nil {
		user.Name = *body.Name
	}
	if body.Surname != nil {
		user.Surname = *body.Surname
	}
	if err := user.UpdateOrCreateUser(c.Get("db").(*gorm.DB)); err != nil {
		return c.JSON(500, SetUserReturn{
			Updated: false,
			Error:   true,
			Message: "Error while updating user",
		})
	}
	return c.JSON(200, SetUserReturn{
		Updated: true,
		Error:   false,
		Message: "User updated",
	})
}

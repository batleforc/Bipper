package route

import (
	"batleforc/bipper/model"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type CreateChanBody struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Picture     string `json:"picture"`
	Private     bool   `json:"private"`
}

type CreateChanReturn struct {
	Created bool   `json:"updated"`
	Error   bool   `json:"error"`
	Message string `json:"message"`
	PassKey string `json:"passkey"`
}

// Create channel
// @Summary Create channel
// @Description Create channel, Name has to be unique
// @Tags Chan
// @Accept  json
// @Security BearerAuth
// @Param Request body route.CreateChanBody true "Create channel body"
// @Success 200 {object} route.CreateChanReturn "Create channel return"
// @Failure 400 {object} route.CreateChanReturn "Create channel return"
// @Failure 500 {object} route.CreateChanReturn "Create channel return"
// @Router /chan [post]
func CreateChan(c echo.Context) error {
	body := new(CreateChanBody)
	if err := c.Bind(body); err != nil {
		return c.JSON(400, CreateChanReturn{
			Created: false,
			Error:   true,
			Message: "Body is not valid",
		})
	}
	channel := model.Channel{}
	if err := channel.GetChannelByName(c.Get("db").(*gorm.DB), body.Name); err != nil && err != gorm.ErrRecordNotFound {
		return c.JSON(500, CreateChanReturn{
			Created: false,
			Error:   true,
			Message: "Error while getting channel",
		})
	}
	if channel.ID != 0 {
		return c.JSON(400, CreateChanReturn{
			Created: false,
			Error:   true,
			Message: "Channel already exist",
		})
	}
	channel.Name = body.Name
	channel.Description = body.Description
	channel.Picture = body.Picture
	channel.Private = body.Private
	passKey := ""
	if body.Private {
		passKey = channel.GeneratePassKey()
	}
	if err := channel.CreateChannel(c.Get("db").(*gorm.DB)); err != nil {
		return c.JSON(500, CreateChanReturn{
			Created: false,
			Error:   true,
			Message: "Error while creating channel",
		})
	}
	return c.JSON(200, CreateChanReturn{
		Created: true,
		Error:   false,
		Message: "Channel created",
		PassKey: passKey,
	})
}

package route

import (
	"batleforc/bipper/model"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type CheckChanNameReturn struct {
	Available bool   `json:"available"`
	Error     bool   `json:"error"`
	Message   string `json:"message"`
}

type CheckChanNameBody struct {
	Name string `json:"name"`
}

// CheckChanName check if channel name is available
// @Summary Check if channel name is available
// @Description Check if channel name is available
// @Tags Chan
// @Accept json
// @Security BearerAuth
// @Param Request body route.CheckChanNameBody true "Check channel name body"
// @Success 200 {object} route.CheckChanNameReturn "Chann available or not"
// @Failure 400 {object} route.CheckChanNameReturn "Body not valid"
// @Failure 500 {object} route.CheckChanNameReturn "Internal server error"
// @Router /chan/name [post]
func CheckChanName(c echo.Context) error {
	body := new(CheckChanNameBody)
	if err := c.Bind(body); err != nil {
		return c.JSON(400, CheckChanNameReturn{
			Available: false,
			Error:     true,
			Message:   "Body is not valid",
		})
	}
	channel := model.Channel{}
	err := channel.GetChannelByName(c.Get("db").(*gorm.DB), body.Name)
	if err != nil && err != gorm.ErrRecordNotFound {
		return c.JSON(500, CheckChanNameReturn{
			Available: false,
			Error:     true,
			Message:   "Error while getting channel",
		})
	}
	if channel.ID != 0 {
		return c.JSON(200, CheckChanNameReturn{
			Available: false,
			Error:     false,
			Message:   "Channel already exist",
		})
	}
	return c.JSON(200, CheckChanNameReturn{
		Available: true,
		Error:     false,
		Message:   "Channel name is available",
	})
}

package route

import (
	"batleforc/bipper/model"
	"strconv"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type GetOneChanReturn struct {
	Error   bool   `json:"error"`
	Message string `json:"message"`
}

// Get One Channel by id
// @Summary Get One Channel by id
// @Description Get One Channel by id
// @Tags Chan
// @Accept json
// @Security BearerAuth
// @Success 200 {object} model.Channel "Channel"
// @Failure 400 {object} route.GetOneChanReturn "Chan Id is not valid"
// @Failure 403 {object} route.GetOneChanReturn "User is not in channel"
// @Failure 500 {object} route.GetOneChanReturn "Error while getting channel, (can be normal if not exist)"
// @Param chanId path string true "Channel id"
// @Router /chan/{chanId} [get]
func GetOneChan(c echo.Context) error {
	chanIdString := c.Param("chanId")
	chanId, err := strconv.ParseUint(chanIdString, 10, 32)
	if err != nil {
		return c.JSON(400, GetOneChanReturn{
			Error:   true,
			Message: "ChanId is not valid",
		})
	}
	channel := model.Channel{}
	if err := channel.GetChannel(c.Get("db").(*gorm.DB), uint(chanId)); err != nil {
		return c.JSON(500, GetOneChanReturn{
			Error:   true,
			Message: "Error while getting channel",
		})
	}
	if channel.Private {
		user := c.Get("User").(*model.User)
		if !channel.IsUserInChannel(user.ID) {
			return c.JSON(403, GetOneChanReturn{
				Error:   true,
				Message: "User is not in channel",
			})
		}
	}
	return c.JSON(200, channel)
}

package route

import (
	"batleforc/bipper/model"
	"strconv"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type RenewChanPasswordReturn struct {
	Updated bool   `json:"updated"`
	Error   bool   `json:"error"`
	Message string `json:"message"`
}

// Reset channel password
// @Summary Reset channel password
// @Description Reset channel password
// @Tags Chan
// @Accept json
// @Security BearerAuth
// @Success 200 {object} route.RenewChanPasswordReturn "Reset channel password return"
// @Param chanId path string true "Channel id"
// @Router /chan/{chanId}/renew [post]
func RenewChanPassword(c echo.Context) error {
	chanIdString := c.Param("chanId")
	chanId, err := strconv.ParseUint(chanIdString, 10, 32)
	if err != nil {
		return c.JSON(400, RenewChanPasswordReturn{
			Updated: false,
			Error:   true,
			Message: "ChanId is not valid",
		})
	}
	user := c.Get("User").(*model.User)
	channel := model.Channel{}
	if err := channel.GetChannel(c.Get("db").(*gorm.DB), uint(chanId)); err != nil {
		return c.JSON(500, RenewChanPasswordReturn{
			Updated: false,
			Error:   true,
			Message: "Error while getting channel",
		})
	}
	if channel.Owner != user.ID {
		userRight, err := channel.GetUserById(c.Get("db").(*gorm.DB), user.ID)
		if err != nil {
			return c.JSON(500, RenewChanPasswordReturn{
				Updated: false,
				Error:   true,
				Message: "Error while getting user right",
			})
		} else if userRight == nil {
			return c.JSON(403, RenewChanPasswordReturn{
				Updated: false,
				Error:   true,
				Message: "User is not in channel",
			})
		} else if !userRight.CanMod {
			return c.JSON(403, RenewChanPasswordReturn{
				Updated: false,
				Error:   true,
				Message: "User can't mod channel",
			})
		}
	}
	chanPassword := channel.GeneratePassKey()
	if err := channel.UpdateChannel(c.Get("db").(*gorm.DB)); err != nil {
		return c.JSON(500, RenewChanPasswordReturn{
			Updated: false,
			Error:   true,
			Message: "Error while updating channel",
		})
	}
	return c.JSON(200, RenewChanPasswordReturn{
		Updated: true,
		Error:   false,
		Message: chanPassword,
	})
}

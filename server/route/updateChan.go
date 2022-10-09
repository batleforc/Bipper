package route

import (
	"batleforc/bipper/model"
	"strconv"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type UpdateChannelReturn struct {
	Updated bool   `json:"updated"`
	Error   bool   `json:"error"`
	Message string `json:"message"`
}

type UpdateChannelBody struct {
	Description string `json:"description"`
	Private     bool   `json:"private"`
}

// Update Channel
// @Summary Update Channel
// @Description Update Channel
// @Tags Chan
// @Security BearerAuth
// @Param chanId path int true "Channel Id"
// @Param body body route.UpdateChannelBody true "Body"
// @Success 200 {object} route.UpdateChannelReturn "Update User Channel Right"
// @Failure 400 {object} route.UpdateChannelReturn "Chan Id is not valid or error while getting body"
// @Failure 403 {object} route.UpdateChannelReturn "User is not in channel or User is not allowed to update channel"
// @Failure 500 {object} route.UpdateChannelReturn "Error while getting channel, channel user or updating channel"
// @Router /chan/{chanId} [post]
func UpdateChannel(c echo.Context) error {
	chanIdString := c.Param("chanId")
	chanId, err := strconv.ParseUint(chanIdString, 10, 32)
	if err != nil {
		return c.JSON(400, UpdateUserChanRightReturn{
			Error:   true,
			Updated: false,
			Message: "ChanId is not valid",
		})
	}
	channel := model.Channel{}
	if err := channel.GetChannel(c.Get("db").(*gorm.DB), uint(chanId)); err != nil {
		return c.JSON(500, UpdateUserChanRightReturn{
			Updated: false,
			Error:   true,
			Message: "Error while getting channel",
		})
	}
	user := c.Get("User").(*model.User)
	if !channel.IsUserInChannel(user.ID) {
		return c.JSON(403, UpdateUserChanRightReturn{
			Updated: false,
			Error:   true,
			Message: "User is not in channel",
		})
	}
	if channel.Owner != user.ID {
		usr, err := channel.GetUserById(user.ID)
		if err != nil {
			return c.JSON(500, UpdateUserChanRightReturn{
				Updated: false,
				Error:   true,
				Message: "Error while getting channel user",
			})
		}
		if !usr.CanMod {
			return c.JSON(403, UpdateUserChanRightReturn{
				Updated: false,
				Error:   true,
				Message: "User is not allowed to update channel",
			})
		}
	}
	body := UpdateChannelBody{}
	if err := c.Bind(&body); err != nil {
		return c.JSON(400, UpdateUserChanRightReturn{
			Updated: false,
			Error:   true,
			Message: "Error while getting body",
		})
	}
	channel.Description = body.Description
	channel.Private = body.Private
	if err := channel.UpdateChannel(c.Get("db").(*gorm.DB)); err != nil {
		return c.JSON(500, UpdateUserChanRightReturn{
			Updated: false,
			Error:   true,
			Message: "Error while updating channel",
		})
	}
	return c.JSON(200, UpdateUserChanRightReturn{
		Updated: true,
		Error:   false,
		Message: "Channel updated",
	})
}

package route

import (
	"batleforc/bipper/model"
	"strconv"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type UpdateUserChanRightBody struct {
	UserId  uint `json:"userId"`
	CanSend bool
	CanRead bool
	CanMod  bool
}

type UpdateUserChanRightReturn struct {
	Updated bool   `json:"updated"`
	Error   bool   `json:"error"`
	Message string `json:"message"`
}

// Update User Channel Right
// @Summary Update User Channel Right
// @Description Update User Channel Right
// @Tags Chan
// @Security BearerAuth
// @Param chanId path int true "Channel Id"
// @Param body body route.UpdateUserChanRightBody true "Body"
// @Success 200 {object} route.UpdateUserChanRightReturn "Update User Channel Right"
// @Failure 400 {object} route.UpdateUserChanRightReturn "Chan Id is not valid or error while getting body"
// @Failure 403 {object} route.UpdateUserChanRightReturn "User is not in channel or User is not allowed to update user right"
// @Failure 500 {object} route.UpdateUserChanRightReturn "Error while getting channel, channel user or updating user right"
// @Router /chan/{chanId}/right [post]
func UpdateUserChanRight(c echo.Context) error {
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
				Message: "User is not allowed to update user right",
			})
		}
	}
	body := UpdateUserChanRightBody{}
	if err := c.Bind(&body); err != nil {
		return c.JSON(400, UpdateUserChanRightReturn{
			Updated: false,
			Error:   true,
			Message: "Error while getting body",
		})
	}
	chanUser, err := channel.GetUserById(user.ID)
	if err != nil {
		return c.JSON(500, UpdateUserChanRightReturn{
			Updated: false,
			Error:   true,
			Message: "Error while getting channel user",
		})
	}
	chanUser.CanRead = body.CanRead
	chanUser.CanSend = body.CanSend
	chanUser.CanMod = body.CanMod
	if err := chanUser.UpdateChannelUser(c.Get("db").(*gorm.DB)); err != nil {
		return c.JSON(500, UpdateUserChanRightReturn{
			Updated: false,
			Error:   true,
			Message: "Error while updating user right",
		})
	}
	return c.JSON(200, UpdateUserChanRightReturn{
		Updated: true,
		Error:   false,
		Message: "User right updated",
	})
}

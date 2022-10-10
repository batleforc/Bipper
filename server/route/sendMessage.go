package route

import (
	"batleforc/bipper/model"
	"strconv"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type SendMessageContent struct {
	Content string `json:"content"`
}

// Send Message
// @Summary Send Message
// @Description Send Message
// @Tags Message
// @Security BearerAuth
// @Success 200 {object} route.GetOneChanReturn "Send Message"
// @Failure 400 {object} route.GetOneChanReturn "Send Message"
// @Failure 403 {object} route.GetOneChanReturn "Send Message"
// @Failure 500 {object} route.GetOneChanReturn "Send Message"
// @Param chanId path int true "Channel Id"
// @Param body body route.SendMessageContent true "Body"
// @Router /chan/{chanId}/message [post]
func SendMessage(c echo.Context) error {
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
	user := c.Get("User").(*model.User)
	if !channel.IsUserInChannel(user.ID) {
		return c.JSON(403, GetOneChanReturn{
			Error:   true,
			Message: "User is not in channel",
		})
	}
	if channel.Owner != user.ID {
		usr, err := channel.GetUserById(user.ID)
		if err != nil {
			return c.JSON(500, GetOneChanReturn{
				Error:   true,
				Message: "Error while getting channel user",
			})
		}
		if !usr.CanSend {
			return c.JSON(403, GetOneChanReturn{
				Error:   true,
				Message: "User is not allowed to update channel",
			})
		}
	}
	body := SendMessageContent{}
	if err := c.Bind(&body); err != nil {
		return c.JSON(400, GetOneChanReturn{
			Error:   true,
			Message: "Error while getting body",
		})
	}
	message := model.Message{
		Content:   body.Content,
		ChannelID: channel.ID,
		UserID:    user.ID,
	}
	if err := message.CreateMessage(c.Get("db").(*gorm.DB)); err != nil {
		return c.JSON(500, GetOneChanReturn{
			Error:   true,
			Message: "Error while creating message",
		})
	}
	return c.JSON(200, GetOneChanReturn{
		Error:   false,
		Message: "Message created",
	})
}

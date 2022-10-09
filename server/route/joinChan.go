package route

import (
	"batleforc/bipper/model"
	"strconv"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type JoinChanReturn struct {
	Joined  bool   `json:"joined"`
	Error   bool   `json:"error"`
	Message string `json:"message"`
}
type JoinChanBody struct {
	Password *string `json:"password,omitempty"`
}

// Get One Channel by id
// @Summary Get One Channel by id
// @Description Get One Channel by id
// @Tags Chan
// @Accept json
// @Security BearerAuth
// @Success 200 {object} model.Channel "Channel"
// @Failure 400 {object} route.JoinChanReturn "Chan Id is not valid"
// @Failure 403 {object} route.JoinChanReturn "Chan Id is not valid"
// @Failure 500 {object} route.JoinChanReturn "Error while getting channel, (can be normal if not exist)"
// @Param chanId path string true "Channel id"
// @Router /chan/{chanId}/join [post]
func JoinChan(c echo.Context) error {
	chanIdString := c.Param("chanId")
	chanId, err := strconv.ParseUint(chanIdString, 10, 32)
	if err != nil {
		return c.JSON(400, JoinChanReturn{
			Joined:  false,
			Error:   true,
			Message: "ChanId is not valid",
		})
	}
	channel := model.Channel{}
	if err := channel.GetChannel(c.Get("db").(*gorm.DB), uint(chanId)); err != nil {
		return c.JSON(500, JoinChanReturn{
			Joined:  false,
			Error:   true,
			Message: "Error while getting channel",
		})
	}
	body := JoinChanBody{}
	if err := c.Bind(&body); err != nil || body.Password == nil {
		return c.JSON(400, JoinChanReturn{
			Joined:  false,
			Error:   true,
			Message: "Error while getting body",
		})
	}
	if !channel.CheckPassKey(*body.Password) {
		return c.JSON(403, JoinChanReturn{
			Joined:  false,
			Error:   true,
			Message: "Wrong passkey",
		})
	}
	chanUser := model.ChannelUser{
		ChannelID: uint(chanId),
		UserID:    c.Get("User").(*model.User).ID,
		CanSend:   false,
		CanRead:   false,
		CanMod:    false,
	}
	if err := chanUser.CreateChannelUser(c.Get("db").(*gorm.DB)); err != nil {
		return c.JSON(500, JoinChanReturn{
			Error:   true,
			Message: "Error while creating channel user",
		})
	}
	channel = model.Channel{}
	if err := channel.GetChannel(c.Get("db").(*gorm.DB), uint(chanId)); err != nil {
		return c.JSON(500, JoinChanReturn{
			Joined:  false,
			Error:   true,
			Message: "Error while getting channel",
		})
	}

	return c.JSON(200, channel)
}

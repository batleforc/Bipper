package route

import (
	"batleforc/bipper/model"
	"strconv"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

// Get One Channel message by id
// @Summary Get One Channel message by id
// @Description Get One Channel message by id, if user not in chan can't see message and if user hasn't the read right can only see past 24 hour message
// @Tags Chan
// @Accept json
// @Security BearerAuth
// @Success 200 {object} []model.Message "message"
// @Failure 400 {object} route.GetOneChanReturn "Chan Id is not valid"
// @Failure 403 {object} route.GetOneChanReturn "User is not in channel"
// @Failure 500 {object} route.GetOneChanReturn "Error while getting channel, (can be normal if not exist)"
// @Param chanId path string true "Channel id"
// @Param limit query int false "Limit of message"
// @Param page query int false "Page of message"
// @Router /chan/{chanId}/message [get]
func GetOneChanMessage(c echo.Context) error {
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
	limitString := c.QueryParam("limit")
	pageString := c.QueryParam("page")
	limit, errLimit := strconv.Atoi(limitString)
	page, errPage := strconv.Atoi(pageString)
	if limitString == "" || errLimit != nil {
		limit = 10
	}
	if pageString == "" || errPage != nil {
		page = 1
	}

	chanUser, err := channel.GetUserById(user.ID)
	if err != nil {
		return c.JSON(500, GetOneChanReturn{
			Error:   true,
			Message: "Error while getting channel",
		})
	}
	if chanUser.CanRead {
		return c.JSON(200, channel.GetPaginatedMessages(c.Get("db").(*gorm.DB), limit, page))
	} else {
		return c.JSON(200, channel.GetPast24HourMessages(c.Get("db").(*gorm.DB), limit, page))
	}
}

package route

import (
	"batleforc/bipper/model"
	"strconv"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type DeleteChanReturn struct {
	Deleted bool   `json:"deleted"`
	Error   bool   `json:"error"`
	Message string `json:"message"`
}

// Delete channel
// @Summary Delete channel
// @Description Delete channel
// @Tags Chan
// @Accept json
// @Security BearerAuth
// @Success 200 {object} route.DeleteChanReturn "Is Deleted ?"
// @Failure 400 {object} route.DeleteChanReturn "Chan Id is not valid"
// @Failure 403 {object} route.DeleteChanReturn "User is not the owner of the channel"
// @Failure 500 {object} route.DeleteChanReturn "Error while getting channel or deleting channel"
// @Param chanId path string true "Channel id"
// @Router /chan/{chanId} [delete]
func DeleteChan(c echo.Context) error {
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
	if channel.Owner != user.ID {
		return c.JSON(403, DeleteChanReturn{
			Deleted: false,
			Error:   true,
			Message: "User is not the owner of the channel",
		})
	}
	if err := channel.DeleteChannel(c.Get("db").(*gorm.DB), channel.ID); err != nil {
		return c.JSON(500, DeleteChanReturn{
			Deleted: false,
			Error:   true,
			Message: "Error while deleting channel",
		})
	}
	return c.JSON(200, DeleteChanReturn{
		Deleted: true,
		Error:   false,
		Message: "Channel deleted",
	})
}

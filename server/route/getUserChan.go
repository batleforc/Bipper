package route

import (
	"batleforc/bipper/model"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type GetUserChanReturn struct {
	OwnChan    []model.Channel  `json:"ownChan"`
	MemberChan *[]model.Channel `json:"memberChan"`
}

// Get user channels
// @Summary Get user channels
// @Description Get user channels
// @Tags Chan
// @Security BearerAuth
// @Success 200 {object} route.GetUserChanReturn "Get User Channels"
// @Router /chan [get]
func GetUserChan(c echo.Context) error {
	user := c.Get("User").(*model.User)
	memberChan, err := user.GetChannels(c.Get("DB").(*gorm.DB))
	if err != nil {
		return c.JSON(500, err)
	}
	return c.JSON(200, GetUserChanReturn{
		OwnChan:    user.MyChannels,
		MemberChan: memberChan,
	})
}

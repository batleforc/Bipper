package route

import (
	"batleforc/bipper/model"
	"strconv"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

// Get Public channel
// @Summary Get Public channel
// @Description Get Public channel
// @Tags Chan
// @Accept json
// @Security BearerAuth
// @Success 200 {object} []model.Channel "Channel"
// @Failure 500 {object} route.GetOneChanReturn "Error while getting channel"
// @Param chanId path string true "Channel id"
// @Param limit query int false "Limit of Channel"
// @Param search query int false "Search in Channel"
// @Param page query int false "Page of Channel"
// @Router /chan/public [get]
func GetPublicChannels(c echo.Context) error {
	limitString := c.QueryParam("limit")
	pageString := c.QueryParam("page")
	search := c.QueryParam("search")
	limit, errLimit := strconv.Atoi(limitString)
	page, errPage := strconv.Atoi(pageString)
	if limitString == "" || errLimit != nil {
		limit = 10
	}
	if pageString == "" || errPage != nil {
		page = 1
	}
	channel := model.Channel{}
	channels, err := channel.GetPublicChannelsSearch(c.Get("db").(*gorm.DB), limit, page, search)
	if err != nil {
		return c.JSON(500, GetOneChanReturn{
			Error:   true,
			Message: "Error while getting channels",
		})
	}
	return c.JSON(200, channels)
}

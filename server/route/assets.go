package route

import (
	"batleforc/bipper/model"
	"fmt"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
)

// Serve static asset
// @Summary Serve static asset
// @Description Serve static asset
// @Tags Asset
// @Success 200 {file} nil
// @Param fileName path string true "fileName"
// @Router /asset/{fileName} [get]
func Asset(c echo.Context) error {
	file := c.Param("file")
	staticImage := &model.StaticImageHandler{}
	fileObject, err := staticImage.GetFile(file)
	if err != nil {
		return c.String(http.StatusNotFound, "File not found")
	}
	extension := strings.Split(file, ".")[2]
	return c.Stream(http.StatusOK, fmt.Sprintf("image/%s", extension), fileObject)
}

package route

import (
	"batleforc/bipper/model"
	"bytes"
	"fmt"
	"io"
	"net/http"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type SetPictureReturn struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

// SetPicture user
// @Summary SetPicture user
// @Description SetPicture user
// @Tags User
// @Security BearerAuth
// @Accept multipart/form-data
// @Param file formData file true ".jpeg, .png, .gif"
// @Success 200 {object} route.SetPictureReturn "SetPicture return"
// @Router /user/setpicture [post]
func SetPicture(c echo.Context) error {
	file, err := c.FormFile("file")
	if err != nil {
		return c.JSON(http.StatusBadRequest, SetPictureReturn{
			Success: false,
			Message: "file not found",
		})
	}
	src, err := file.Open()
	if err != nil {
		return c.JSON(http.StatusBadRequest, SetPictureReturn{
			Success: false,
			Message: "file not found",
		})
	}
	defer src.Close()
	staticImageHandler := &model.StaticImageHandler{}
	buf, _ := io.ReadAll(src)
	mimeType := staticImageHandler.GetImageType(buf)
	if mimeType == "" {
		return c.JSON(http.StatusBadRequest, SetPictureReturn{
			Success: false,
			Message: "Invalid File",
		})
	}

	// TODO : Crop Image
	newSrc := io.NopCloser(bytes.NewBuffer(buf))
	user := c.Get("User").(*model.User)
	fileName := fmt.Sprintf("%s.%s", user.Pseudo, file.Filename)
	err = staticImageHandler.SaveImage(newSrc, fileName)
	if err != nil {
		return c.JSON(http.StatusBadRequest, SetPictureReturn{
			Success: false,
			Message: "could not save file",
		})
	}
	staticImageHandler.DeleteOldestFileIfNeeded(user.Pseudo)
	user.Picture = fileName
	err = user.UpdateOrCreateUser(c.Get("db").(*gorm.DB))
	if err != nil {
		return c.JSON(http.StatusBadRequest, SetPictureReturn{
			Success: false,
			Message: "Could not change profile picture",
		})
	}
	return c.JSON(http.StatusOK, SetPictureReturn{
		Success: true,
		Message: "success",
	})
}

package controller

import (
	"ecommerce-project/config"
	"ecommerce-project/feature/user/entities"
	"ecommerce-project/middlewares"
	"ecommerce-project/utils/helper"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
)

type Delivery struct {
	FromTo entities.ServiceInterface
}

func New(e *echo.Echo, data entities.ServiceInterface) {
	handler := &Delivery{
		FromTo: data,
	}

	e.GET("/profile", handler.SeeProfile, middlewares.JWTMiddleware())
	e.PUT("/profile/image", handler.NewImages, middlewares.JWTMiddleware())
	e.PUT("/profile/name", handler.NewName, middlewares.JWTMiddleware())
	e.PUT("/profile/username", handler.NewUserName, middlewares.JWTMiddleware())
	e.PUT("/profile/password", handler.NewPassword, middlewares.JWTMiddleware())
	e.PUT("/profile/email", handler.NewEmail, middlewares.JWTMiddleware())
	e.DELETE("/profile/delete", handler.Delete, middlewares.JWTMiddleware())
}

func (user *Delivery) SeeProfile(c echo.Context) error {
	userid := middlewares.ExtractToken(c)
	core, err := user.FromTo.MyProfile(userid)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	response := CoreToResponse(core)
	return c.JSON(http.StatusOK, helper.SuccessGet("Your Profile", response))
}

func (user *Delivery) NewImages(c echo.Context) error {
	userid := middlewares.ExtractToken(c)

	dataFoto, infoFoto, fotoerr := c.Request().FormFile("foto")

	if fotoerr == http.ErrMissingFile || fotoerr != nil {
		return c.JSON(http.StatusInternalServerError, helper.Failed("Failed TO Upload Image"))
	}

	format, errf := helper.CheckFile(infoFoto.Filename)
	if errf != nil {
		return c.JSON(http.StatusBadRequest, helper.Failed("Format Error"))
	}
	waktu := fmt.Sprintf("%v", time.Now())
	imageName := strconv.Itoa(userid) + "_" + waktu + "." + format

	imageaddress, errupload := helper.UploadFileToS3(config.FolderName, imageName, config.FileType, dataFoto)
	if errupload != nil {
		fmt.Println(errupload)
		return c.JSON(http.StatusInternalServerError, helper.Failed("failed to upload file"))
	}

	msg, err := user.FromTo.UpdateImage(imageaddress, userid)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.Failed("Terjadi Kesalahan"))
	}

	return c.JSON(http.StatusOK, helper.Success(msg))
}

func (user *Delivery) NewUserName(c echo.Context) error {
	userid := middlewares.ExtractToken(c)
	var bind UserNameUpdate
	errbind := c.Bind(&bind)
	if errbind != nil {
		return c.JSON(http.StatusBadRequest, helper.Failed("Gagal bind data"))
	}
	core := UserNameToCore(bind)
	msg, err := user.FromTo.UpdateUserName(core, userid)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.Failed("Terjadi kesalahan"))
	}

	return c.JSON(http.StatusOK, helper.Success(msg))
}

func (user *Delivery) NewPassword(c echo.Context) error {
	userid := middlewares.ExtractToken(c)
	var bind PasswordUpdate
	errbind := c.Bind(&bind)
	if errbind != nil {
		return c.JSON(http.StatusBadRequest, helper.Failed("Gagal bind data"))
	}
	core := PasswordToCore(bind)
	msg, err := user.FromTo.UpdatePassword(core, userid)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.Failed("Terjadi kesalahan"))
	}

	return c.JSON(http.StatusOK, helper.Success(msg))
}

func (user *Delivery) NewEmail(c echo.Context) error {
	userid := middlewares.ExtractToken(c)
	var bind EmailUpdate
	errbind := c.Bind(&bind)
	if errbind != nil {
		return c.JSON(http.StatusBadRequest, helper.Failed("Gagal bind data"))
	}
	core := EmailToCore(bind)
	msg, err := user.FromTo.UpdateEmail(core, userid)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.Failed("Terjadi kesalahan"))
	}

	return c.JSON(http.StatusOK, helper.Success(msg))
}

func (user *Delivery) NewName(c echo.Context) error {
	userid := middlewares.ExtractToken(c)
	var bind NameUpdate
	errbind := c.Bind(&bind)
	if errbind != nil {
		return c.JSON(http.StatusBadRequest, helper.Failed("Gagal bind data"))
	}
	core := NameToCore(bind)
	msg, err := user.FromTo.UpdateName(core, userid)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.Failed("Terjadi kesalahan"))
	}

	return c.JSON(http.StatusOK, helper.Success(msg))
}

func (user *Delivery) Delete(c echo.Context) error {
	userid := middlewares.ExtractToken(c)
	msg, err := user.FromTo.Delete(userid)
	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.Failed("Terjadi Kesalahan"))
	}

	return c.JSON(http.StatusOK, helper.Success(msg))
}

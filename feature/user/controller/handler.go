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
	e.PUT("/profile", handler.Update, middlewares.JWTMiddleware())
	e.DELETE("/profile", handler.Delete, middlewares.JWTMiddleware())
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

func (user *Delivery) Update(c echo.Context) error {
	userid := middlewares.ExtractToken(c)
	var userreq Request
	err := c.Bind(&userreq)
	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.Failed("Gagal Membind data"))
	}

	dataFoto, infoFoto, fotoerr := c.Request().FormFile("foto")

	if fotoerr != http.ErrMissingFile || fotoerr == nil {
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
		userreq.Foto = imageaddress
	}
	core := userreq.RequestToCore()
	msg, err := user.FromTo.Update(core, userid)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.Failed("Terjadi Kesalahan"))
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

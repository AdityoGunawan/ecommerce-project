package controller

import (
	"ecommerce-project/feature/auth/entities"
	"ecommerce-project/utils/helper"
	"net/http"

	"github.com/labstack/echo/v4"
)

type Delivery struct {
	FromTo entities.ServiceInterface
}

func New(e *echo.Echo, data entities.ServiceInterface) {
	handler := &Delivery{
		FromTo: data,
	}

	e.POST("/register", handler.Register)
	e.GET("/login", handler.Login)
}

func (user *Delivery) Register(c echo.Context) error {
	var request RegisterReq
	errbind := c.Bind(&request)
	if errbind != nil {
		return c.JSON(http.StatusBadRequest, helper.Failed("Failed to Bind Data"))
	}
	core := request.ReqToCore()
	row, err := user.FromTo.FirstRegister(core)
	if err != nil || row != 1 {
		return c.JSON(http.StatusInternalServerError, helper.Failed("Terjadi kesalahan"))
	}

	return c.JSON(http.StatusOK, helper.Success("Berhasil Mendaftar"))
}

func (user *Delivery) Login(c echo.Context) error {
	var login LoginReq
	errbind := c.Bind(&login)
	if errbind != nil {
		return c.JSON(http.StatusBadRequest, helper.Failed("Gagal Bind Data"))
	}
	core := login.ReqToCoreLogin()
	token, err := user.FromTo.Login(core.Email, core.Password)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.Failed("Password Tidak sesuai"))
	}

	return c.JSON(http.StatusOK, helper.SuccessGet("Berhasil Login", token))
}

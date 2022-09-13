package controller

import (
	"ecommerce-project/config"
	"ecommerce-project/feature/product/entities"
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

	e.POST("/product", handler.AddProduct, middlewares.JWTMiddleware())
	e.GET("/product", handler.Get8All)
	e.GET("/profile/product", handler.GetMyProduct, middlewares.JWTMiddleware())
	e.DELETE("/product:id", handler.DeleteProduct, middlewares.JWTMiddleware())
}

func (user *Delivery) AddProduct(c echo.Context) error {
	userid := middlewares.ExtractToken(c)
	var request Request
	errbind := c.Bind(&request)
	if errbind != nil {
		return c.JSON(http.StatusBadRequest, helper.Failed("Failed Bind Data"))
	}
	dataFoto, infoFoto, fotoerr := c.Request().FormFile("foto")
	if fotoerr == http.ErrMissingFile || fotoerr != nil {
		return c.JSON(http.StatusInternalServerError, helper.Failed("Failed TO Upload Image"))
	}

	format, errf := helper.CheckFile(infoFoto.Filename)
	if errf != nil {
		return c.JSON(http.StatusBadRequest, helper.Failed("Format Error"))
	}
	//
	err_image_size := helper.CheckSize(infoFoto.Size)
	if err_image_size != nil {
		return c.JSON(http.StatusBadRequest, err_image_size)
	}
	//
	waktu := fmt.Sprintf("%v", time.Now())
	imageName := strconv.Itoa(userid) + "_" + request.Name + waktu + "." + format

	imageaddress, errupload := helper.UploadFileToS3(config.FolderName, imageName, config.FileType, dataFoto)
	if errupload != nil {
		fmt.Println(errupload)
		return c.JSON(http.StatusInternalServerError, helper.Failed("failed to upload file"))
	}

	core := request.ReqToCore(imageaddress, uint(userid))

	row, erri := user.FromTo.AddProductI(core)
	if erri != nil || row != 1 {
		return c.JSON(http.StatusInternalServerError, helper.Failed("Gagal ke database"))
	}

	return c.JSON(http.StatusCreated, helper.Success("Sukses menambahkan"))
}

func (user *Delivery) Get8All(c echo.Context) error {

	page, _ := strconv.Atoi(c.QueryParam("page"))

	listcore, err := user.FromTo.GetAll(page)
	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.Failed("Terjadi Kesalahan"))
	}
	listRes := CoreToResponseList(listcore)

	return c.JSON(http.StatusOK, helper.SuccessGet("Sukses mendapatkan data", listRes))
}

func (user *Delivery) DeleteProduct(c echo.Context) error {
	userid := middlewares.ExtractToken(c)
	productid, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusOK, helper.Failed("Id hanya bisa nomor"))
	}
	msg, errs := user.FromTo.Delete(userid, productid)
	if errs != nil {
		return c.JSON(http.StatusInternalServerError, helper.Failed(msg))
	}

	return c.JSON(http.StatusOK, helper.Success(msg))
}

func (user *Delivery) GetMyProduct(c echo.Context) error {
	userid := middlewares.ExtractToken(c)
	mycore, err := user.FromTo.GetMyProduct(userid)
	if err != nil {
		c.JSON(http.StatusInternalServerError, helper.Failed("Terjadi Kesalahan"))
	}
	myres := CoreToResponseList(mycore)

	return c.JSON(http.StatusOK, helper.SuccessGet("Sukses mendapatkan data", myres))
}

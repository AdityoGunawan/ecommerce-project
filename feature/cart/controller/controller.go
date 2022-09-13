package controller

import (
	"ecommerce-project/feature/cart/repository"
	"ecommerce-project/feature/cart/services"
	"ecommerce-project/middlewares"
	"ecommerce-project/models"
	"ecommerce-project/response"
	"net/http"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

// controller untuk membuat cart
func CreateCartControllers(c echo.Context) error {

	Cart := models.Cart{}
	c.Bind(&Cart)
	v := validator.New()
	e := v.Var(Cart.Quantity, "required,gt=0")
	if e == nil {
		logged := middlewares.ExtractToken(c)

		id_user_cart, _ := repository.GetIDUserProduct(int(Cart.ProductID))
		harga_product, _ := repository.GetHargaProduct(int(Cart.ProductID))

		Cart.UserID = uint(logged)
		Cart.TotalPrice = Cart.Quantity * harga_product

		if uint(logged) == id_user_cart {
			return c.JSON(http.StatusBadRequest, response.AccessForbiddenResponse())
		}
		_, e = services.CreateCart(&Cart)
	}
	if e != nil {
		return c.JSON(http.StatusBadRequest, response.BadRequestResponse())
	}
	return c.JSON(http.StatusOK, response.SuccessResponseNonData())
}

// controller untuk mendapatkan semua cart by id user
func GetAllCartControllers(c echo.Context) error {
	logged := middlewares.ExtractToken(c)
	cart, err := repository.GetAllCart(logged)
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.BadRequestResponse())
	}
	return c.JSON(http.StatusOK, response.SuccessResponseData(cart))
}

// controller untuk menghapus cart by id
func DeleteCartControllers(c echo.Context) error {
	id := c.Param("id")
	conv_id, err := strconv.Atoi(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.FalseParamResponse())
	}
	id_user_cart, _, _ := repository.GetIDUserCart(conv_id)
	logged := middlewares.ExtractToken(c)
	if uint(logged) != id_user_cart {
		return c.JSON(http.StatusBadRequest, response.AccessForbiddenResponse())
	}
	repository.DeleteCart(conv_id)
	return c.JSON(http.StatusOK, response.SuccessResponseNonData())
}

func UpdateCartControllers(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.FalseParamResponse())
	}

	cart := models.Cart{}
	c.Bind(&cart)
	v := validator.New()
	e := v.Var(cart.Quantity, "required,gt=0")
	if e == nil {
		return c.JSON(http.StatusBadRequest, response.BadRequestResponse())
	}

	// mengecek user id nya sama dan ada pada tabel
	id_user_cart, id_product, _ := repository.GetIDUserCart(id)
	logged := middlewares.ExtractToken(c)
	if id_user_cart == 0 {
		return c.JSON(http.StatusBadRequest, response.BadRequestResponse())
	} else if uint(logged) != id_user_cart {
		return c.JSON(http.StatusBadRequest, response.AccessForbiddenResponse())
	}

	// mengupdate total harga
	harga_product, _ := repository.GetHargaProduct(int(id_product))
	cart.TotalPrice = cart.Quantity * harga_product

	// untuk mengupdate
	repository.UpdateCart(id, &cart)

	return c.JSON(http.StatusOK, response.SuccessResponseNonData())
}

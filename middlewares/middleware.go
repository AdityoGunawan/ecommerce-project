package middlewares

import (
	// "clean/config"
	"ecommerce-project/config"
	"strings"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func JWTMiddleware() echo.MiddlewareFunc {
	return middleware.JWTWithConfig(middleware.JWTConfig{
		SigningMethod: "HS256",
		SigningKey:    []byte(config.SECRET),
	})
}

func CreateToken(userId int) (string, error) {
	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["userId"] = userId                             //menyimpan userid ke dalam token
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix() //token akan expired setelah 1 jam
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// fmt.Println("token:", token)
	return token.SignedString([]byte(config.SECRET))
}

func ExtractToken(c echo.Context) int {

	headerData := c.Request().Header.Get("Authorization")
	dataAuth := strings.Split(headerData, " ")
	token := dataAuth[len(dataAuth)-1]
	datajwt, _ := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		return []byte(config.SECRET), nil
	})
	// fmt.Println("data:", datajwt)
	if datajwt.Valid {
		claims := datajwt.Claims.(jwt.MapClaims)
		// fmt.Println("calaims:", claims)
		userId := claims["userId"].(float64)
		// fmt.Println(userId)
		return int(userId)
	}

	return -1
}

package factory

import (
	controllerauth "ecommerce-project/feature/auth/controller"
	repositoryauth "ecommerce-project/feature/auth/repository"
	servicesauth "ecommerce-project/feature/auth/service"
	controller "ecommerce-project/feature/product/controller"
	repository "ecommerce-project/feature/product/repository"
	services "ecommerce-project/feature/product/services"
	controlleruser "ecommerce-project/feature/user/controller"
	repositoryuser "ecommerce-project/feature/user/repository"
	servicesuser "ecommerce-project/feature/user/services"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func InitFact(e *echo.Echo, db *gorm.DB) {
	authrepo := repositoryauth.New(db)
	authserv := servicesauth.New(authrepo)
	controllerauth.New(e, authserv)

	userrepo := repositoryuser.New(db)
	userserv := servicesuser.New(userrepo)
	controlleruser.New(e, userserv)

	fromrepo := repository.New(db)
	fromserv := services.New(fromrepo)
	controller.New(e, fromserv)
}

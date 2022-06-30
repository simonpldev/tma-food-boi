package route

import (
	"github.com/labstack/echo/v4"
	"tma-food-boi/controller"
	"tma-food-boi/validator"
)

func user(e *echo.Echo) {
	g := e.Group("/users")

	g.GET("", controller.UserList)

	g.POST("", controller.UserCreate, validator.UserCreate)
}

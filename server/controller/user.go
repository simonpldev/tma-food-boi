package controller

import (
	"github.com/labstack/echo/v4"
	"tma-food-boi/model"
	"tma-food-boi/service"
	"tma-food-boi/util"
)

func UserCreate(c echo.Context) error {
	var (
		payload = c.Get("body").(model.UserCreatePayload)
	)

	rs, err := service.UserCreate(payload)
	if err != nil {
		return util.Response404(c, err.Error())
	}

	return util.Response200(c, "", rs)
}

func UserList(c echo.Context) error {

	rs, err := service.UserList()

	if err != nil {
		return util.Response404(c, err.Error())
	}

	return util.Response200(c, "", rs)
}

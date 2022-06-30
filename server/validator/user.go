package validator

import (
	"github.com/asaskevich/govalidator"
	"github.com/labstack/echo/v4"
	"tma-food-boi/model"
	"tma-food-boi/util"
)

func UserCreate(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		var (
			payload model.UserCreatePayload
		)

		if err := c.Bind(&payload); err != nil {
			return util.Response400(c, err.Error())
		}

		_, err := govalidator.ValidateStruct(payload)
		if err != nil {
			return util.Response400(c, err.Error())
		}

		c.Set("body", payload)
		return next(c)
	}
}

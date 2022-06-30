package route

import "github.com/labstack/echo/v4"

func Bootstrap(e *echo.Echo) {
	user(e)
}

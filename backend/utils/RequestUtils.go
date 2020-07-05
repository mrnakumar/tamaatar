package utils

import "github.com/labstack/echo/v4"

type RequestUtils struct {
}

func (c RequestUtils) GetHeader(ctx echo.Context) (header string, found bool) {
	found = false
	header = ctx.Request().Header.Get("uid")
	if len(header) > 0 {
		found = true
	} else {
		found = false
	}
	return header, found
}

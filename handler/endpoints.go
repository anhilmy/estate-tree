package handler

import (
	"github.com/labstack/echo/v4"
)

func (s *Server) PostEstate(ctx echo.Context) error {
	est := new(PostEstateJSONRequestBody)
	if err := ctx.Bind(est); err != nil {
		return ctx.JSON(400, BadRequestError(err.Error()))
	}

	if err := est.Validate(); err != nil {
		return ctx.JSON(400, err.Error())
	}
	return nil
}

package handler

import (
	"github.com/SawitProRecruitment/UserService/generated"
	"github.com/labstack/echo/v4"
)

func (s *Server) PostEstate(ctx echo.Context) error {
	est := new(generated.PostEstateJSONRequestBody)
	if err := ctx.Bind(est); err != nil {
		return ctx.JSON(BadRequestError(err.Error()))
	}
	return nil
}

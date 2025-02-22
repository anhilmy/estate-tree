package handler

import (
	"errors"

	"github.com/SawitProRecruitment/UserService/generated"
	"github.com/labstack/echo/v4"
)

type ErrorResponse struct {
	generated.ErrorResponse
}

func (e ErrorResponse) Error() string {
	return e.Message
}

func BadRequestError(ctx echo.Context, err error) error {
	return ctx.JSON(400, ErrorResponse{
		generated.ErrorResponse{
			Message: err.Error(),
		},
	})
}

func InternalServerError(ctx echo.Context, err error) error {
	return ctx.JSON(500, ErrorResponse{
		generated.ErrorResponse{
			Message: err.Error(),
		},
	})
}

func NotFoundError(ctx echo.Context, err error) error {
	return ctx.JSON(404, ErrorResponse{
		generated.ErrorResponse{
			Message: err.Error(),
		},
	})
}
type SuccessResponse struct {
}

func SuccessCreateResponse(ctx echo.Context, data interface{}) error {
	return ctx.JSON(200, data)
}

func SuccessGetResponse(ctx echo.Context, data interface{}) error {
	return ctx.JSON(200, data)
}

type PostEstateJSONRequestBody struct {
	generated.PostEstateJSONRequestBody
}

func (req PostEstateJSONRequestBody) Validate() error {
	if req.Width < 1 || req.Width > 50000 {
		return errors.New("width is not a valid number")
	}

	if req.Length < 1 || req.Length > 50000 {
		return errors.New("length is not a valid number")
	}

	return nil
}

type PostEstateIdTreeJSONRequestBody struct {
	generated.PostEstateIdTreeJSONRequestBody
}

func (req PostEstateIdTreeJSONRequestBody) Validate() error {
	if req.Height < 1 || req.Height > 30 {
		return errors.New("height is not a valid number")
	}

	if req.X < 1 || req.Y < 1 {
		return errors.New("coordinates is out of maps bound")
	}

	return nil
}

package handler

import (
	"github.com/SawitProRecruitment/UserService/generated"
	"github.com/SawitProRecruitment/UserService/repository"
	"github.com/labstack/echo/v4"
)

func (s *Server) PostEstate(c echo.Context) error {
	ctx := c.Request().Context()

	req := new(generated.PostEstateJSONRequestBody)
	if err := c.Bind(req); err != nil {
		return BadRequestError(c, err)
	}

	body := PostEstateJSONRequestBody{
		generated.PostEstateJSONRequestBody(*req),
	}

	if err := body.Validate(); err != nil {
		return BadRequestError(c, err)
	}

	var res generated.UuidResponse
	var input repository.CreateEstateInput = repository.CreateEstateInput{
		Length: req.Length,
		Width:  req.Width,
	}
	uuid, err := s.Repository.InsertEstate(ctx, input)
	if err != nil {
		return InternalServerError(c, err)
	}

	res.Uuid = uuid.Uuid
	return SuccessCreateResponse(c, res)
}

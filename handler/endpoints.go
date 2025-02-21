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

// PostEstateIdTree implements generated.ServerInterface.
func (s *Server) PostEstateIdTree(c echo.Context, id string) error {
	ctx := c.Request().Context()

	req := new(generated.PostEstateIdTreeJSONRequestBody)
	if err := c.Bind(req); err != nil {
		return BadRequestError(c, err)
	}

	body := PostEstateIdTreeJSONRequestBody{
		generated.PostEstateIdTreeJSONRequestBody(*req),
	}

	if err := body.Validate(); err != nil {
		return BadRequestError(c, err)
	}

	var res generated.UuidResponse
	var input repository.CreateTreeInput = repository.CreateTreeInput{
		X:        req.X,
		Y:        req.Y,
		Height:   req.Height,
		EstateId: id,
	}

	uuid, err := s.Repository.InsertTree(ctx, input)
	if err != nil {
		return InternalServerError(c, err)
	}

	res.Uuid = uuid.Uuid
	return SuccessCreateResponse(c, res)
}

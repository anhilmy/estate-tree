package handler

import (
	"database/sql"
	"errors"

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

	estate, err := s.Repository.GetEstate(ctx, repository.UuidInput{
		Uuid: id,
	})
	if err == sql.ErrNoRows {
		err = errors.New("estate not found")
		return NotFoundError(c, err)
	} else if err != nil {
		return InternalServerError(c, err)
	} else if estate.Length < req.X || estate.Width < req.Y {
		return BadRequestError(c, errors.New("tree location cannot outside the estate"))
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

// GetEstateIdStats implements generated.ServerInterface.
func (s *Server) GetEstateIdStats(c echo.Context, id string) error {
	ctx := c.Request().Context()

	_, err := s.Repository.GetEstate(ctx, repository.UuidInput{
		Uuid: id,
	})
	if err == sql.ErrNoRows {
		err = errors.New("estate not found")
		return NotFoundError(c, err)
	} else if err != nil {
		return InternalServerError(c, err)
	}

	stat, err := s.Repository.GetEstateStats(ctx, repository.UuidInput{
		Uuid: id,
	})
	if err != nil {
		return InternalServerError(c, err)
	}

	var res generated.EstateStatResponse = generated.EstateStatResponse{
		Count:  stat.Count,
		Max:    stat.Max,
		Median: stat.Median,
		Min:    stat.Min,
	}
	return SuccessGetResponse(c, res)
}

package handler

import (
	"database/sql"
	"errors"
	"fmt"
	"net/http"

	"github.com/SawitProRecruitment/UserService/generated"
	"github.com/SawitProRecruitment/UserService/repository"
	"github.com/labstack/echo/v4"
)

func (s *Server) GetHello(ctx echo.Context, params generated.GetHelloParams) error {
	var resp generated.HelloResponse
	resp.Message = fmt.Sprintf("Hello User %d", params.Id)
	return ctx.JSON(http.StatusOK, resp)
}

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

	res.Id = uuid.Uuid
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

	res.Id = uuid.Uuid
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

// GetEstateIdDronePlan implements generated.ServerInterface.
func (s *Server) GetEstateIdDronePlan(c echo.Context, id string, params generated.GetEstateIdDronePlanParams) error {
	ctx := c.Request().Context()

	estate, err := s.Repository.GetEstate(ctx, repository.UuidInput{
		Uuid: id,
	})
	if err == sql.ErrNoRows {
		err = errors.New("estate not found")
		return NotFoundError(c, err)
	} else if err != nil {
		return InternalServerError(c, err)
	}

	var res generated.DronePlanResponse
	var input repository.UuidInput = repository.UuidInput{
		Uuid: id,
	}

	allTree, err := s.Repository.GetAllTree(ctx, input)
	if err != nil {
		return InternalServerError(c, err)
	}

	queue := make([]repository.TreeModel, 0)
	reversedQueue := make([]repository.TreeModel, 0) // will be always inserted from index 0

	for _, tree := range allTree {
		if tree.Y%2 == 0 {
			reversedQueue = append([]repository.TreeModel{tree}, reversedQueue...)
		} else {
			if len(reversedQueue) > 0 {
				queue = append(queue, reversedQueue...)
				reversedQueue = make([]repository.TreeModel, 0)
			}
			queue = append(queue, tree)
		}
	}

	allDeltaHeight := 0
	var lastHeight repository.TreeModel = repository.TreeModel{}
	// rest := generated.DronePlanResponseRest{}

	for _, currHeight := range queue {
		deltaHeight := 0
		if lastHeight.Height > currHeight.Height {
			deltaHeight = lastHeight.Height - currHeight.Height
		} else {
			deltaHeight = currHeight.Height - lastHeight.Height
		}
		fmt.Println(deltaHeight)
		allDeltaHeight = allDeltaHeight + deltaHeight
		// max distance is not 0 && rest is initiated && current
		lastHeight = currHeight
	}
	allDeltaHeight += lastHeight.Height

	// +2 from drone need fly above 1 meters, calculating on takeoff and on landing
	res.Distance = allDeltaHeight + ((estate.Length * estate.Width * 10) - 10) + 2

	return SuccessGetResponse(c, res)
}

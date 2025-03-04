package handler

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/SawitProRecruitment/UserService/generated"
	"github.com/SawitProRecruitment/UserService/repository"
	"github.com/golang/mock/gomock"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestPostEstate(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockRepo := repository.NewMockRepositoryInterface(mockCtrl)
	// server := NewServer(NewServerOptions{
	// 	Repository: mockRepo,

	// })

	testcase := []struct {
		name    string
		req     *generated.PostEstateJSONRequestBody
		res     *generated.UuidResponse
		repoIn  *repository.CreateEstateInput
		repoOut *repository.UuidOutput
		repoErr error
		err     error
		status  int
	}{
		{
			name:    "normal",
			req:     &generated.PostEstateRequest{Length: 10, Width: 200},
			res:     &generated.UuidResponse{Id: "uuid-uuid"},
			repoIn:  &repository.CreateEstateInput{Length: 10, Width: 200},
			repoOut: &repository.UuidOutput{Uuid: "uuid-uuid"},
			repoErr: nil,
			err:     nil,
			status:  200,
		},
		{
			name: "width over",
			req: &generated.PostEstateRequest{
				Length: 99,
				Width:  100000,
			},
			res:     &generated.UuidResponse{},
			repoIn:  nil,
			repoOut: nil,
			repoErr: nil,
			err:     errors.New("width is not a valid number"),
			status:  400,
		},
		{
			name: "length under",
			req: &generated.PostEstateRequest{
				Length: -100,
				Width:  90,
			},
			res:     &generated.UuidResponse{},
			repoIn:  nil,
			repoOut: nil,
			repoErr: nil,
			err:     errors.New("length is not a valid number"),
			status:  400,
		},
		{
			name:    "empty",
			req:     &generated.PostEstateRequest{},
			res:     &generated.UuidResponse{},
			repoIn:  nil,
			repoOut: nil,
			repoErr: nil,
			err:     errors.New("width is not a valid number"),
			status:  400,
		},
	}

	e := echo.New()
	for _, tc := range testcase {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			if tc.repoIn != nil {
				mockRepo.EXPECT().InsertEstate(gomock.Any(), *tc.repoIn).Return(*tc.repoOut, tc.repoErr)
			} else {
				mockRepo.EXPECT().InsertEstate(gomock.Any(), gomock.Nil()).Return(repository.UuidOutput{}, tc.repoErr).Times(0)
			}

			reqbody, _ := json.Marshal(tc.req)
			req := httptest.NewRequest(http.MethodPost, "/estate", bytes.NewBuffer(reqbody))
			req.Header.Set("Content-Type", "application/json")

			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)

			server := NewServer(NewServerOptions{
				Repository: mockRepo,
			})
			err := server.PostEstate(c)
			if err != nil {
				t.Fail()
			}

			assert.Equal(t, tc.status, rec.Code)
			if tc.err != nil {
				var response generated.ErrorResponse
				json.Unmarshal(rec.Body.Bytes(), &response)
				assert.Equal(t, tc.err.Error(), response.Message)
			} else {
				var response generated.UuidResponse
				json.Unmarshal(rec.Body.Bytes(), &response)
				assert.Equal(t, *tc.res, response)
			}
		})
	}
}

func TestPostEstateIdTree(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockRepo := repository.NewMockRepositoryInterface(mockCtrl)

	estateInput := repository.UuidInput{
		Uuid: "123456ESTATE",
	}
	estateOutput := repository.EstateModel{
		Uuid:   "123456ESTATE",
		Length: 30,
		Width:  200,
	}
	insertOutput := repository.UuidOutput{
		Uuid: "123456TREE",
	}

	testcase := []struct {
		name      string
		req       generated.PostEstateIdTreeJSONRequestBody
		res       generated.UuidResponse
		estateIn  repository.UuidInput
		estateOut repository.EstateModel
		estateErr error
		insertIn  repository.CreateTreeInput
		insertOut repository.UuidOutput
		insertErr error
		err       error //resposne error
		status    int
	}{
		{
			name: "normal",
			req: generated.PostEstateIdTreeJSONRequestBody{
				Height: 30,
				X:      20,
				Y:      150,
			},
			res: generated.UuidResponse{
				Id: "123456TREE",
			},
			estateIn:  estateInput,
			estateOut: estateOutput,
			estateErr: nil,
			insertIn: repository.CreateTreeInput{
				X:        20,
				Y:        150,
				Height:   30,
				EstateId: "123456ESTATE",
			},
			insertOut: insertOutput,
			insertErr: nil,
			err:       nil,
			status:    200,
		},
		{
			name: "coordinate over",
			req: generated.PostEstateIdTreeJSONRequestBody{
				Height: 25,
				X:      45,
				Y:      250,
			},
			res:       generated.UuidResponse{},
			estateIn:  estateInput,
			estateOut: estateOutput,
			estateErr: nil,
			insertIn:  repository.CreateTreeInput{},
			insertOut: repository.UuidOutput{},
			insertErr: nil,
			err:       errors.New("tree location cannot outside the estate"),
			status:    400,
		},
		{
			name: "height over",
			req: generated.PostEstateIdTreeJSONRequestBody{
				Height: 50,
				X:      23,
				Y:      140,
			},
			res:       generated.UuidResponse{},
			estateIn:  repository.UuidInput{},
			estateOut: repository.EstateModel{},
			estateErr: nil,
			insertIn:  repository.CreateTreeInput{},
			insertOut: repository.UuidOutput{},
			insertErr: nil,
			err:       errors.New("height is not a valid number"),
			status:    400,
		},
		{
			name: "estate not found",
			req: generated.PostEstateIdTreeJSONRequestBody{
				Height: 23,
				X:      25,
				Y:      100,
			},
			res:       generated.UuidResponse{},
			estateIn:  estateInput,
			estateOut: repository.EstateModel{},
			estateErr: sql.ErrNoRows,
			insertIn:  repository.CreateTreeInput{},
			insertOut: repository.UuidOutput{},
			insertErr: nil,
			err:       errors.New("estate not found"),
			status:    404,
		},
	}

	e := echo.New()
	for _, tc := range testcase {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			estateExp := mockRepo.EXPECT().GetEstate(gomock.Any(), tc.estateIn).Return(tc.estateOut, tc.estateErr)
			treeExp := mockRepo.EXPECT().InsertTree(gomock.Any(), tc.insertIn).Return(tc.insertOut, tc.insertErr)
			if (tc.estateIn == repository.UuidInput{}) {
				estateExp.Times(0)
			}
			if (tc.insertIn == repository.CreateTreeInput{}) {
				treeExp.Times(0)
			}

			reqbody, _ := json.Marshal(tc.req)
			req := httptest.NewRequest(http.MethodPost, "/estate/1234/tree", bytes.NewBuffer(reqbody))
			req.Header.Set("Content-Type", "application/json")

			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)

			server := NewServer(NewServerOptions{
				Repository: mockRepo,
			})
			err := server.PostEstateIdTree(c, estateInput.Uuid)
			if err != nil {
				t.Fail()
			}

			assert.Equal(t, tc.status, rec.Code)
			if tc.err != nil {
				var response generated.ErrorResponse
				json.Unmarshal(rec.Body.Bytes(), &response)
				assert.Equal(t, tc.err.Error(), response.Message)
			} else {
				var response generated.UuidResponse
				json.Unmarshal(rec.Body.Bytes(), &response)
				assert.Equal(t, tc.res, response)
			}
		})
	}
}

func TestGetEstateIdStat(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockRepo := repository.NewMockRepositoryInterface(mockCtrl)

	estateInput := repository.UuidInput{
		Uuid: "123456ESTATE",
	}
	estateOutput := repository.EstateModel{
		Uuid:   "123456ESTATE",
		Length: 30,
		Width:  200,
	}
	statOutput := repository.EstateStatsOutput{
		Count:  15,
		Max:    39,
		Min:    12,
		Median: 23,
	}

	testcase := []struct {
		name       string
		res        generated.EstateStatResponse
		estateIn   repository.UuidInput
		estateOut  repository.EstateModel
		estateErr  error
		statOutput repository.EstateStatsOutput
		statError  error
		err        error //resposne error
		status     int
	}{
		{
			name: "normal",
			res: generated.EstateStatResponse{
				Count:  statOutput.Count,
				Max:    statOutput.Max,
				Median: statOutput.Median,
				Min:    statOutput.Min,
			},
			estateIn:   estateInput,
			estateOut:  estateOutput,
			estateErr:  nil,
			statOutput: statOutput,
			statError:  nil,
			err:        nil,
			status:     200,
		},
		{
			name:       "no estate",
			res:        generated.EstateStatResponse{},
			estateIn:   estateInput,
			estateOut:  repository.EstateModel{},
			estateErr:  sql.ErrNoRows,
			statOutput: repository.EstateStatsOutput{},
			statError:  nil,
			err:        nil,
			status:     404,
		},
		{
			name: "no tree",
			res: generated.EstateStatResponse{
				Count:  0,
				Max:    0,
				Median: 0,
				Min:    0,
			},
			estateIn:  estateInput,
			estateOut: estateOutput,
			estateErr: nil,
			statOutput: repository.EstateStatsOutput{
				Count:  0,
				Max:    0,
				Min:    0,
				Median: 0,
			},
			statError: nil,
			err:       nil,
			status:    200,
		},
	}

	e := echo.New()
	for _, tc := range testcase {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			mockRepo.EXPECT().GetEstate(gomock.Any(), tc.estateIn).Return(tc.estateOut, tc.estateErr)
			stat := mockRepo.EXPECT().GetEstateStats(gomock.Any(), tc.estateIn).Return(tc.statOutput, tc.statError)

			if tc.estateErr != nil {
				stat.Times(0)
			}

			req := httptest.NewRequest(http.MethodPost, "/estate/1234/stats", nil)
			req.Header.Set("Content-Type", "application/json")

			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)

			server := NewServer(NewServerOptions{
				Repository: mockRepo,
			})
			err := server.GetEstateIdStats(c, tc.estateIn.Uuid)
			if err != nil {
				t.Fail()
			}

			assert.Equal(t, tc.status, rec.Code)
			if tc.err != nil {
				var response generated.ErrorResponse
				json.Unmarshal(rec.Body.Bytes(), &response)
				assert.Equal(t, tc.err.Error(), response.Message)
			} else {
				var response generated.EstateStatResponse
				json.Unmarshal(rec.Body.Bytes(), &response)
				assert.Equal(t, tc.res, response)
			}
		})
	}
}

func TestGetDronePlan(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockRepo := repository.NewMockRepositoryInterface(mockCtrl)

	estateInput := repository.UuidInput{
		Uuid: "123456ESTATE",
	}
	estateOutput := repository.EstateModel{
		Uuid:   "123456ESTATE",
		Length: 5,
		Width:  1,
	}

	treeOutput := []repository.TreeModel{
		{
			Uuid:       "1TREE",
			X:          2,
			Y:          1,
			Height:     10,
			EstateUuid: "123456ESTATE",
		},

		{
			Uuid:       "2TREE",
			X:          3,
			Y:          1,
			Height:     20,
			EstateUuid: "123456ESTATE",
		},

		{
			Uuid:       "3TREE",
			X:          4,
			Y:          1,
			Height:     10,
			EstateUuid: "123456ESTATE",
		},
		// {
		// 	Uuid:       "6TREE",
		// 	X:          1,
		// 	Y:          2,
		// 	Height:     2,
		// 	EstateUuid: "123456ESTATE",
		// },
		// {
		// 	Uuid:       "5TREE",
		// 	X:          2,
		// 	Y:          2,
		// 	Height:     8,
		// 	EstateUuid: "123456ESTATE",
		// },
		// {
		// 	Uuid:       "4TREE",
		// 	X:          3,
		// 	Y:          2,
		// 	Height:     15,
		// 	EstateUuid: "123456ESTATE",
		// },
		// {
		// 	Uuid:       "7TREE",
		// 	X:          2,
		// 	Y:          3,
		// 	Height:     2,
		// 	EstateUuid: "123456ESTATE",
		// },
		// {
		// 	Uuid:       "8TREE",
		// 	X:          4,
		// 	Y:          3,
		// 	Height:     2,
		// 	EstateUuid: "123456ESTATE",
		// },
	}

	testcase := []struct {
		name       string
		param      generated.GetEstateIdDronePlanParams
		res        generated.DronePlanResponse
		estateIn   repository.UuidInput
		estateOut  repository.EstateModel
		estateErr  error
		treeOutput []repository.TreeModel
		treeError  error
		err        error //resposne error
		status     int
	}{
		{
			name:  "normal no max distance",
			param: generated.GetEstateIdDronePlanParams{},
			res: generated.DronePlanResponse{
				// Distance: 174,
				Distance: 82,
			},
			estateIn:   estateInput,
			estateOut:  estateOutput,
			estateErr:  nil,
			treeOutput: treeOutput,
			treeError:  nil,
			err:        nil,
			status:     200,
		},
		// {
		// 	name: "normal with max distance",
		// },
		{
			name:       "estate not found",
			param:      generated.GetEstateIdDronePlanParams{},
			res:        generated.DronePlanResponse{},
			estateIn:   estateInput,
			estateOut:  repository.EstateModel{},
			estateErr:  sql.ErrNoRows,
			treeOutput: nil,
			treeError:  nil,
			err:        errors.New("estate not found"),
			status:     404,
		},
		// {
		// 	name: "max distance is 1",
		// },
	}

	e := echo.New()
	for _, tc := range testcase {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			mockRepo.EXPECT().GetEstate(gomock.Any(), tc.estateIn).Return(tc.estateOut, tc.estateErr)
			tree := mockRepo.EXPECT().GetAllTree(gomock.Any(), tc.estateIn).Return(tc.treeOutput, tc.treeError)

			if tc.estateErr != nil {
				tree.Times(0)
			}

			req := httptest.NewRequest(http.MethodPost, "/estate/1234/drone-plan", nil)
			req.Header.Set("Content-Type", "application/json")

			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)

			server := NewServer(NewServerOptions{
				Repository: mockRepo,
			})
			err := server.GetEstateIdDronePlan(c, tc.estateIn.Uuid, tc.param)
			if err != nil {
				t.Fail()
			}

			assert.Equal(t, tc.status, rec.Code)
			if tc.err != nil {
				var response generated.ErrorResponse
				json.Unmarshal(rec.Body.Bytes(), &response)
				assert.Equal(t, tc.err.Error(), response.Message)
			} else {
				var response generated.DronePlanResponse
				json.Unmarshal(rec.Body.Bytes(), &response)
				assert.Equal(t, tc.res, response)
			}
		})
	}
}

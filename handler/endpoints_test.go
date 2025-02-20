package handler

import (
	"bytes"
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
			res:     &generated.UuidResponse{Uuid: "uuid-uuid"},
			repoIn:  &repository.CreateEstateInput{Length: 10, Width: 200},
			repoOut: &repository.UuidOutput{Uuid: "uuid-uuid"},
			repoErr: nil,
			err:     nil,
			status:  201,
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

package handler

import "github.com/SawitProRecruitment/UserService/generated"

type ErrorResponse struct {
	generated.ErrorResponse
}

func (e ErrorResponse) Error() string {
	return e.Message
}

func BadRequestError(msg string) ErrorResponse {
	return ErrorResponse{
		generated.ErrorResponse{
			Message: msg,
		},
	}
}

type PostEstateJSONRequestBody generated.PostEstateJSONRequestBody

func (req PostEstateJSONRequestBody) Validate() error {
	if req.Width < 1 || req.Width > 50000 {
		return BadRequestError("width is not a valid number")
	}

	if req.Length < 1 || req.Length > 50000 {
		return BadRequestError("length is not a valid number")
	}

	return nil
}

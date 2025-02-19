package handler

import "github.com/SawitProRecruitment/UserService/generated"

type ErrorResponse struct {
	generated.ErrorResponse

	Code int
}

func (e ErrorResponse) Error() string {
	return e.Message
}

func BadRequestError(msg string) (int, generated.ErrorResponse) {
	return 400, generated.ErrorResponse{
		Message: msg,
	}
}

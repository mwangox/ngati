package model

type ApiErrorResponse struct {
	StatusCode int `json:"status_code"`
	StatusDescription string `json:"status_description"`
}

func NewApiErrorResponse(code int, message string) ApiErrorResponse {
	return ApiErrorResponse{
		StatusCode:        code,
		StatusDescription: message,
	}
}
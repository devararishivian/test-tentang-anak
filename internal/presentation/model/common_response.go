package model

type CommonResponse struct {
	Message string `json:"message"`
}

type CommonResponseWithErrors struct {
	Message string `json:"message"`
	Errors  any    `json:"errors"`
}

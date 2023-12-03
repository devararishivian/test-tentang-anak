package model

type CommonResponse struct {
	Message string `json:"message"`
}

type CommonResponseWithData struct {
	Message string `json:"message"`
	Data    any    `json:"data"`
}

type CommonResponseWithErrors struct {
	Message string `json:"message"`
	Errors  any    `json:"errors"`
}

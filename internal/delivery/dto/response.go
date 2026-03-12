package dto

type Response struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Data    any    `json:"data"`
}

type NoDataResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

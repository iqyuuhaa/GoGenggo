package types

type ErrorData struct {
	HttpCode  int
	ErrorCode string
}

type DefaultAPIResponse struct {
	Header DefaultAPIResponseHeader `json:"header"`
	Data   interface{}              `json:"data,omitempty"`
}

type DefaultAPIResponseHeader struct {
	Status      string  `json:"status"`
	StatusCode  int     `json:"status_code"`
	ProcessTime string  `json:"process_time"`
	Message     *string `json:"message,omitempty"`
	ErrorCode   *string `json:"error_code,omitempty"`
}

type DefaultAPISuccessResponse struct {
	IsSuccess bool `json:"is_success" example:"true"`
}

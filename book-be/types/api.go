package types

type ResponseSuccess struct {
	Status  string `json:"status"`
	Message string `json:"message"`
	Data    *any   `json:"data,omitempty"`
}

type ResponseError struct {
	Status  string `json:"status"`
	Message string `json:"message"`
	Error   any    `json:"error,omitempty"`
}

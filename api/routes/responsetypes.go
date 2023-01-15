package routes

type SuccessResponse struct {
	Message string `json:"message"`
}

type ErrorResponse struct {
	Reason string `json:"reason"`
}

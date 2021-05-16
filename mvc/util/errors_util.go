package util

type ApplicationError struct {
	StatusCode int    `json:"status"`
	Message    string `json:"message"`
	Code       string `json:"code"`
}

package models

import "github.com/google/uuid"

// Response is a struct that contains the information of a response
type Response struct {
	Ok      bool        `json:"ok"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

// ResponseError is a struct that contains the information of a response error
type ResponseError struct {
	Ok         bool   `json:"ok"`
	StatusCode int    `json:"status_code"`
	Message    string `json:"message"`
	Error      error  `json:"-"` // This is not serialized
	ErrorUUID  string `json:"error_uuid"`
}

// NewResponse is a function that creates a new response
// message is the message of the response
// data is the data of the response
func NewResponse(message string, data interface{}) *Response {
	return &Response{
		Ok:      true,
		Message: message,
		Data:    data,
	}
}

// NewResponseError is a function that creates a new response error
// statusCode is the status code of the error
// message is the message of the error
// error is the error
func NewResponseError(statusCode int, message string, error error) *ResponseError {
	uuid := uuid.New().String()
	return &ResponseError{
		Ok:         false,
		StatusCode: statusCode,
		Message:    message,
		Error:      error,
		ErrorUUID:  uuid,
	}
}

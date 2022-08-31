package helper

import "strings"

// Response is used for static shape json return
type Response struct {
	Status  bool        `json:"status"`
	Message string      `json:"message"`
	Errors  interface{} `json:"errors"`
	Data    interface{} `json:"data"`
}

// Empty object is used when data doesn't want to be null on json
type EmptyObj struct{}

// BuildSuccessResponse method is to inject data value to dynamic success response
func BuildSuccessResponse(message string, data interface{}) Response {
	res := Response{
		Status:  true,
		Message: message,
		Errors:  nil,
		Data:    data,
	}
	return res
}

// BuildFailedResponse method is to inject data value to dynamic failed response
func BuildFailedResponse(message string, err string, data interface{}) Response {
	splittedError := strings.Split(err, "\n")
	res := Response{
		Status:  false,
		Message: message,
		Errors:  splittedError,
		Data:    data,
	}
	return res
}

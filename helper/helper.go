package helper

import "github.com/go-playground/validator/v10"

type Response struct {
	Meta Meta        `json:"meta"`
	Data interface{} `json:"data"`
}
type Meta struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
	Status  string `json:"status"`
}

func APIresponse(message string, code int, status string, data interface{}) Response {
	Meta := Meta{
		Message: message,
		Code:    code,
		Status:  status,
	}
	json_response := Response{
		Meta: Meta,
		Data: data,
	}

	return json_response
}

func FormatError(err error) []string {
	var errors []string

	for _, e := range err.(validator.ValidationErrors) {
		errors = append(errors, e.Error())
	}

	return errors
}

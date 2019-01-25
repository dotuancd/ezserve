package errors

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gopkg.in/go-playground/validator.v8"
	"io"
	"net/http"
)

type Error struct {
	Ok bool `json:"ok"`
	Code string `json:"code"`
	Message string `json:"message"`
	Status int `json:"status"`
	Fields map[string]interface{} `json:"fields,omitempty"`
}

var (
	BadRequest = New("bad_request", "Bad Request", http.StatusBadRequest)

	NotFound = New("page_not_found", "Page Not Found", http.StatusNotFound)

	Forbidden = New("forbidden", "Forbidden", http.StatusForbidden)

	InternalServerError = New("internal_server_error", "Internal Server Error", http.StatusInternalServerError)
)

type HandlerFunc = func(*gin.Context) error

func Handler(next HandlerFunc) gin.HandlerFunc {
	return func(c *gin.Context) {
		err := next(c)

		if err == nil {
			return
		}

		if err == io.EOF {
			c.JSON(BadRequest.Status, BadRequest)
			return
		}

		validationErr, isValidationErr := err.(validator.ValidationErrors)
		if isValidationErr {
			validationErrorToResponse(c, validationErr)
			return
		}

		e, ok := err.(*Error)

		// Convert to Error in case receive an other error
		if !ok {
			e = toResponse(err)
		}

		c.JSON(e.Status, e)
	}
}

func toResponse(err error) *Error {
	return New("unknown", err.Error(), 500)
}

func (re *Error) Error() string {
	return fmt.Sprintf("[%s] %s %s", string(re.Status), re.Code, re.Message)
}

func New(code string, message string, status int) *Error {
	return &Error{Ok: false, Code:code, Message: message, Status: status}
}

func validationErrorToResponse(c *gin.Context, err validator.ValidationErrors) {

	status := http.StatusUnprocessableEntity
	e := New("validation_failed", "Validation fail", status)
	e.Fields = formatValidationErrors(err)

	c.JSON(status, e)
}

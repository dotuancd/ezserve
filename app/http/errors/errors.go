package errors

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Error struct {
	Code string `json:"code"`
	Message string `json:"message"`
	Status int `json:"status"`

}

type HandlerFunc = func(*gin.Context) error

func Handler(next HandlerFunc) gin.HandlerFunc {
	return func(c *gin.Context) {
		err := next(c)

		if err == nil {
			return
		}

		e, ok := err.(*Error)

		// Convert to Error in case receive an other error
		if !ok {
			e = convertToResponse(err)
		}

		c.JSON(e.Status, e)
	}
}

func convertToResponse(err error) (*Error) {
	return New("unknown", err.Error(), 500)
}

var (
	NotFound = New("page_not_found", "Page Not Found", http.StatusNotFound)

	Forbidden = New("forbidden", "Forbidden", http.StatusForbidden)

	InternalServerError = New("internal_server_error", "Internal Server Error", http.StatusInternalServerError)
)

func (re *Error) Error() string {
	return fmt.Sprintf("[%s] %s %s", string(re.Status), re.Code, re.Message)
}

func New(code string, message string, status int) *Error {
	return &Error{Code:code, Message: message, Status: status}
}

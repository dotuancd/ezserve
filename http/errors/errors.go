package errors

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Error struct {
	Code string
	Message string
	Status int
}

type ErrorHandlerFunc = func(*gin.Context) error

type HandlerFuncAdapter struct {
	Next func(*gin.Context) error
}

func (h *HandlerFuncAdapter) Handle() gin.HandlerFunc {
	return func(c *gin.Context) {
		_ = h.Next(c)
		//
		//if err != nil {
		//	//c.JSON()
		//}

		//return nil
	}
}

var (
	NotFound = &Error{Status: http.StatusNotFound, Code: "page_not_found", Message: "Page Not Found"}

	Forbidden = &Error{Status: http.StatusForbidden, Code: "forbidden", Message: "Forbidden"}

	InternalServerError = &Error{Status:http.StatusInternalServerError, Code: "internal_server_error", Message: "Internal Server Error"}

)

func (re *Error) Error() string {
	return fmt.Sprintf("[%s] %s %s", string(re.Status), re.Code, re.Message)
}

func HandleErrorFunc(next ErrorHandlerFunc) gin.HandlerFunc {
	return func(c *gin.Context) {
		err := next(c)

		if err == nil {
			return
		}

		if e, ok := err.(*Error); ok {
			c.JSON(e.Status, e)
		}
	}
}




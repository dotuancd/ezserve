package res

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Error struct {
	Context *gin.Context
	StatusCode int
	Content gin.H
}

type Success struct {
	context *gin.Context
	statusCode int
	content gin.H
}

func NewSuccess(c *gin.Context) *Success {
	return &Success{context: c, content: map[string]interface{}{"ok": true}, statusCode: 200}
}

func (s *Success) Message (message string) *Success {
	s.content["message"] = message
	return s
}

func (s *Success) Content() gin.H {
	return s.content
}

func (s *Success) Extra(key string, value interface{}) *Success {
	s.content[key] = value
	return s
}

func (s *Success) Status (status int) *Success {
	s.statusCode = status
	return s
}

func (s *Success) Send() {
	s.context.JSON(s.statusCode, s.content)
	s.context.Abort()
}

func NotFound(c *gin.Context) *Error {
	return NewError(c).
		Code("resource_not_found").
		Message("Resource not found").
		Status(http.StatusNotFound)
}

func Unauthorized(c *gin.Context) *Error {
	return NewError(c).
		Code("cannot_authorized").
		Message("Cannot authorized").
		Status(http.StatusUnauthorized)
}

func Forbidden(c *gin.Context) *Error {
	return NewError(c).
		Code("permission_required").
		Message("You haven't permission to perform the action").
		Status(http.StatusForbidden)
}

func NewError(c *gin.Context) *Error {
	return &Error{Context: c, Content: map[string]interface{}{"ok": false}, StatusCode: http.StatusInternalServerError}
}

func SendError(c *gin.Context, status int, code string, message string) {
	NewError(c).
		Status(status).
		Code(code).
		Message(message).
		Send()
}

func SendInternalError(c *gin.Context, code string, message string) {
	NewError(c).
		Code(code).
		Message(message).
		Send()
}

func (r *Error) Status(status int) *Error {
	r.StatusCode = status
	return r
}

func (r *Error) Code(code string) *Error {
	r.Content["code"] = code
	return r
}

func (r *Error) Message(message string) *Error {
	r.Content["message"] = message
	return r
}

func (r *Error) Send() {
	r.Context.AbortWithStatusJSON(r.StatusCode, r.Content)
}

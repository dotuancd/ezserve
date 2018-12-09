package res

import "github.com/gin-gonic/gin"

const (
	HttpBadRequest = 400
	HttpUnauthorized = 401
	HttpForbidden = 403
	HttpNotFound = 404
	HttpInternalError = 500
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
		Status(HttpNotFound)
}

func Unauthorized(c *gin.Context) *Error {
	return NewError(c).
		Code("cannot_authorized").
		Message("Cannot authorized").
		Status(HttpUnauthorized)
}

func Forbidden(c *gin.Context) *Error {
	return NewError(c).
		Code("permission_required").
		Message("You haven't permission to perform the action").
		Status(HttpForbidden)
}

func NewError(c *gin.Context) *Error {
	return &Error{Context: c, Content: map[string]interface{}{"ok": false}, StatusCode: HttpInternalError}
}

func SendError(c *gin.Context, status int, code string, message string) {
	NewError(c).
		Status(status).
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
	r.Context.JSON(r.StatusCode, r.Content)
	r.Context.Abort()
}

//
//func Error(c *gin.Context, status int, code string, message string)  {
//	c.JSON(status, gin.H{
//		"ok": false,
//		"code": code,
//		"error": message,
//	})
//	c.Abort()
//}

//func NewError(c *gin.Context, code string, message string) {
//	//c.JSON(404, gin.H{
//	//	"ok": false,
//	//	"code": code,
//	//	"error": message,
//	//})
//	//c.Abort()
//}

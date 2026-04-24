package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
	Errors  interface{} `json:"errors,omitempty"`
}

func Success(c *gin.Context, status int, message string, data interface{}) {
	c.JSON(status, Response{
		Success: true,
		Message: message,
		Data:    data,
	})
}

func Error(c *gin.Context, status int, message string, errs interface{}) {
	c.JSON(status, Response{
		Success: false,
		Message: message,
		Errors:  errs,
	})
}

func InternalError(c *gin.Context, message string) {
	Error(c, http.StatusInternalServerError, message, nil)
}

func BadRequest(c *gin.Context, message string, errs interface{}) {
	Error(c, http.StatusBadRequest, message, errs)
}

func NotFound(c *gin.Context, message string) {
	Error(c, http.StatusNotFound, message, nil)
}

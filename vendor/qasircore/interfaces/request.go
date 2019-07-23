package interfaces

import (
	"github.com/gin-gonic/gin"
	"github.com/thedevsaddam/govalidator"
)

type Request interface {
	Rules() map[string][]string
	Message() map[string][]string
	Authorize() bool
	Validation(rules govalidator.MapData, messages govalidator.MapData) []string
	SetContextGin(c *gin.Context)
	SetBody(body map[string]interface{})
	Header(key string) string
}

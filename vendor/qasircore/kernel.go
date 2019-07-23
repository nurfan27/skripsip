package qasircore

import (
	"qasircore/interfaces"
	"reflect"

	"github.com/gin-gonic/gin"
	"github.com/thedevsaddam/govalidator"
)

var Kernel = &kernel{}

type kernel struct {
	c        *gin.Context
	response ApiResponse
}

func (k *kernel) SetContext(c *gin.Context) {
	k.c = c
}

func (k *kernel) Handle(request interfaces.Request, controllerFunction interface{}) func(c *gin.Context) {
	var rules govalidator.MapData
	var messages govalidator.MapData

	if request == nil {
		request = &Request{}
	}

	return func(c *gin.Context) {
		if c == nil {
			c = k.c
		}
		reflectFunctionController := reflect.ValueOf(controllerFunction)

		request.SetContextGin(c)
		SetSubdomain(c)

		rules = request.Rules()
		messages = request.Message()

		if !request.Authorize() {
			c.JSON(403, map[string]interface{}{
				"errors": "unauthorized",
			})
			return
		}

		var errors []string

		if len(rules) > 0 {
			errors = request.Validation(rules, messages)
		}

		if errors != nil && len(errors) > 0 {
			k.response.SetStatus(STATUS_ERR_INVALID_PARAMETER)
			k.response.SetData(map[string]interface{}{
				"errors": errors,
			})
			c.JSON(400, k.response.ToResponse())
			return
		}
		reqInterface := reflect.ValueOf(request)
		req := reqInterface.Interface().(HttpRequest)

		params := []reflect.Value{reflect.ValueOf(req), reflect.ValueOf(req.getResponse())}
		reflectFunctionController.Call(params)

		return
	}
}
